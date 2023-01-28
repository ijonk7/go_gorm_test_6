package model

import (
	"gorm.io/gorm"
	"time"
)

type User2 struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Alamat    string
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
