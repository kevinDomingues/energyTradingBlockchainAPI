package handlers

import (
	"energyTradingBlockchainAPI/pkg/database"
	"energyTradingBlockchainAPI/pkg/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetConsumptionByYear(c *gin.Context) {
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year format"})
		return
	}
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userClaims := claims.(jwt.MapClaims)
	userID := userClaims["sum"].(string)

	db := database.GetDatabase()

	var consumptions []models.Consumption
	err = db.Model(&models.Consumption{}).
		Select("user_id, consumption_year, consumption_month, energy_type_id, SUM(energy_consumed) AS energy_consumed").
		Where("user_id = ? AND consumption_year = ?", userID, year).
		Group("user_id, consumption_year, consumption_month, energy_type_id").
		Order("consumption_month, energy_type_id").
		Scan(&consumptions).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, consumptions)
}
