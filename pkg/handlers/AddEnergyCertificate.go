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

func AddEnergyCertificate(c *gin.Context) {
	blockhainURL := os.Getenv("BLOCKCHAIN_URL")
	addEnergyCertificateUrl := blockhainURL + "/invoke/channel1/energyTradingBlockchain"

	var energyCertificate models.EnergyCertificate

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err := c.ShouldBindJSON(&energyCertificate)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	userClaims := claims.(jwt.MapClaims)
	userID := userClaims["sum"].(string)
	blockchainToken := userClaims["btkn"].(string)

	blockchainMethod := models.BlockchainMethod{
		Method: "EnergyCertificateContract:CreateEnergyCertificate",
		Args:   []string{userID, userID, time.Now().Format("2006-01-02"), fmt.Sprintf("%d", energyCertificate.UsableMonth), fmt.Sprintf("%d", energyCertificate.UsableYear), energyCertificate.RegulatoryAuthorityID},
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

	response, error := client.Do(request)
	if error != nil {
		c.JSON(400, gin.H{
			"error": "Cannot do request: " + error.Error(),
		})
		return
	}

	defer response.Body.Close()
	c.Status(200)
}
