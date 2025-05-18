package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"github.com/gorilla/mux"
	"restapi.com/app/routes"
	"restapi.com/app/models"
	"restapi.com/app/database"
)

func isPortAvailable(port string)bool{
	listener, err := net.Listen("tcp", port)
	if err != nil{
		return false
	}
	_ = listener.Close()
	return true
}

func main(){
	port := ":8080"

	if !isPortAvailable(port){
		log.Fatalf("Port %s is already in use. Please close the conflicting application and run again", port)
	}

	database.Connect()
	database.DB.AutoMigrate(&models.User{})

	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	fmt.Println("Server running at http://localhost%s\n", port)
	if err := http.ListenAndServe(port, router); err != nil{
		log.Fatal(err)
	}
}