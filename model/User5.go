package model

import (
	"gorm.io/gorm"
	"time"
)

type User5 struct {
	gorm.Model
	ID          uint `gorm:"primarykey"`
	Name        string
	CreditCards []CreditCard5
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
