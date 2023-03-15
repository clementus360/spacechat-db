package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/clementus360/spacechat-db/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func FindUser(UserDB *gorm.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var user models.User

		vars := mux.Vars(req)

		result := UserDB.Where("phone_hash = ?", vars["id"]).First(&user)
		if result.Error != nil {
			HandleError(result.Error, "Failed to find user", res)
			return
		}

		json.NewEncoder(res).Encode(user)
	}
}
