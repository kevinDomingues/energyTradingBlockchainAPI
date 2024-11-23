package routes

import (
	"energyTradingBlockchainAPI/pkg/handlers"
	"energyTradingBlockchainAPI/pkg/middlewares"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	main := router.Group("api/v1")
	{
		Admin := main.Group("admins")
		{
			Admin.GET("/:id", handlers.GetAdminById)
			Admin.GET("/", handlers.GetAllAdmins)
			Admin.POST("/", handlers.AddAdmin)
			Admin.DELETE("/:id", handlers.DelAdmin)
			Admin.PUT("/", handlers.UpdateAdmin)
		}
		User := main.Group("user")
		{
			User.POST("/create", handlers.AddUser)
		}
		EnergyCertificate := main.Group("certificate", middlewares.Auth())
		{
			EnergyCertificate.POST("/create", handlers.AddEnergyCertificate)
			EnergyCertificate.GET("/producer/:id", handlers.GetEnergyCertificateByProducerId)
			EnergyCertificate.GET("/owned", handlers.GetEnergyCertificateByOwnerId)
			EnergyCertificate.GET("/from/:usableMonth/:usableYear", handlers.GetCertificatesAvailableFromSpecificMonth)
			EnergyCertificate.GET("/from/:usableMonth/:usableYear/type/:energyType", handlers.GetCertificatesAvailableFromSpecificMonthAndEnergyType)
			EnergyCertificate.POST("/transfer", handlers.TransferEnergyCertificate)
		}
		Consumptions := main.Group("consumptions", middlewares.Auth())
		{
			Consumptions.GET("/:year", handlers.GetConsumptionByYear)
			Consumptions.GET("/from/:month/:year", handlers.GetConsumptionFromSpecificMonth)
		}
		main.POST("login", handlers.Login)
	}
	return router
}
