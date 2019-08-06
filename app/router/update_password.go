package main 

import(
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	database "app/Model/UseDatabase"
	CD "app/Model/CreateDatabase"
    util "app/Utility"
)

type Update_password_header struct{
	Cookie string `header:"Miti-Cookie"`
}


func update_password(w http.ResponseWriter,r *http.Request){
	update_password_header:=Update_password_header{}
	util.GetHeader(r,&update_password_header)
	session_id:=update_password_header.Cookie

	user_id,status:=database.Get_user_id_from_session(session_id)
	if status!="OK"{
		fmt.Println("Session does not exist")
		util.Message(w,1003)
		return
	}

	//Read body data
	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}

	//UNMARSHILING DATA
	password_change_data :=CD.Password_change{}
	err_user_data:=json.Unmarshal(requestBody,&password_change_data)
	if err_user_data!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}

	sanatization_status :=CD.Sanatize(password_change_data)
	if sanatization_status =="ERROR"{
		fmt.Println("User data invalid")
		util.Message(w,1002)
		return
	}

	status=database.Check_user_by_id(user_id,password_change_data.Old_Password)

	if status=="OK"{
		//UPDATE PASSWORD
		database.Update_Password(user_id,password_change_data.New_Password)
		util.Message(w,200)
	} else{
		//SEND ERROR MESSAGE

	}
}