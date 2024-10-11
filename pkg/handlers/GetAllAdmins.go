package handlers

import (
	"energyTradingBlockchainAPI/pkg/database"
	"energyTradingBlockchainAPI/pkg/models"

	"github.com/gin-gonic/gin"
)

func GetAllAdmins(c *gin.Context) {
	db := database.GetDatabase()

	var admin []models.Admin
	err := db.Find(&admin).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list user: " + err.Error(),
		})
		return
	}
	c.JSON(200, admin)
}
