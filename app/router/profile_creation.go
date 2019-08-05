package main

import(
	"fmt"
	"net/http"
	// "log"
	"io/ioutil"
	// "strings"
	"encoding/json"
   CD "app/Model/CreateDatabase"
   database "app/Model/UseDatabase"
   util "app/Utility"
)

type Profile_creation_Header struct{
	Method1 string `header:"method"`
	Agent1 string `header:"agent"`
	Cookie string `header:"Miti-Cookie"`
}

func profile_creation(w http.ResponseWriter, r *http.Request){
	header:=Profile_creation_Header{}
	util.GetHeader(r,&header)


	session_id:=header.Cookie

	user_id,d_err:=database.Get_user_id_from_session(session_id)
	if d_err==""{
		fmt.Println("Session Does not exist")
		util.Message(w,1003)
		return
	}


	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}
	profile_data:=CD.Profile{}
	err_profile_data:=json.Unmarshal(requestBody,&profile_data)
	if err_profile_data!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}
	profile_data.User_id=user_id
	sanatization_status:=CD.Sanatize(profile_data)
	if sanatization_status =="ERROR"{
		fmt.Println("profile creation data invalid")
		util.Message(w,1002)
		return
	}
	profile_data_handle(w,profile_data)

}

func profile_data_handle(w http.ResponseWriter,profile_data CD.Profile){
	database.Enter_profile_data(profile_data)
	util.Message(w,200)
	return
}