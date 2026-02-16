package main

import (
	"bankingSystem/config"
	"bankingSystem/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// here Connect DB
	config.ConnectDatabase()

	// then Run migrations
	// Config.Migrate()

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")
}
