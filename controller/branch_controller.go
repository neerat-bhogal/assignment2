package controller

import (
	"net/http"
	"strconv"

	"bankingSystem/config"
	"bankingSystem/models"

	"github.com/gin-gonic/gin"
)

// Creating Branch
func CreateBranch(c *gin.Context) {
	var branch models.Branch

	if err := c.ShouldBindJSON(&branch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if Bank exists
	var bank models.Bank
	if err := config.DB.First(&bank, branch.BankID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bank not found"})
		return
	}

	if err := config.DB.Create(&branch).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create branch"})
		return
	}

	c.JSON(http.StatusCreated, branch)
}

// gettig all Branches
func GetBranches(c *gin.Context) {
	var branches []models.Branch

	if err := config.DB.Preload("Bank").Find(&branches).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch branches"})
		return
	}

	c.JSON(http.StatusOK, branches)
}

// geting branch by ID
func GetBranchByID(c *gin.Context) {
	id := c.Param("id")

	branchID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var branch models.Branch

	if err := config.DB.
		Preload("Bank").
		Preload("Accounts").
		First(&branch, branchID).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": "Branch not found"})
		return
	}

	c.JSON(http.StatusOK, branch)
}

// Update Branch
func UpdateBranch(c *gin.Context) {
	id := c.Param("id")

	var branch models.Branch

	if err := config.DB.First(&branch, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Branch not found"})
		return
	}

	var input models.Branch
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	branch.Name = input.Name

	if err := config.DB.Save(&branch).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update branch"})
		return
	}

	c.JSON(http.StatusOK, branch)
}

//  Delete Branch
func DeleteBranch(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Branch{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete branch"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Branch deleted successfully"})
}
