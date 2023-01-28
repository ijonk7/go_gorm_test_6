package model

import "gorm.io/gorm"

type User6 struct {
	gorm.Model
	Languages []*Language `gorm:"many2many:user_languages;"`
}
