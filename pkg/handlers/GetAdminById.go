package handlers

import (
	"energyTradingBlockchainAPI/pkg/database"
	"energyTradingBlockchainAPI/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAdminById(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var admin models.Admin
	err = db.First(&admin, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find admin: " + err.Error(),
		})
		return
	}

	c.JSON(200, admin)

}
