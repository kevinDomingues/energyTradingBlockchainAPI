package handlers

import (
	"energyTradingBlockchainAPI/pkg/database"
	"energyTradingBlockchainAPI/pkg/models"
	"energyTradingBlockchainAPI/pkg/services"
	"os"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	db := database.GetDatabase()

	blockhainURL := os.Getenv("BLOCKCHAIN_URL")
	getTokenUrl := blockhainURL + "/user/enroll"

	var login models.Login

	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	var dbUser models.User
	dbError := db.Where("email = ?", login.Email).First(&dbUser).Error

	if dbError != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find admin",
		})
		return
	}

	if dbUser.Password != services.SHA256ENCODER(login.Password) {
		c.JSON(401, gin.H{
			"error": "Invalid Credentials",
		})
		return
	}

	blockchainToken, err := services.GetBearerToken(c, getTokenUrl, dbUser.ID, dbUser.Password)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to retrieve token: " + err.Error(),
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(dbUser.ID, blockchainToken, dbUser.UserType)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
		"role":  dbUser.UserType,
	})
}
