package main

import(
	"fmt"
	"net/http"
	// "log"
	"io/ioutil"
	"encoding/json"
   CD "app/Model/CreateDatabase"
   database "app/Model/UseDatabase"
   util "app/Utility"
)

func profile_creation(w http.ResponseWriter, r *http.Request){
	util.GetHeader(r)
	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,200)
		return 
	}
	profile_data:=CD.Profile{}
	err_profile_data:=json.Unmarshal(requestBody,&profile_data)
	if err_profile_data!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,201)
		return
	}
	profile_data.User_id="1234"
	sanatization_status:=CD.Sanatize(profile_data)
	if sanatization_status =="ERROR"{
		fmt.Println("profile creation data invalid")
		util.Message(w,202)
		return
	}
	profile_data_handle(w,profile_data)

}

func profile_data_handle(w http.ResponseWriter,profile_data CD.Profile){
	database.Enter_profile_data(profile_data)
	util.Message(w,203)
	return
}