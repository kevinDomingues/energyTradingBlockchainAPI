package handlers

import (
	"bytes"
	"encoding/json"
	"energyTradingBlockchainAPI/pkg/database"
	"energyTradingBlockchainAPI/pkg/models"
	"energyTradingBlockchainAPI/pkg/services"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {

	db := database.GetDatabase()

	blockhainURL := os.Getenv("BLOCKCHAIN_URL")
	getTokenUrl := blockhainURL + "/user/enroll"
	addUserUrl := blockhainURL + "/user/register"

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	user.Password = services.SHA256ENCODER(user.Password)

	user.ID = uuid.NewString()

	token, err := services.GetBearerToken(c, getTokenUrl, os.Getenv("BLOCKCHAIN_ADMIN"), os.Getenv("BLOCKCHAIN_ADMIN_SECRET"))

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to retrieve token: " + err.Error(),
		})
		return
	}

	blockchainUser := models.BlockchainUser{
		Id:     user.ID,
		Secret: user.Password,
	}

	user.BlockchainUser = user.ID

	data, err := json.Marshal(blockchainUser)
	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(data)

	request, errorPost := http.NewRequest("POST", addUserUrl, reader)
	if errorPost != nil {
		panic(errorPost)
	}

	request.Header.Set("Authorization", "Bearer "+token)

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	response, error := client.Do(request)
	if error != nil {
		c.JSON(400, gin.H{
			"error": "Cannot do request: " + error.Error(),
		})
		return
	}

	err = db.Create(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create user: " + err.Error(),
		})
		return
	}

	defer response.Body.Close()
	c.Status(204)
}
