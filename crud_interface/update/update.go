package update

import (
	"fmt"
	"gorm.io/gorm"
	"gorm_test_6/model"
)

/*
	Save All Fields
*/
func SaveAllFields(db *gorm.DB) {
	var user model.User
	db.First(&user)

	user.Name = "jinzhu 22"
	user.Age = 100
	db.Save(&user)

	fmt.Println(user)
}

/*
	Update single column
*/
func UpdateSingleColumn(db *gorm.DB) {
	// Update with conditions
	//db.Model(&model.User{}).Where("age = ?", 10).Update("name", "jinzhu 33")

	//db.Model(model.User{ID: 25}).Update("name", "jinzhu 25")

	// Update with conditions and model value
	db.Model(model.User{ID: 3}).Where("name = ?", "jinzhu 33").Update("age", 34)
}

/*
	Updates multiple columns
*/
func UpdatesMultipleColumns(db *gorm.DB) {
	// Update attributes with `struct`, will only update non-zero fields
	//db.Model(model.User{ID: 7}).Updates(model.User{Name: "hello", Age: 18, Active: false})

	// Update attributes with `map`
	db.Model(model.User{ID: 7}).Updates(map[string]interface{}{"name": "jinzhu 77", "age": 18, "active": false})
}

/*
	Update Selected Fields
*/
func UpdateSelectedFields(db *gorm.DB) {
	// Select with Map
	//db.Model(model.User{ID: 26}).Select("name", "age", "active").Updates(map[string]interface{}{"name": "jinzhu 88", "age": 55, "active": false})

	// Select with Map
	//db.Model(model.User{ID: 27}).Select("name").Updates(map[string]interface{}{"name": "jinzhu 27", "age": 27, "active": false})

	// Select with Struct (select zero value fields)
	db.Model(model.User{ID: 6}).Select("Name", "Age").Updates(model.User{Name: "jinzhu 6", Age: 35})
}

/*
	Update Selected Fields
*/
func TesUpdateHooks(db *gorm.DB) {
	db.Model(&model.User{ID: 1, Name: "jinzhu 1"}).Update("name", "jinzhu 1")

	//db.Model(&model.User{ID: 1, Name: "jinzhu 1"}).Updates(model.User{Name: "jinzhu 1"})
}
