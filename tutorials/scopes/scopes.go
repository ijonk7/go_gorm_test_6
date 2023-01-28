package scopes

import (
	"fmt"
	"gorm.io/gorm"
	"gorm_test_6/model"
)

func CheckName(db *gorm.DB) *gorm.DB {
	return db.Where("name = ?", "jinzhu 4")
}

/*
	Scopes - Query
*/
func Query(db *gorm.DB) {
	var users []model.User
	db.Scopes(CheckName).Find(&users)

	fmt.Println(users)
}
