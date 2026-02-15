package routes

import (
	"bankingSystem/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	api := router.Group("/api/v1")
	{

		// ---------------- BANK ROUTES ----------------
		banks := api.Group("/banks")
		{
			banks.POST("/", controller.CreateBank)
			banks.GET("/", controller.GetBanks)
			banks.GET("/:id", controller.GetBankByID)
			banks.PUT("/:id", controller.UpdateBank)
			banks.DELETE("/:id", controller.DeleteBank)
		}

		// ---------------- BRANCH ROUTES ----------------
		branches := api.Group("/branches")
		{
			branches.POST("/", controller.CreateBranch)
			branches.GET("/", controller.GetBranches)
			branches.GET("/:id", controller.GetBranchByID)
			branches.PUT("/:id", controller.UpdateBranch)
			branches.DELETE("/:id", controller.DeleteBranch)
		}

		// ---------------- CUSTOMER ROUTES ----------------
		customers := api.Group("/customers")
		{
			customers.POST("/", controller.CreateCustomer)
			customers.GET("/", controller.GetCustomers)
			customers.GET("/:id", controller.GetCustomerByID)
			customers.DELETE("/:id", controller.DeleteCustomer)
			customers.GET("/:id/accounts", controller.GetCustomerAccounts)
		}

		// ---------------- ACCOUNT ROUTES ----------------
		accounts := api.Group("/accounts")
		{
			accounts.POST("/", controller.CreateAccount)
			accounts.GET("/:id", controller.GetAccount)
			accounts.PUT("/:id/joint", controller.AddJointHolder)
			accounts.PUT("/:id/close", controller.CloseAccount)
		}

		// ---------------- LOAN ROUTES ----------------
		loans := api.Group("/loans")
		{
			loans.POST("/", controller.CreateLoan)
			loans.GET("/", controller.GetLoans)
			loans.GET("/:id", controller.GetLoanByID)
			loans.GET("/:id/interest", controller.CalculateInterest)
			loans.PUT("/:id/repay", controller.RepayLoan)
		}

		// ---------------- TRANSACTION ROUTES ----------------
		transactions := api.Group("/transactions")
		{
			transactions.POST("/deposit", controller.Deposit)
			transactions.POST("/withdraw", controller.Withdraw)
			transactions.GET("/", controller.GetTransactions)
			transactions.GET("/account/:account_id", controller.GetTransactionsByAccount)
		}
	}
}
