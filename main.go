package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/clementus360/spacechat-db/config"
	"github.com/clementus360/spacechat-db/controllers"
	"github.com/clementus360/spacechat-db/models"
	"github.com/clementus360/spacechat-db/services"
	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()

	UserDB := config.ConnectDB()

	config.AutoMigrate(UserDB, &models.User{})
	config.AutoMigrate(UserDB, &models.EncryptionKey{})

	router := mux.NewRouter()
	router.HandleFunc("/api/user", controllers.CreateUser(UserDB)).Methods("POST")
	router.HandleFunc("/api/user/{id}", controllers.FindUser(UserDB)).Methods("GET")
	router.HandleFunc("/api/encryption/{id}", controllers.FindEncryptionKey(UserDB)).Methods("GET")
	router.HandleFunc("/api/user/{id}", controllers.ActivateUser(UserDB)).Methods("PUT")

	go func() {
		services.DeleteUsersHandler(UserDB)
	}()

	err := http.ListenAndServe(":3001", router)
	if err != nil {
		fmt.Println("Failed to start server")
		log.Fatal(err)
	} else {
		fmt.Println("Server runnning on port 5000")
	}
}
