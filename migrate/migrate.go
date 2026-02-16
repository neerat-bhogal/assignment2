package main

import (
	"bankingSystem/config"
	"bankingSystem/models"
	"log"
)

func main() {

	// Connect to database
	config.ConnectDatabase()

	log.Println("Running database migrations...")

	// Ran AutoMigrate
	err := config.DB.AutoMigrate(
		&models.Bank{},
		&models.Branch{},
		&models.Customer{},
		&models.Account{},
		&models.Loan{},
		&models.Transaction{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println(" Database migrated successfully")
}
