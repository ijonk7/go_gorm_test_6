package model

import "gorm.io/gorm"

type User3 struct {
	gorm.Model
	Name         string
	CompanyRefer int
	Company      Company `gorm:"foreignKey:CompanyRefer;references:Code"`
}
