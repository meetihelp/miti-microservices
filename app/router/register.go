package main

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	// "io"
	"encoding/json"
   CD "app/Model/CreateDatabase"
   database "app/Model/UseDatabase"
   util "app/Utility"
)

func register(w http.ResponseWriter, r *http.Request){
	util.GetHeader(r)

	//Read body data
	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,100)
		return 
	}

	//UNMARSHILING DATA
	user_data :=CD.User{}
	err_user_data:=json.Unmarshal(requestBody,&user_data)
	if err_user_data!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,101)
		return 
	}

	//SANITIZE USER AND PROFILE DATA
	sanatization_status :=CD.Sanatize(user_data)
	if sanatization_status =="ERROR"{
		fmt.Println("User data invalid")
		util.Message(w,102)
		return
	}
	user_data_handle(w,user_data)
}

func user_data_handle(w http.ResponseWriter, user_data CD.User){
	_,db_status:=database.Enter_user_data(user_data)
	if db_status ==1{
		log.Println("User Already exist")
		util.Message(w,103)
		return
	} else{
		log.Println("User data entered successfully")
		util.Message(w,104)
		return
	}
}

