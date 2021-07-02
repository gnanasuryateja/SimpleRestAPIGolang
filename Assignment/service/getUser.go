package service

import (
	"github.com/gorilla/mux"
	"net/http"
	cnts "assignment/constants"
	"strconv"
	"assignment/model"
	"log"
)

func GetUser(r *http.Request)(model.User,model.Error){
	log.Println("Inside GetUser()")
	req := mux.Vars(r)
	key := req["id"]
	id,_:=strconv.Atoi(key)
	res,present:=cnts.Users[id]
	if present{
		log.Println("Got the data...")
		log.Println("Completed GetUser()")
		return res,model.Error{}
	}
	log.Println("Data not found...")
	log.Println("Completed GetUser()")
	return res,model.Error{Message:"User not found"}
}