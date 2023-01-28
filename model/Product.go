package model

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        uint
	Code      string
	Price     uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
