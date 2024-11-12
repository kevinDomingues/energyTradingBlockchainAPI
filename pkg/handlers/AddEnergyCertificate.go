package handlers

import (
	"bytes"
	"encoding/json"
	"energyTradingBlockchainAPI/pkg/models"
	"fmt"
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

	createCertificateRequest.RegulatoryAuthorityID = "2"

	userClaims := claims.(jwt.MapClaims)
	userID := userClaims["sum"].(string)
	blockchainToken := userClaims["btkn"].(string)

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
