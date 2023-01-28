package preloading

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm_test_6/app"
	"gorm_test_6/model"
	"gorm_test_6/model/web"
	"net/http"
)

/*
	Preload
*/
func Preload(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	db := app.NewDB()

	var users []model.User5
	db.Preload("CreditCards").Find(&users)

	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   users,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(response)
}

/*
	Preload All
*/
func PreloadAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	db := app.NewDB()

	var users []model.User5
	db.Preload(clause.Associations).Find(&users)

	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   users,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(response)
}

/*
	Preload with conditions
*/
func PreloadWithConditions(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	db := app.NewDB()

	var users []model.User5
	db.Preload("CreditCards", "number = ?", "5555555").Find(&users)

	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   users,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(response)
}

/*
	Custom Preloading SQL
*/
func CustomPreloadingSql(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	db := app.NewDB()

	var users []model.User5
	db.Preload("CreditCards", func(db *gorm.DB) *gorm.DB {
		return db.Order("credit_card5.number DESC")
	}).Find(&users)

	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   users,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(response)
}
