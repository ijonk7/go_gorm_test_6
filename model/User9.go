package model

import (
	"gorm.io/gorm"
	"time"
)

type User9 struct {
	ID        uint `gorm:"primaryKey; column:id_user;"`
	Name      string
	Alamat    string `gorm:"primaryKey; column:alamat_rumah;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
