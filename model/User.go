package model

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint
	Name         string
	CreditCard   CreditCard
	Email        *string `gorm:"default:''"`
	Age          uint8   `gorm:"default:10"`
	Birthday     *time.Time
	MemberNumber sql.NullString
	Active       sql.NullBool `gorm:"default:false"`
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Age >= 50 {
		return errors.New("Usia Lebih dari 20 Tahun")
	}
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed() {
		fmt.Println("DATA BERUBAH")
	} else {
		fmt.Println("DATA TIDAK BERUBAH")
	}

	//if tx.Statement.Changed("Name") {
	//	tx.Statement.SetColumn("Age", 36)
	//}
	return nil
}
