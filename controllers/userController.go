package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "restapi.com/app/models"
    "restapi.com/app/database"
)

var users []models.User

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
    var users []models.User
    database.DB.Find(&users)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)
    database.DB.Create(&user)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var user models.User
    database.DB.Delete(&user, params["id"])
    w.WriteHeader(http.StatusNoContent)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var user models.User
    database.DB.First(&user, params["id"])

    if user.ID == 0{
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode("User not found")
        return
    }

    json.NewDecoder(r.Body).Decode(&user)
    database.DB.Save(&user)
    json.NewEncoder(w).Encode(user)
}