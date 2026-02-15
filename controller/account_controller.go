package controller

import (
	"net/http"

	"bankingSystem/config"
	"bankingSystem/models"

	"github.com/gin-gonic/gin"
)

// ✅ Open Account (Primary Holder)
func CreateAccount(c *gin.Context) {
	var account models.Account

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check customer exists
	var customer models.Customer
	if err := config.DB.First(&customer, account.PrimaryCustomerID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found"})
		return
	}

	// Check branch exists
	var branch models.Branch
	if err := config.DB.First(&branch, account.BranchID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch not found"})
		return
	}

	account.Status = "ACTIVE"
	account.Balance = 0

	if err := config.DB.Create(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create account"})
		return
	}

	c.JSON(http.StatusCreated, account)
}

// ✅ Add Joint Holder
func AddJointHolder(c *gin.Context) {
	accountID := c.Param("id")

	var input struct {
		CustomerID uint `json:"customer_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var account models.Account
	if err := config.DB.First(&account, accountID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	// verify customer
	var customer models.Customer
	if err := config.DB.First(&customer, input.CustomerID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found"})
		return
	}

	account.SecondaryCustomerID = &input.CustomerID

	if err := config.DB.Save(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add joint holder"})
		return
	}

	c.JSON(http.StatusOK, account)
}

// ✅ Get Account Details
func GetAccount(c *gin.Context) {
	id := c.Param("id")

	var account models.Account

	if err := config.DB.
		Preload("PrimaryCustomer").
		Preload("SecondaryCustomer").
		Preload("Branch").
		First(&account, id).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	c.JSON(http.StatusOK, account)
}

// ✅ Get All Accounts of Customer
func GetCustomerAccounts(c *gin.Context) {
	id := c.Param("customer_id")

	var accounts []models.Account

	if err := config.DB.
		Where("primary_customer_id = ? OR secondary_customer_id = ?", id, id).
		Find(&accounts).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch accounts"})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

// ✅ Close Account
func CloseAccount(c *gin.Context) {
	id := c.Param("id")

	var account models.Account
	if err := config.DB.First(&account, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	if account.Balance != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Account balance must be zero to close"})
		return
	}

	account.Status = "CLOSED"

	if err := config.DB.Save(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to close account"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account closed successfully"})
}
