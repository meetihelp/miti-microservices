package main

import (
	"net/http"
	"fmt"
	// "strings"
	"time"
	database "app/Model/UseDatabase"
	util "app/Utility"
	mail "app/Notification/Mail"
	"github.com/nu7hatch/gouuid"
	// "net/smtp"
)


type Email_Verification_Header struct{
	Cookie string `header:"Miti-Cookie"`
}

func generate_verification_email(w http.ResponseWriter,r *http.Request){
	email_verification_header:=Email_Verification_Header{}
	util.GetHeader(r,&email_verification_header)
	session_id:=email_verification_header.Cookie

	user_id,err:=database.Get_user_id_from_session(session_id)
	if err==""{
		fmt.Println("Session Does not exist")
		util.Message(w,1003)
		return
	}
	verified:= database.IsUserVerified(user_id)
	if !verified{
		user_email,_:=database.Get_user_detail(user_id)

		if user_email!=""{
			send_verification_email(w,user_id,user_email)
		} else{
			fmt.Println("Email id does not exist")
			util.Message(w,1201)
		}
	} else{
		fmt.Println("User already verified")
		util.Message(w,1004)
	}

}

func send_verification_email(w http.ResponseWriter,id string,email string){
	count,last_modified:=database.Get_Email_verification_count(id)
	if count < MAX_COUNT{
		u, _ := uuid.NewV4()
		token:=u.String()
		url:="http://localhost:9000/verify_email?token="+token
		database.Enter_email_verification(id,token)
		mail.Send_email(email,url)
		fmt.Println("Email sent for verification")
		util.Message(w,200)
		return
	}

	time_elasped:=time.Since(last_modified)
	if time_elasped.Hours() < MAX_HOUR{
		fmt.Println("Link sent more than limit")
		util.Message(w,1202)
		return
	} else{
		database.Delete_all_email_verification(id)
		u, _ := uuid.NewV4()
		token:=u.String()
		url:="http://localhost:9000/verify_email?token="+token
		database.Enter_email_verification(id,token)
		mail.Send_email(email,url)
		fmt.Println("Email sent for verification")
		util.Message(w,200)
		return
	}
}