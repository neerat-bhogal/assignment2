package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name  string
	Email string
	Phone string

	// One-to-Many (Customer -> Loans)
	Loans []Loan

	// Many-to-Many (Customer <-> Accounts)
	Accounts []Account `gorm:"many2many:account_customers;"`
}
