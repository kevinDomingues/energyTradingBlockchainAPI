package handlers

import (
	"bytes"
	"encoding/json"
	"energyTradingBlockchainAPI/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetTransactionsByToUserId(c *gin.Context) {
	blockhainURL := os.Getenv("BLOCKCHAIN_URL")
	getTransactionUrl := blockhainURL + "/invoke/channel1/energyTradingBlockchain"

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userClaims := claims.(jwt.MapClaims)
	userID := userClaims["sum"].(string)
	blockchainToken := userClaims["btkn"].(string)

	blockchainMethod := models.BlockchainMethod{
		Method: "EnergyCertificateContract:GetTransactionsByToUserID",
		Args:   []string{userID},
	}

	data, err := json.Marshal(blockchainMethod)
	if err != nil {
		log.Fatal(err)
		c.Status(http.StatusBadRequest)
	}

	reader := bytes.NewBuffer(data)

	request, errorPost := http.NewRequest("POST", getTransactionUrl, reader)
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

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read response body: " + err.Error(),
		})
		return
	}

	var responseBody map[string]interface{}
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to unmarshal JSON response: " + err.Error(),
		})
		return
	}

	c.JSON(response.StatusCode, responseBody)
}