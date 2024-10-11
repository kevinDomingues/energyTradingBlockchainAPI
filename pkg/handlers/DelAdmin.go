package handlers

import (
	"energyTradingBlockchainAPI/pkg/database"
	"energyTradingBlockchainAPI/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DelAdmin(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Admin{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot delete admin: " + err.Error(),
		})
		return
	}
	c.Status(204)
}
