package session

import (
	"fmt"
	"gorm.io/gorm"
	"gorm_test_6/model"
)

/*
	DryRun
*/
func DryRun(db *gorm.DB) {
	var user model.User

	stmt := db.Session(&gorm.Session{DryRun: true}).Find(&user, 6).Statement

	fmt.Println(stmt.SQL.String()) // Generate SQL
	fmt.Println(stmt.Vars)         // Generate Arguments

	// Untuk menghasilkan SQL akhir, Anda dapat menggunakan kode berikut:
	result := db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...)
	fmt.Println(result)
}
