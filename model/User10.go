package model

import (
	"gorm.io/gorm"
	"time"
)

type User10 struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
