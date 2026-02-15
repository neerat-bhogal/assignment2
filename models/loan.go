package models

import "gorm.io/gorm"

type Loan struct {
	gorm.Model
	Principal    float64
	Remaining    float64
	InterestRate float64
	Status       string

	CustomerID uint
	BankID     uint
}
