package main 

import (
	"net/http"
	"fmt"
	CD "app/Model/CreateDatabase"
	database "app/Model/UseDatabase"
	util "app/Utility"
	"io/ioutil"
	"encoding/json"
	"strings"
)

type Login_header struct{
	Cookie string `header:"Cookie"`
}
func login(w http.ResponseWriter,r *http.Request){
	ip_address:=util.Get_IP_address(r)

	login_header:=Login_header{}
	util.GetHeader(r,&login_header)
	//WRITE CODE IF COOKIE EXIST
	session_id:=login_header.Cookie
	x:=strings.Split(session_id,";")
	x=strings.Split(x[1],"=")
	session_id=x[1]
	user_id,login_status:=database.Get_user_id_from_session(session_id)
	if login_status=="OK"{
		util.Message(w,405)
		return
	}

	//Read body data
	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,400)
		return 
	}

	//UNMARSHILING DATA
	user_data :=CD.User{}
	err_user_data:=json.Unmarshal(requestBody,&user_data)
	if err_user_data!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,401)
		return 
	}

	sanatization_status :=CD.Sanatize(user_data)
	if sanatization_status =="ERROR"{
		fmt.Println("User data invalid")
		util.Message(w,402)
		return
	}

	user_id,login_status=database.Check_user(user_data)

	if login_status=="WRONG_PASSWORD"{
		util.Message(w,403)
		return
	}
	if login_status=="NO_USER"{
		util.Message(w,403)
		return
	}
	if login_status=="UNVERIFIED"{
		cookie:=database.Insert_session(user_id,ip_address)
		http.SetCookie(w,&cookie)
		util.Message(w,404)
		return
	} 
	if login_status=="OK"{
		cookie:=database.Insert_session(user_id,ip_address)
		http.SetCookie(w,&cookie)
		util.Message(w,405)
		return
	}
}


// func Check_user_already_login(session_id string) (string,string){
// 	user_id,err:=database.Get_user_id_from_session(session_id)
// 	if err==""{
// 		fmt.Println("Session Does not exist")
// 		return user_id."ERROR"
// 	}


// }