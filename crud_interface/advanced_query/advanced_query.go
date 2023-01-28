package advanced_query

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/hints"
	"gorm_test_6/model"
	"time"
)

/*
	Smart Select Fields
*/
func SmartSelectFields(db *gorm.DB) {
	var users []model.User

	type APIUser struct {
		ID       uint
		Name     string
		Age      uint8
		Birthday *time.Time
		Active   sql.NullBool `gorm:"default:true"`
	}
	var apiUser []APIUser

	db.Model(&users).Limit(10).Find(&apiUser)
	//db.Model(&users).Select("id", "name", "age").Limit(10).Find(&apiUser)
	fmt.Println(apiUser)
}

/*
	First Or Init
*/
func FirstOrInit(db *gorm.DB) {
	var user model.User

	// Found user with `name` = `jinzhu`
	//db.Where(model.User{Name: "jinzhu"}).FirstOrInit(&user)
	// user -> User{ID: 111, Name: "Jinzhu", Age: 18}

	// Found user with `name` = `jinzhu`
	//db.FirstOrInit(&user, map[string]interface{}{"name": "jinzhu"})
	// user -> User{ID: 111, Name: "Jinzhu", Age: 18}

	// User not found, initialize it with give conditions and Attrs
	//db.Where(model.User{Name: "non_existing"}).Attrs(model.User{Age: 20}).FirstOrInit(&user)
	// SELECT * FROM USERS WHERE name = 'non_existing' ORDER BY id LIMIT 1;
	// user -> User{Name: "non_existing", Age: 20}

	db.Where(model.User{Name: "jinzhu2"}).Assign(model.User{Age: 20}).FirstOrInit(&user)

	fmt.Println(user)
}

/*
	First Or Create
	Dapatkan record pertama yang cocok atau buat yang baru dengan kondisi yg diberikan (hanya berfungsi dengan kondisi struct & map),
	RowsAffected mengembalikan jumlah record yang di create/update
*/
func FirstOrCreate(db *gorm.DB) {
	var user model.User

	//result := db.FirstOrCreate(&user, model.User{Name: "non_existing"})

	//result := db.Where(model.User{Name: "non_existing2"}).Attrs(model.User{Age: 20}).FirstOrCreate(&user)

	// User not found, create it with give conditions and Attrs
	//result := db.Where(model.User{Name: "non_existing"}).Attrs(model.User{Age: 30}).FirstOrCreate(&user)

	result := db.Where(model.User{Name: "non_existing_2"}).Assign(model.User{Age: 20}).FirstOrCreate(&user)

	fmt.Println(result.RowsAffected)
	fmt.Println(user)
}

/*
	Optimizer/Index Hints
*/
func OptimizerIndexHints(db *gorm.DB) {
	var user model.User

	result := db.Clauses(hints.New("hint")).Find(&model.User{})
	//result := db.Clauses(hints.New("MAX_EXECUTION_TIME(10000)")).Find(&model.User{})
	//result := db.Clauses(hints.UseIndex("idx_user_name")).Find(&model.User{})

	fmt.Println(user)
	fmt.Println(result)
}

/*
	Iteration
*/
func Iteration(db *gorm.DB) {
	rows, _ := db.Model(&model.User{}).Where("name = ?", "jinzhu").Rows()
	defer rows.Close()

	for rows.Next() {
		var user model.User
		// ScanRows is a method of `gorm.DB`, it can be used to scan a row into a struct
		db.ScanRows(rows, &user)

		// do something
		fmt.Println(user.Name)
		fmt.Println(user.Age)
	}
}

/*
	Pluck
	Query satu kolom dari database dan scan menjadi slice, jika Anda ingin query multiple (beberapa) kolom, gunakan Select dengan Scan sebagai gantinya
*/
func Pluck(db *gorm.DB) {
	var names []string
	db.Model(&model.User{}).Pluck("name", &names)
	fmt.Println(names)

	//type APIUser struct {
	//	Name string
	//	Age  uint8
	//}
	//var apiUser []APIUser

	//db.Model(&model.User{}).Select("name", "age").Scan(&apiUser)
	//db.Model(&model.User{}).Select("name", "age").Find(&apiUser)

	//fmt.Println(apiUser)
}

/*
	Scopes
*/
func CheckName(db *gorm.DB) *gorm.DB {
	return db.Where("name = ?", "jinzhu3")
}

func Scopes(db *gorm.DB) {
	var users []model.User

	db.Scopes(CheckName).Find(&users)
	// Find user by Name
	fmt.Println(users)
}

/*
	Count
*/

func Count(db *gorm.DB) {
	//var users []model.User
	var count int64

	db.Model(&model.User{}).Where("name = ?", "jinzhu").Count(&count)
	// SELECT count(1) FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2'

	fmt.Println(count)
}
