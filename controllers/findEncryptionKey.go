package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/clementus360/spacechat-db/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func FindEncryptionKey(UserDB *gorm.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var encryption models.EncryptionKey

		vars := mux.Vars(req)

		result := UserDB.Where("user_id = ?", vars["id"]).First(&encryption)
		if result.Error != nil {
			HandleError(result.Error, "Failed to find user", res)
			return
		}

		json.NewEncoder(res).Encode(encryption)
		// res.WriteHeader(http.StatusOK)
	}
}
