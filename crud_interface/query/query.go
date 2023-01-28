package query

import (
	"fmt"
	"gorm.io/gorm"
	"gorm_test_6/model"
)

/*
	Retrieving a single object using First()
*/
func RetrievingSingleObject1(db *gorm.DB) {
	var user model.User

	// Get the first record ordered by primary key
	//db.First(&user)
	//fmt.Println(user)

	// Get the first record ordered by primary key
	//result := db.First(&user)
	//fmt.Println(user)
	//fmt.Println(result.RowsAffected) // returns count of records found
	//fmt.Println(result.Error)        // returns error or nil

	// check error ErrRecordNotFound
	//result := db.First(&user, 100)
	//checkError := errors.Is(result.Error, gorm.ErrRecordNotFound)
	//fmt.Println(checkError)

	// works because model is specified using `db.Model()`
	result := map[string]interface{}{}
	db.Model(&user).First(&result)
	fmt.Println(user)
}

/*
	Retrieving a single object using Take()
*/
func RetrievingSingleObject2(db *gorm.DB) {
	var user model.User

	// Get one record, no specified order
	db.Take(&user)
	fmt.Println(user)

	// works with Take
	//result := map[string]interface{}{}
	//db.Table("users").Take(&result)
	//fmt.Println(result)
}

/*
	Retrieving a single object using Last()
*/
func RetrievingSingleObject3(db *gorm.DB) {
	var user model.User

	// Get last record, ordered by primary key desc
	db.Last(&user)
	fmt.Println(user)

}

/*
	Retrieving objects with primary key
*/
func RetrievingObjectsWithPrimaryKey(db *gorm.DB) {
	var user model.User
	var users []model.User

	// SELECT * FROM users WHERE id = 10;
	db.First(&user, 10)
	fmt.Println(user)

	// SELECT * FROM users WHERE id = 10;
	//db.First(&user, "10")

	// SELECT * FROM users WHERE id IN (1,2,3);
	db.Find(&users, []int{1, 2, 3})
	fmt.Println(users)

	// If the primary key is a string (for example, like a uuid), the query will be written as follows:
	// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";
	//db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")

	// SELECT * FROM users WHERE id = 3;
	//var user2 = model.User{ID: 3}
	//db.First(&user2)
	//fmt.Println(user2)

	// SELECT * FROM users WHERE id = 3;
	//var result model.User
	//db.Model(model.User{ID: 3}).First(&result)
}

/*
	Retrieving all objects
*/
func RetrievingAllObjects(db *gorm.DB) {
	var users []model.User

	// Get all records
	//SELECT * FROM users;
	result := db.Find(&users)

	fmt.Println(result.RowsAffected) // returns found records count, equals `len(users)`
	fmt.Println(result.Error)        // returns error
}

/*
	Conditions
*/
func ConditionsWhere(db *gorm.DB) {
	var user model.User
	//var users []model.User

	/*
		String Conditions
	*/
	// Get first matched record
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
	db.Where("name = ?", "jinzhu").First(&user)

	// Get all matched records
	// SELECT * FROM users WHERE name <> 'jinzhu';
	//db.Where("name <> ?", "jinzhu").Find(&users)

	// IN
	// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');
	//db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)

	// LIKE
	// SELECT * FROM users WHERE name LIKE '%jin%';
	//db.Where("name LIKE ?", "%jin%").Find(&users)

	// AND
	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;
	//db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)

	/*
		Struct & Map Conditions
		NOTE When querying with struct, GORM will only query with non-zero fields,
		that means if your field’s value is 0, '', false or other zero values, it won’t be used to build query conditions, for example:
	*/
	//result := db.Where(&model.User{Name: "jinzhu", Age: 0}).Find(&users)
	//fmt.Println(users)
	//fmt.Println(result.RowsAffected)

	/*
		Struct & Map Conditions
		To include zero values in the query conditions, you can use a map,
		which will include all key-values as query conditions, for example:
	*/
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;
	//result := db.Where(map[string]interface{}{"name": "jinzhu", "age": 0}).Find(&users)
	//fmt.Println(users)
	//fmt.Println(result.RowsAffected)

	/*
		Specify Struct search fields
	*/
	//result := db.Where(&model.User{Name: "jinzhu"}, "name", "Age").Find(&users)
	//fmt.Println(users)
	//fmt.Println(result.RowsAffected) // returns found records count, equals `len(users)`

}

/*
	Order
	Specify order when retrieving records from the database
*/
func Order(db *gorm.DB) {
	var users []model.User

	result := db.Order("name desc").Find(&users)
	//result := db.Order("name").Find(&users)
	fmt.Println(users)
	fmt.Println(result.RowsAffected) // returns found records count, equals `len(users)`
}

/*
	Limit & Offset
	Limit specify the max number of records to retrieve
	Offset specify the number of records to skip before starting to return the records
*/
func LimitOffset(db *gorm.DB) {
	var users []model.User

	result := db.Limit(3).Offset(7).Find(&users)
	fmt.Println(users)
	fmt.Println(result.RowsAffected) // returns found records count, equals `len(users)`
}

/*
	Group By & Having
*/
func GroupByHaving(db *gorm.DB) {
	type result struct {
		Name  string
		Total int
	}

	var hasil result

	db.Model(&model.User{}).Select("name, sum(age) as total").Where("name = ?", "jinzhu1").Group("name").Find(&hasil)
	fmt.Println(hasil)
}

/*
	Distinct
*/
func Distinct(db *gorm.DB) {
	var users []model.User

	db.Distinct("name", "age").Order("name, age desc").Find(&users)
	fmt.Println(users)
}

/*
	Scan
	Hasil Scan() dimasukkan ke dalam struct bekerja mirip dengan cara ketika kita menggunakan Find()
*/
func Scan(db *gorm.DB) {
	type Result struct {
		Name string
		Age  int
	}

	var result []Result
	db.Table("users").Select("name", "age").Where("name = ?", "jinzhu2").Scan(&result)
	//db.Table("users").Select("name", "age").Where("name = ?", "jinzhu2").Find(&result)
	fmt.Println(result)
}
