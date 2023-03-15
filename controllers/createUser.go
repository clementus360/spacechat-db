package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/clementus360/spacechat-db/models"
	"gorm.io/gorm"
)

type CreateUserRequestData struct {
	User models.User `json:"user"`
	Key string `json:"key"`
}

// Handles Error logs in the console and the response
func HandleError(err error,message string, res http.ResponseWriter) {
	fmt.Println(message)
	fmt.Println(err)
	http.Error(res, message, http.StatusInternalServerError)
}

// Adds user to the database
func CreateUser(UserDB *gorm.DB) http.HandlerFunc{
	return func(res http.ResponseWriter, req *http.Request) {
		var body CreateUserRequestData
		var user models.User
		var EncryptionKey models.EncryptionKey

		// Gets the user data from request body
		err := json.NewDecoder(req.Body).Decode(&body)
		if err!=nil {
			HandleError(err, "Failed to decode request body", res)
			return
		}

		user = body.User

		// Adds a new user to the database
		if err := UserDB.Create(&user).Error; err != nil {
			HandleError(err, "Failed to add user to DB", res)
			return
		}

		EncryptionKey = models.EncryptionKey{
			UserID: user.ID,
			Key: body.Key,
		}

		// Adds a new encryption key to the database
		if err := UserDB.Create(&EncryptionKey).Error; err != nil {
			HandleError(err, "Failed to add encryption key to DB", res)
			return
		}

		// Writes a response when successful
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("User created successfully"))
	}
}
