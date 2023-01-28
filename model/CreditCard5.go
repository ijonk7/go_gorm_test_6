package model

import "gorm.io/gorm"

type CreditCard5 struct {
	gorm.Model
	Number  string
	User5ID uint
}
