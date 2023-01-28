package belongs_to

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gorm_test_6/app"
	"gorm_test_6/model"
	"gorm_test_6/model/web"
	"net/http"
)

/*
	Belongs To
*/
func BelongsTo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//var user model.User
	var users []model.User3

	db := app.NewDB()
	db.Joins("Company").Find(&users)

	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   users,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(response)
}
