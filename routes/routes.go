package routes

import(
	"github.com/gorilla/mux"
	"restapi.com/app/controllers"
)

func RegisterRoutes(router *mux.Router){
	router.HandleFunc("/api/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/users/{id}", controllers.UpdateUser).Methods("PUT")
}