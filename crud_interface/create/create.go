package create

import (
	"fmt"
	"gorm.io/gorm"
	"gorm_test_6/model"
	"time"
)

/*
	Create Record
*/
func CreateRecord(db *gorm.DB) {
	layoutFormat := "2006-01-02 15:04:05"
	value := time.Now().Format("2006-01-02 15:04:05")
	date, _ := time.Parse(layoutFormat, value)
	user := model.User{Name: "Jinzhu", Age: 18, Birthday: &date}

	result := db.Create(&user) // pass pointer of data to Create

	fmt.Println(user.ID)             // returns inserted data's primary key
	fmt.Println(result.Error)        // returns error
	fmt.Println(result.RowsAffected) // returns inserted records count
}

/*
	Create Record With Selected Fields
*/
func CreateRecord2(db *gorm.DB) {
	layoutFormat := "2006-01-02 15:04:05"
	value := time.Now().Format("2006-01-02 15:04:05")
	date, _ := time.Parse(layoutFormat, value)
	user := &model.User{Name: "Jinzhu 2", Age: 10, Birthday: &date}
	result := db.Select("Name", "Age", "Birthday").Create(&user)

	fmt.Println(user.ID)             // returns inserted data's primary key
	fmt.Println(result.Error)        // returns error
	fmt.Println(result.RowsAffected) // returns inserted records count
}

/*
	Batch Insert
*/
func BatchInsert(db *gorm.DB) {
	var users = []model.User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	db.Create(&users)

	for _, user := range users {
		fmt.Println(user.ID) // 1,2,3
	}
}

/*
	Batch Insert with size
*/
func BatchSize(db *gorm.DB) {
	var users = []model.User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	db.CreateInBatches(&users, 1)

	for _, user := range users {
		fmt.Println(user.ID) // 1,2,3
	}
}

/*
	Create From Map
*/
func CreateFromMap1(db *gorm.DB) {
	result := db.Model(model.User{}).Create(map[string]interface{}{
		"Name": "jinzhu", "Age": 18,
	})

	fmt.Println(result.Error)        // returns error
	fmt.Println(result.RowsAffected) // returns inserted records count
}

/*
	Create From Map (NOT WORKING)
	batch insert from `[]map[string]interface{}{}`
*/
func CreateFromMap2(db *gorm.DB) {
	db.Model(&model.User{}).Create([]map[string]interface{}{
		{"Name": "jinzhu_1", "Age": 14},
		{"Name": "jinzhu_2", "Age": 15},
	})
}

/*
	Create With Associations
*/
func CreateWithAssociations(db *gorm.DB) {
	db.Create(&model.User{
		Name:       "jinzhu",
		CreditCard: model.CreditCard{Number: "411111111112"},
	})
}
