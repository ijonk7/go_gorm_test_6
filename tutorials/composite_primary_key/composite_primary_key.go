package composite_primary_key

import "gorm.io/gorm"

type Product2 struct {
	ID           string `gorm:"primaryKey"`
	LanguageCode string `gorm:"primaryKey"`
	Code         string
	Name         string
}

/*
	Composite Primary Key
*/
func CompositePrimaryKey(db *gorm.DB) {
	db.AutoMigrate(Product2{})
}
