package routes

import (
	"energyTradingBlockchainAPI/pkg/handlers"

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
	}
	return router
}
