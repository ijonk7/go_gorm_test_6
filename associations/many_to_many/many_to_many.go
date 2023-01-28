package many_to_many

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gorm_test_6/app"
	"gorm_test_6/model"
	"gorm_test_6/model/web"
	"net/http"
)

/*
	Many To Many
*/
func ManyToMany(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var users []model.User6
	//var languages []model.Language

	db := app.NewDB()
	err := db.Model(&model.User6{}).Preload("Languages").Find(&users).Error
	//err := db.Model(&model.Language{}).Preload("Users").Find(&languages).Error

	if err != nil {
		panic(err)
	}

	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   users,
		//Data: languages,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(response)
}
