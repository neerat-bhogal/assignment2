package controller

import (
	"net/http"
	"strconv"

	"bankingSystem/config"
	"bankingSystem/models"

	"github.com/gin-gonic/gin"
)

// ✅ Deposit Money
func Deposit(c *gin.Context) {
	var input struct {
		AccountID uint    `json:"account_id"`
		Amount    float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid deposit amount"})
		return
	}

	tx := config.DB.Begin()

	var account models.Account
	if err := tx.First(&account, input.AccountID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	account.Balance += input.Amount

	transaction := models.Transaction{
		Type:      "DEPOSIT",
		Amount:    input.Amount,
		AccountID: account.ID,
	}

	if err := tx.Save(&account).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update balance"})
		return
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"message":     "Deposit successful",
		"new_balance": account.Balance,
	})
}

// ✅ Withdraw Money
func Withdraw(c *gin.Context) {
	var input struct {
		AccountID uint    `json:"account_id"`
		Amount    float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid withdrawal amount"})
		return
	}

	tx := config.DB.Begin()

	var account models.Account
	if err := tx.First(&account, input.AccountID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	if account.Balance < input.Amount {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
		return
	}

	account.Balance -= input.Amount

	transaction := models.Transaction{
		Type:      "WITHDRAW",
		Amount:    input.Amount,
		AccountID: account.ID,
	}

	if err := tx.Save(&account).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update balance"})
		return
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"message":     "Withdrawal successful",
		"new_balance": account.Balance,
	})
}

// ✅ Get All Transactions
func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction

	if err := config.DB.Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

// ✅ Get Transactions By Account
func GetTransactionsByAccount(c *gin.Context) {
	id := c.Param("account_id")

	accountID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	var transactions []models.Transaction

	if err := config.DB.
		Where("account_id = ?", accountID).
		Find(&transactions).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
