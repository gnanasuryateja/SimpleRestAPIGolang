package testing

import (
	cnts "assignment/constants"
	"assignment/model"
	"net/http"
	"testing"
	"io/ioutil"
	"encoding/json"
)

func TestGetUser1(t *testing.T){
	tmp:=make(map[int]model.User)
	tmp[1]=model.User{ID:1,Fname:"Steve",City:"LA",Phone:1234567890,Height:5.8,Married:true}
	tmp[2]=model.User{ID:2,Fname:"Balmer",City:"NY",Phone:9876543210,Height:5.10,Married:true}
	tmp[3]=model.User{ID:3,Fname:"George",City:"NJ",Phone:9012345678,Height:6.2,Married:false}
	cnts.Users=tmp
	url:="http://localhost:8080/getUser/1"
	req,_:=http.NewRequest("GET",url,nil)
	client:=&http.Client{}
	response,_:=client.Do(req)
	bytesVal,_:=ioutil.ReadAll(response.Body)
	var res model.User
	json.Unmarshal(bytesVal,&res)
	if res.ID!=1{
		t.Errorf("Failed during TestGetUser1")
	}
}

func TestGetUser2(t *testing.T){
	tmp:=make(map[int]model.User)
	tmp[1]=model.User{ID:1,Fname:"Steve",City:"LA",Phone:1234567890,Height:5.8,Married:true}
	tmp[2]=model.User{ID:2,Fname:"Balmer",City:"NY",Phone:9876543210,Height:5.10,Married:true}
	tmp[3]=model.User{ID:3,Fname:"George",City:"NJ",Phone:9012345678,Height:6.2,Married:false}
	cnts.Users=tmp
	url:="http://localhost:8080/getUser/4"
	req,_:=http.NewRequest("GET",url,nil)
	client:=&http.Client{}
	response,_:=client.Do(req)
	bytesVal,_:=ioutil.ReadAll(response.Body)
	var res model.User
	json.Unmarshal(bytesVal,&res)
	if res.ID!=0{
		t.Errorf("Failed during TestGetUser2")
	}
}

func TestGetUsers1(t *testing.T){
	tmp:=make(map[int]model.User)
	tmp[1]=model.User{ID:1,Fname:"Steve",City:"LA",Phone:1234567890,Height:5.8,Married:true}
	tmp[2]=model.User{ID:2,Fname:"Balmer",City:"NY",Phone:9876543210,Height:5.10,Married:true}
	tmp[3]=model.User{ID:3,Fname:"George",City:"NJ",Phone:9012345678,Height:6.2,Married:false}
	cnts.Users=tmp
	url:="http://localhost:8080/getUsers/1,2,3"
	req,_:=http.NewRequest("GET",url,nil)
	client:=&http.Client{}
	response,_:=client.Do(req)
	bytesVal,_:=ioutil.ReadAll(response.Body)
	var res []model.User
	json.Unmarshal(bytesVal,&res)
	if len(res)!=3{
		t.Errorf("Failed during TestGetUsers1")
	}
}

func TestGetUsers2(t *testing.T){
	tmp:=make(map[int]model.User)
	tmp[1]=model.User{ID:1,Fname:"Steve",City:"LA",Phone:1234567890,Height:5.8,Married:true}
	tmp[2]=model.User{ID:2,Fname:"Balmer",City:"NY",Phone:9876543210,Height:5.10,Married:true}
	tmp[3]=model.User{ID:3,Fname:"George",City:"NJ",Phone:9012345678,Height:6.2,Married:false}
	cnts.Users=tmp
	url:="http://localhost:8080/getUsers/4"
	req,_:=http.NewRequest("GET",url,nil)
	client:=&http.Client{}
	response,_:=client.Do(req)
	bytesVal,_:=ioutil.ReadAll(response.Body)
	var res []model.User
	json.Unmarshal(bytesVal,&res)
	if len(res)!=0{
		t.Errorf("Failed during TestGetUsers2")
	}
}