package main

import (
	"assignment/model"
	cnts "assignment/constants"
	"assignment/router"
	"log"
)

func main() {
	log.Println("Intializing DB.....")
	tmp:=make(map[int]model.User)
	tmp[1]=model.User{ID:1,Fname:"Steve",City:"LA",Phone:1234567890,Height:5.8,Married:true}
	tmp[2]=model.User{ID:2,Fname:"Balmer",City:"NY",Phone:9876543210,Height:5.10,Married:true}
	tmp[3]=model.User{ID:3,Fname:"George",City:"NJ",Phone:9012345678,Height:6.2,Married:false}
	cnts.Users=tmp
	log.Println("Starting the server")
    router.Controller()
}