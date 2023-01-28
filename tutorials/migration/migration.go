package migration

import (
	"fmt"
	"gorm.io/gorm"
	"gorm_test_6/model"
)

/*
	Migration - CurrentDatabase
*/
func CurrentDatabase(db *gorm.DB) {
	dbName := db.Migrator().CurrentDatabase()
	fmt.Println(dbName)
}

/*
	Migration - Tables
*/
func Tables(db *gorm.DB) {
	// Create table for `User`
	//db.Migrator().CreateTable(&model.User7{})

	// Rename old table to new table
	db.Migrator().RenameTable("user7", "user77")
}

/*
	Migration - Columns
*/
func Columns(db *gorm.DB) {
	// Rename column to new name
	db.Migrator().RenameColumn(&model.User7{}, "address", "alamat")
}
