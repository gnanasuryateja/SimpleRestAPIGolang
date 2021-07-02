package router

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"assignment/service"
	"encoding/json"
)

func Controller() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/getUser/{id}", getUserHandler).Methods("GET")
	myRouter.HandleFunc("/getUsers/{id}", getUsersHandler).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func getUserHandler(w http.ResponseWriter, r *http.Request){
	log.Println("API to get single user...")
	response,err:=service.GetUser(r)
	w.Header().Set("content-type", "application/json")
	if err.Message!=""{
		log.Println("Completed getUser API...")
		json.NewEncoder(w).Encode(err)
		return
	}
	log.Println("Completed getUser API...")
	json.NewEncoder(w).Encode(response)
	return
}

func getUsersHandler(w http.ResponseWriter, r *http.Request){
	log.Println("API to get multiple users")
	response,err:=service.GetUsers(r)
	w.Header().Set("content-type", "application/json")
	if err.Message!=""{
		log.Println("Completed getUsers API...")
		json.NewEncoder(w).Encode(err)
		return
	}
	log.Println("Completed getUsers API...")
	json.NewEncoder(w).Encode(response)
	return
}