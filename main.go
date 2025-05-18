package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"restapi.com/app/routes"
	"restapi.com/app/models"
	"restapi.com/app/database"
)


func main(){
	database.Connect()
	database.DB.AutoMigrate(&models.User{})

	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}