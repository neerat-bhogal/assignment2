package main

import (
	"bankingSystem/config"
	"bankingSystem/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// Connect DB
	config.ConnectDatabase()

	// Run migrations
	// Config.Migrate()

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")
}
