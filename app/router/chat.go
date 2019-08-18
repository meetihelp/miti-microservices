package main 

import (
	"net/http"
	"fmt"
	CD "app/Model/CreateDatabase"
	database "app/Model/UseDatabase"
	util "app/Utility"
	"io/ioutil"
	"encoding/json"
	"time"
)

type Chat_header struct{
	Cookie string `header:"Miti-Cookie"`
}




func chat(w http.ResponseWriter,r *http.Request){
	chat_header:=Chat_header{}
	util.GetHeader(r,&chat_header)
	session_id:=chat_header.Cookie

	_,login_status:=database.Get_user_id_from_session(session_id)

	if login_status=="ERROR"{
		util.Message(w,1003)
		return
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}

	chat_data :=CD.Chat{}
	err_user_data:=json.Unmarshal(requestBody,&chat_data)
	if err_user_data!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}
	chat_data.Message_id=util.Generate_token()
	chat_data.CreatedAt=time.Now()
	db:=database.GetDB()
	db.Create(&chat_data)
}
