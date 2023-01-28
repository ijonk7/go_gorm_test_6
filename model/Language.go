package model

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	Name  string
	Users []*User6 `gorm:"many2many:user_languages;"`
}
