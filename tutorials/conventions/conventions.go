package conventions

import (
	"gorm.io/gorm"
	"gorm_test_6/model"
	"time"
)

/*
	Conventions - PluralizedTableName
*/
func PluralizedTableName(db *gorm.DB) {
	db.AutoMigrate(&model.User8{})

	//var users []model.User8
	//db.Find(&users)
	////db.Table("profiles_8").Find(&users)
	//fmt.Println(users)
}

/*
	Conventions - Temporarily specify a name
*/
func TemporarilySpecifyName(db *gorm.DB) {
	// Create table `deleted_users` with struct User's fields
	db.Table("deleted_users").AutoMigrate(&model.User8{})
}

/*
	Conventions - Column Name
*/
func ColumnName(db *gorm.DB) {
	db.AutoMigrate(&model.User9{})
}

/*
	Conventions - Timestamp Tracking
*/
func TimestampTracking(db *gorm.DB) {
	//db.AutoMigrate(&model.User10{})

	//user2 := model.User10{Name: "jinzhu 3", CreatedAt: time.Now()}
	//db.Create(&user2)

	user3 := model.User10{ID: 2}
	db.Model(&user3).Update("CreatedAt", time.Now())
}
