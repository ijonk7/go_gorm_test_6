package model

import "gorm.io/gorm"

type User4 struct {
	gorm.Model
	CreditCard CreditCard4
}
