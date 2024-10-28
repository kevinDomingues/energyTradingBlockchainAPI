package main

import (
	"log"

	"energyTradingBlockchainAPI/pkg/database"
	"energyTradingBlockchainAPI/pkg/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}

}

func Run() {
	router := routes.ConfigRoutes(gin.Default())

	log.Print("Server is running at port: ", "5000")
	router.Run(":5000")
}

func main() {
	database.StartDB()

	Run()
}
