package has_one

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gorm_test_6/app"
	"gorm_test_6/model"
	"gorm_test_6/model/web"
	"net/http"
)

/*
	Has One
*/
func HasOne(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var users []model.User4

	db := app.NewDB()
	err := db.Model(&model.User4{}).Preload("CreditCard").Find(&users).Error

	if err != nil {
		panic(err)
	}

	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   users,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(response)
}
