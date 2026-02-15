package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model

	Balance float64
	Status  string

	BranchID uint
	Branch   Branch

	PrimaryCustomerID uint
	PrimaryCustomer   Customer

	SecondaryCustomerID *uint
	SecondaryCustomer   *Customer
}
