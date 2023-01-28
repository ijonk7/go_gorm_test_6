package serializer

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gorm_test_6/app"
	"gorm_test_6/model/web"
	"net/http"
)

/*
	Serializer
*/
func Serializer(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	type Roles []string

	type Job struct {
		Title    string
		Location string
		IsIntern bool
	}

	type User11 struct {
		ID          uint                   `gorm:"primaryKey"`
		Name        []byte                 `gorm:"serializer:json"`
		Roles       Roles                  `gorm:"serializer:json"`
		Contracts   map[string]interface{} `gorm:"serializer:json"`
		JobInfo     Job                    `gorm:"type:bytes;serializer:gob"`
		CreatedTime int64                  `gorm:"serializer:unixtime;type:time"` // store int as datetime into database
	}

	db := app.NewDB()

	//db.AutoMigrate(Job{})
	//db.AutoMigrate(User11{})

	//createdAt := time.Date(2020, 1, 1, 0, 8, 0, 0, time.UTC)
	//data := User11{
	//	ID:          1,
	//	Name:        []byte("jinzhu"),
	//	Roles:       []string{"admin", "owner"},
	//	Contracts:   map[string]interface{}{"name": "jinzhu", "age": 10},
	//	CreatedTime: createdAt.Unix(),
	//	JobInfo: Job{
	//		Title:    "Developer",
	//		Location: "NY",
	//		IsIntern: false,
	//	},
	//}
	//db.Create(&data)

	//var result User11
	//db.First(&result, "id = ?", 1)
	//fmt.Println(result)

	var result User11
	db.Where(User11{Name: []byte("jinzhu")}).Take(&result)
	fmt.Println(result)

	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(response)
}
