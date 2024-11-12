package handlers

import (
	"bytes"
	"encoding/json"
	"energyTradingBlockchainAPI/pkg/database"
	"energyTradingBlockchainAPI/pkg/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TransferEnergyCertificate(c *gin.Context) {
	blockhainURL := os.Getenv("BLOCKCHAIN_URL")
	transferEnergyCertificateUrl := blockhainURL + "/invoke/channel1/energyTradingBlockchain"

	var transferCertificateBody models.TransferCertificate

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err := c.ShouldBindJSON(&transferCertificateBody)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind transferCertificate JSON: " + err.Error(),
		})
		return
	}

	userClaims := claims.(jwt.MapClaims)
	userID := userClaims["sum"].(string)
	blockchainToken := userClaims["btkn"].(string)

	energyCertificate, err := getEnergyCertificate(transferCertificateBody.EnergyCertificateID, blockchainToken, transferEnergyCertificateUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	oldOwnerId := energyCertificate.OwnerID

	price := 12.3

	if err := transferEnergyCertificate(transferCertificateBody.EnergyCertificateID, userID, price, blockchainToken, transferEnergyCertificateUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := createConsumption(userID, energyCertificate.UsableYear, energyCertificate.UsableMonth, energyCertificate.EnergyType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create consumption record: " + err.Error()})
		return
	}

	if err := createConsumption(oldOwnerId, energyCertificate.UsableYear, energyCertificate.UsableMonth, energyCertificate.EnergyType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create consumption record: " + err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func makeBlockchainRequest(method string, args []string, blockchainToken string, url string) ([]byte, error) {
	blockchainMethod := models.BlockchainMethod{
		Method: method,
		Args:   args,
	}

	data, err := json.Marshal(blockchainMethod)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal blockchain method: %v", err)
	}

	reader := bytes.NewBuffer(data)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	request.Header.Set("Authorization", "Bearer "+blockchainToken)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return body, nil
}

func getEnergyCertificate(energyCertificateID, blockchainToken, url string) (models.EnergyCertificate, error) {
	body, err := makeBlockchainRequest("EnergyCertificateContract:ReadEnergyCertificate", []string{energyCertificateID}, blockchainToken, url)
	if err != nil {
		return models.EnergyCertificate{}, err
	}

	var energyCertificateResponse struct {
		Response models.EnergyCertificate `json:"response"`
	}
	if err := json.Unmarshal(body, &energyCertificateResponse); err != nil {
		return models.EnergyCertificate{}, fmt.Errorf("failed to unmarshal energy certificate: %v", err)
	}

	return energyCertificateResponse.Response, nil
}

func transferEnergyCertificate(energyCertificateID string, userID string, price float64, blockchainToken string, url string) error {
	_, err := makeBlockchainRequest("EnergyCertificateContract:TransferEnergyCertificate", []string{energyCertificateID, userID, fmt.Sprintf("%.2f", price)}, blockchainToken, url)
	if err != nil {
		return fmt.Errorf("failed to transfer energy certificate: %v", err)
	}
	return nil
}

func createConsumption(userID string, year int, month int, energyType int) error {
	db := database.GetDatabase()

	var consumption models.Consumption
	err := db.Where("user_id = ? AND consumption_year = ? AND consumption_month = ? AND energy_type_id = ?", userID, year, month, energyType).First(&consumption).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return fmt.Errorf("failed to find existing consumption record: %v", err)
	}

	if err == gorm.ErrRecordNotFound {
		newConsumption := models.Consumption{
			UserID:           userID,
			ConsumptionYear:  year,
			ConsumptionMonth: month,
			EnergyTypeId:     energyType,
			EnergyConsumed:   1,
		}
		if err := db.Create(&newConsumption).Error; err != nil {
			return fmt.Errorf("failed to create new consumption record: %v", err)
		}
	} else {
		consumption.EnergyConsumed = consumption.EnergyConsumed + 1
		if err := db.Save(&consumption).Error; err != nil {
			return fmt.Errorf("failed to update consumption record: %v", err)
		}
	}

	return nil
}
