package model

import (
	"gorm.io/gorm"
	"time"
)

type User8 struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Alamat    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

//type Tabler interface {
//	TableName() string
//}

// TableName overrides the table name used by User to `profiles`
func (User8) TableName() string {
	return "profiles_8"
}
