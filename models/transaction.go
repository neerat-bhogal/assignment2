package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Type      string // deposit / withdraw
	Amount    float64
	AccountID uint
}
