package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name  string
	Email string
	Phone string

	// One-to-Many type
	Loans []Loan

	// Many-to-Many type for account
	Accounts []Account `gorm:"many2many:account_customers;"`
}
