package services

import (
	"bytes"
	"encoding/json"
	"energyTradingBlockchainAPI/pkg/models"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBearerToken(c *gin.Context, getTokenUrl string, userId string, userSecret string) (string, error) {
	blockchainUser := models.BlockchainUser{
		Id:     userId,
		Secret: userSecret,
	}

	data, err := json.Marshal(blockchainUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to marshal JSON: " + err.Error(),
		})
		return "", err
	}

	reader := bytes.NewReader(data)

	request, err := http.NewRequest("POST", getTokenUrl, reader)
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var tokenResponse models.TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", err
	}

	return tokenResponse.Token, nil
}
