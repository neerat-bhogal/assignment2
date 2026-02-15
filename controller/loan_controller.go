package controller

import (
	"net/http"

	"bankingSystem/config"
	"bankingSystem/models"

	"github.com/gin-gonic/gin"
)

// ✅ Create Loan
func CreateLoan(c *gin.Context) {
	var loan models.Loan

	if err := c.ShouldBindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if Customer exists
	var customer models.Customer
	if err := config.DB.First(&customer, loan.CustomerID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found"})
		return
	}

	// Default values
	loan.InterestRate = 12
	loan.Remaining = loan.Principal
	loan.Status = "ACTIVE"

	if err := config.DB.Create(&loan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create loan"})
		return
	}

	c.JSON(http.StatusCreated, loan)
}

// ✅ Get All Loans
func GetLoans(c *gin.Context) {
	var loans []models.Loan

	if err := config.DB.Find(&loans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch loans"})
		return
	}

	c.JSON(http.StatusOK, loans)
}

// ✅ Get Loan By ID
func GetLoanByID(c *gin.Context) {
	id := c.Param("id")

	var loan models.Loan

	if err := config.DB.First(&loan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	c.JSON(http.StatusOK, loan)
}

// ✅ Calculate Yearly Interest
func CalculateInterest(c *gin.Context) {
	id := c.Param("id")

	var loan models.Loan

	if err := config.DB.First(&loan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	if loan.Status != "ACTIVE" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Loan is not active"})
		return
	}

	interest := loan.Remaining * (loan.InterestRate / 100)

	c.JSON(http.StatusOK, gin.H{
		"loan_id":         loan.ID,
		"remaining":       loan.Remaining,
		"interest_rate":   loan.InterestRate,
		"yearly_interest": interest,
	})
}

// ✅ Repay Loan (Atomic Transaction)
func RepayLoan(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Amount float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := config.DB.Begin()

	var loan models.Loan
	if err := tx.First(&loan, id).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	if loan.Status != "ACTIVE" {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Loan already closed"})
		return
	}

	if input.Amount <= 0 {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid repayment amount"})
		return
	}

	loan.Remaining -= input.Amount

	if loan.Remaining <= 0 {
		loan.Remaining = 0
		loan.Status = "CLOSED"
	}

	if err := tx.Save(&loan).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update loan"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, loan)
}
