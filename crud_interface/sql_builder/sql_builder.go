package sql_builder

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"gorm_test_6/model"
)

/*
	Raw SQL
*/
func RawSql(db *gorm.DB) {
	//type Result struct {
	//	ID   int
	//	Name string
	//	Age  int
	//}
	//var result Result
	//
	//db.Raw("SELECT id, name, age FROM users WHERE id = ?", 8).Scan(&result)
	//fmt.Println(result)

	// Exec with Raw SQL
	db.Exec("UPDATE users SET name = ? WHERE id = ?", "jinzhu 666", 6)
}

/*
	Named Argument (Argumen Bernama)
	GORM mendukung Named Argument (Argumen Bernama) dengan  sql.NamedArg, map[string]interface{}{} or struct
*/
func NamedArgument(db *gorm.DB) {
	var users []model.User
	db.Where("name = @name OR age = @age", sql.Named("name", "jinzhu 88"), sql.Named("age", 55)).Find(&users)
	fmt.Println(users)
}

/*
	DryRun Mode
	Hasilkan SQL dan argumennya tanpa mengeksekusi, dapat digunakan untuk menyiapkan atau menguji SQL yang dihasilkan, lihat Session untuk detailnya
*/
func DryRunMode(db *gorm.DB) {
	var user model.User

	stmt := db.Session(&gorm.Session{DryRun: true}).Find(&user).Statement

	fmt.Println(stmt.SQL.String()) // Generate SQL
	fmt.Println(stmt.Vars)         // Generate Arguments

	// Untuk menghasilkan SQL akhir, Anda dapat menggunakan kode berikut:
	result := db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...)
	fmt.Println(result)
}

/*
	ToSQL
	Returns yang generated SQL tanpa dieksekusi.
*/
func ToSQL(db *gorm.DB) interface{} {
	var user model.User

	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Find(&user)
	})
	return sql
}

/*
	Row & Rows
*/
func RowRows(db *gorm.DB) {
	// Get result as *sql.Row
	var name string
	var age int

	row := db.Table("users").Where("name = ?", "jinzhu 88").Select("name", "age").Row()
	row.Scan(&name, &age)
	fmt.Println(name)
	fmt.Println(age)

	// Get result as *sql.Rows
	//rows, _ := db.Model(&model.User{}).Where("name = ?", "jinzhu 88").Select("name", "age").Rows()
	//defer rows.Close()
	//for rows.Next() {
	//	var name string
	//	var age uint8
	//
	//	rows.Scan(&name, &age)
	//	fmt.Println(name)
	//	fmt.Println(age)
	//
	//	// Scan *sql.Rows into struct
	//	//var user model.User
	//	//db.ScanRows(rows, &user)
	//	//fmt.Println(user.Name)
	//	//fmt.Println(user.Age)
	//}
}
