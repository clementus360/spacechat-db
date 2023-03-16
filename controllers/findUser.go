package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/clementus360/spacechat-db/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func FindUser(UserDB *gorm.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var user models.User

		vars := mux.Vars(req)

		fmt.Println(vars["id"])

		result := UserDB.Where("phone_hash = ?", vars["id"]).First(&user)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				HandleError(result.Error, "Failed to find user", res)
			}
			return
		}

		json.NewEncoder(res).Encode(user)
		res.WriteHeader(http.StatusOK)
	}
}
