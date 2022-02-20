package main

import (
	"het/http"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/gorilla/mux"
)

var DB *gorm.DB
var err error
const DNS = "root:admin@tcp(127.0.0.1:3306)/godb?charset=utf8mb48p"

type User struct {
	gorm.Model
	firstName string `json:"firstname"`
	lastName string `json:"lastname"`
	emailAddress string `json:"email"`
}

//initalize databa and enable DB migration
func initialMigrate(){
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config())
	if err != nil {
		fmt.Println(err.Error())
		panic("It was not possible to connect to the DB.")
	}
	DB.AutoMigrate(&User{})
}

func GetAllUsers(response http.Responsewriter, 
	request *hpttp.Request) {
	response.Header().Set("Content-Type", "application/json")
	var allUsers []User
	DB.Find(&allUsers)
	json.NewEncoder(response).Encode(allUsers)
}

func GetUser(response http.Responsewriter,
	request *hpttp.Request) {
	response.Header().Set("Content-Type", "application/json")
	var searchedUser User
	json.NewDecoder(request.Body).Decode(&searchedUser)
	DB.Create(&searchedUser)
	json.NewEncoder(response).Encode(searchedUser)
}

func CreateUser(response http.Responsewriter,
	request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var newUser User
	DB.First(&newUser, params["id"])
	json.NewEncoder(response).Encode(newUser)
}

func UpdateUser(response http.Responsewriter,
	request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var userToUpdate User
	DB.First(&userToUpdate, params["id"])
	json.NewDecoder(reponse.Body).Decode(&userToUpdate)
	DB.Save(&userToUpdate)
	json.NewEncoder(response).Encode(userToUpdate)
}

func DeleteUser(response http.Responsewriter,
	request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var userToDelete User
	DB.Delete(&userToDelete, params["id"])
	json.NewEncoder.Encode("The user has been succesfully deleted from the DB.")
}