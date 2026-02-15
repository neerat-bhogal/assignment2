package models

import "gorm.io/gorm"

type Branch struct {
	gorm.Model
	Name   string
	BankID uint

	Bank     Bank
	Accounts []Account
}
