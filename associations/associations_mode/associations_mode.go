package associations_mode

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gorm_test_6/app"
	"gorm_test_6/model"
	"gorm_test_6/model/web"
	"net/http"
)

/*
	Associations - Auto Create/Update
	Associations - Skip Auto Create/Update
*/
func AutoCreateUpdate(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	db := app.NewDB()

	user := model.User5{
		Name: "Fika",
		CreditCards: []model.CreditCard5{
			{Number: "6666666"},
			{Number: "7777777"},
		},
	}

	//db.Select("name").Where("name = ?", "Deni").Create(&user)
	//db.Where("name = ?", "Fika").Create(&user)
	//db.Create(&user)
	db.Select("CreditCards").Create(&user)
}

/*
	Association Mode
*/
func AssociationMode(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	db := app.NewDB()

	var user model.User5
	result := db.Model(&user).Association("CreditCards")
	fmt.Println(result)
	fmt.Println(user)

	var user2 model.User5
	err := db.Model(&user2).Association("CreditCards").Error
	if err != nil {
		panic(err)
	}
}

/*
	Find Associations
*/
func FindAssociations(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	db := app.NewDB()

	//var user model.User5
	user := model.User5{
		ID: 4,
		//Name: "Deni",
		//Name: "Fika",
		Name: "Dodi",
		//CreditCards: []model.CreditCard5{
		//	{Number: "7777777", User5ID: 5},
		//	{Number: "4444444", User5ID: 8},
		//},
		//CreditCards: []model.CreditCard5{{Number: "7777777"}, {Number: "4444444"}},
		//CreditCards: []model.CreditCard5{},
	}
	//var creditCard5 model.CreditCard5

	//db.Model(&user).Association("CreditCards")
	//db.Model(&user).Association("CreditCards").Find(&creditCard5)
	//number := []string{"7777777", "4444444"}
	//db.Model(&user).Where("number IN ?", number).Association("CreditCards").Find(&creditCard5)
	//name := []string{"Dika", "Fika"}
	//db.Model(&user).Where("name IN ?", name).Association("CreditCards").Find(&creditCard5)
	//db.Model(&user).Association("CreditCards").Count()
	db.Model(&user).Association("CreditCards").Delete(model.CreditCard5{Number: "111111111"})

	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(response)
}

/*
	Delete with Select
*/
func DeleteWithSelect(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	db := app.NewDB()

	user := model.User5{ID: 2}
	db.Select("CreditCards").Delete(&user)

	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(response)
}
