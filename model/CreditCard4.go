package model

import "gorm.io/gorm"

type CreditCard4 struct {
	gorm.Model
	Number  string
	User4ID uint
}
