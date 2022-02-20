package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func initializeRouter() {
	router := mux.NewRouter()
	
	router.HandleFunc("/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.Handlefunc("/users/{id}", UpdateUser).Methods("PUT")
	router.Handlefunc("/users/{id}", DeleteUser).Methods("PUT")

	log.Fatal(http.ListenAndServer(":9000", router))
}

func main(){
	initialMigrate()
	initializeRouter()
}