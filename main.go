package main

import (
	"bankingSystem/config"
	"bankingSystem/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// here Connect DB
	config.ConnectDatabase()

<<<<<<< HEAD
	// then Run migrations
	// Config.Migrate()
=======
	// Run migrations
	// config.Migrate() // TODO: Implement Migrate in config package or remove this line if not needed
>>>>>>> f88b8e0 (Initial commit: transport all files)

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")
}
