package service

import (
	"log"
	"net/http"
	"assignment/model"
	"github.com/gorilla/mux"
	cnts "assignment/constants"
	"strconv"
	"strings"
)

func GetUsers(r *http.Request)([]model.User,model.Error){
	log.Println("Inside GetUsers()")
	req := mux.Vars(r)
	ids_str := req["id"]
	ids:=strings.Split(ids_str,",")
	var i int
	var users []model.User
	for i=0;i<len(ids);i++{
		id,_:=strconv.Atoi(ids[i])
		res,present:=cnts.Users[id]
		if present{
			users=append(users,res)
		}
	}
	if len(users)>0{
		log.Println("Got the datta...")
		log.Println("Completed GetUsers()")
		return users,model.Error{}
	}
	log.Println("Data not found...")
	log.Println("Completed GetUsers()")
	return users,model.Error{Message:"Users not found"}
}