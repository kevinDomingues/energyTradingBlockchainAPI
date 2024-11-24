package handlers

import (
	"bytes"
	"encoding/json"
	"energyTradingBlockchainAPI/pkg/models"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CreateCertificateRequest struct {
	models.EnergyCertificate
	Quantity int `json:"quantity"`
}

func AddEnergyCertificate(c *gin.Context) {
	blockhainURL := os.Getenv("BLOCKCHAIN_URL")
	addEnergyCertificateUrl := blockhainURL + "/invoke/channel1/energyTradingBlockchain"

	var createCertificateRequest CreateCertificateRequest

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err := c.ShouldBindJSON(&createCertificateRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	quantity := createCertificateRequest.Quantity
	if quantity < 1 {
		quantity = 1
	}

	userClaims := claims.(jwt.MapClaims)
	userID := userClaims["sum"].(string)
	blockchainToken := userClaims["btkn"].(string)

	validatedUser, err := validateUser(userID)
	if err != nil {
		c.JSON((http.StatusInternalServerError), gin.H{"error": err.Error()})
		return
	}

	if !validatedUser.Accepted {
		c.JSON((http.StatusForbidden), gin.H{"error": "Regulatory Authority refused this request"})
		return
	}

	createCertificateRequest.RegulatoryAuthorityID = validatedUser.RegulatoryAuthority

	for i := 0; i < quantity; i++ {
		blockchainMethod := models.BlockchainMethod{
			Method: "EnergyCertificateContract:CreateEnergyCertificate",
			Args: []string{
				userID,
				userID,
				time.Now().Format("2006-01-02"),
				fmt.Sprintf("%d", createCertificateRequest.UsableMonth),
				fmt.Sprintf("%d", createCertificateRequest.UsableYear),
				createCertificateRequest.RegulatoryAuthorityID,
				fmt.Sprintf("%d", createCertificateRequest.EnergyType)},
		}

		data, err := json.Marshal(blockchainMethod)
		if err != nil {
			log.Fatal(err)
			c.Status(http.StatusBadRequest)
		}

		reader := bytes.NewBuffer(data)

		request, errorPost := http.NewRequest("POST", addEnergyCertificateUrl, reader)
		if errorPost != nil {
			panic(errorPost)
		}

		request.Header.Set("Authorization", "Bearer "+blockchainToken)

		request.Header.Set("Content-Type", "application/json")
		client := &http.Client{}

		response, err := client.Do(request)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Cannot do request: " + err.Error(),
			})
			return
		}

		defer response.Body.Close()
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d certificates created", quantity)})
}

func validateUser(userId string) (models.UserValidationResponse, error) {
	mockServerURL := os.Getenv("MOCK_SERVER_URL")
	validateUserUrl := mockServerURL + "/validate-user"

	userValidationRequest := models.UserValidation{
		UserId: userId,
	}

	data, err := json.Marshal(userValidationRequest)
	if err != nil {
		return models.UserValidationResponse{}, fmt.Errorf("failed to marshal userValidationRequest method: %v", err)
	}

	reader := bytes.NewBuffer(data)
	request, err := http.NewRequest("POST", validateUserUrl, reader)
	if err != nil {
		return models.UserValidationResponse{}, fmt.Errorf("failed to create request: %v", err)
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return models.UserValidationResponse{}, fmt.Errorf("failed to send request: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return models.UserValidationResponse{}, fmt.Errorf("failed to read response body: %v", err)
	}

	var userValidationResponse models.UserValidationResponse
	err = json.Unmarshal(body, &userValidationResponse)
	if err != nil {
		return models.UserValidationResponse{}, err
	}

	return userValidationResponse, nil
}
