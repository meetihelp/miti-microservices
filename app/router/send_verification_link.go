package main

import (
	"net/http"
	"fmt"
	"strings"
	database "app/Model/UseDatabase"
	util "app/Utility"
	mail "app/Notification/Mail"
	"github.com/nu7hatch/gouuid"
	// "net/smtp"
)

type Verification_Header struct{
	Cookie string `header:"Cookie"`
}
func send_verification_link(w http.ResponseWriter,r *http.Request){
	verification_header:=Verification_Header{}
	util.GetHeader(r,&verification_header)
	session_id:=verification_header.Cookie
	x:=strings.Split(session_id,";")
	x=strings.Split(x[1],"=")
	session_id=x[1]
	fmt.Println(session_id)
	user_id,err:=database.Get_user_id_from_session(session_id)
	if err==""{
		util.Message(w,200)
		return
	}
	fmt.Println("Session query"+user_id)
	user_email,user_phone:=database.Get_user_detail(user_id)
	if user_email !=""{
		send_verification_email(w,user_id,user_email)
	} else{
		send_verification_otp(w,user_id,user_phone)
	}
}

func send_verification_email(w http.ResponseWriter,id string,email string){
	u, _ := uuid.NewV4()
	token:=u.String()
	url:="http://localhost:9000/verify_email?token="+token
	database.Enter_verification_email(id,token)
	mail.Send_email(email,url)
	util.Message(w,201)
}	

func send_verification_otp(w http.ResponseWriter,id string,email string){
	util.Message(w,202)
}