package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"bankingSystem/config"
	"bankingSystem/models"

	"github.com/gin-gonic/gin"
)

// to Create bank
func CreateBank(c *gin.Context) {
	var bank models.Bank
	fmt.Println("here")
	if err := c.ShouldBindJSON(&bank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if bank.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bank name is required"})
		return
	}

	if err := config.DB.Create(&bank).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bank"})
		return
	}

	c.JSON(http.StatusCreated, bank)
}

// then Get all banks
func GetBanks(c *gin.Context) {
	var banks []models.Bank

	if err := config.DB.
		Preload("Branches").
		Find(&banks).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch banks"})
		return
	}

	c.JSON(http.StatusOK, banks)
}

// then Get bank by ID type
func GetBankByID(c *gin.Context) {
	id := c.Param("id")

	bankID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bank ID"})
		return
	}

	var bank models.Bank

	if err := config.DB.
		Preload("Branches").
		Preload("Loans").
		First(&bank, bankID).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": "Bank not found"})
		return
	}

	c.JSON(http.StatusOK, bank)
}

//  Update Bank
func UpdateBank(c *gin.Context) {
	id := c.Param("id")

	var bank models.Bank

	if err := config.DB.First(&bank, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bank not found"})
		return
	}

	var input models.Bank
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name != "" {
		bank.Name = input.Name
	}

	if err := config.DB.Save(&bank).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bank"})
		return
	}

	c.JSON(http.StatusOK, bank)
}

// Delete Bank
func DeleteBank(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Bank{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete bank"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bank deleted successfully"})
}
