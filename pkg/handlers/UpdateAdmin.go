package handlers

import (
	"energyTradingBlockchainAPI/pkg/database"
	"energyTradingBlockchainAPI/pkg/models"

	"github.com/gin-gonic/gin"
)

func UpdateAdmin(c *gin.Context) {

	db := database.GetDatabase()

	var admin models.Admin

	err := c.ShouldBindJSON(&admin)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&admin).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot update admin: " + err.Error(),
		})
		return
	}

	c.JSON(200, admin)

}
