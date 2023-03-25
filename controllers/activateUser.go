package controllers

import (
	"net/http"

	"github.com/clementus360/spacechat-db/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func ActivateUser(UserDB *gorm.DB) http.HandlerFunc{
	return func(res http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)

		result := UserDB.Model(&models.User{}).Where("phone_hash = ?", vars["id"]).Update("activated", true)
		if result.Error!=nil {
			HandleError(result.Error, "Failed to activate user", res)
			return
		}

		res.WriteHeader(http.StatusOK)
		res.Write([]byte("User activated successfully"))
	}
}
