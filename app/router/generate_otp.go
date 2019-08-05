package main

import (
	"net/http"
	"fmt"
	// "strings"
	"time"
	database "app/Model/UseDatabase"
	util "app/Utility"
	sms "app/Notification/SMS"
	// "github.com/nu7hatch/gouuid"
	// "net/smtp"
)

const (
	MAX_COUNT = 5
	MAX_HOUR = 4
)

type Verification_Header struct{
	Cookie string `header:"Miti-Cookie"`
}

func generate_otp(w http.ResponseWriter,r *http.Request){
	verification_header:=Verification_Header{}
	util.GetHeader(r,&verification_header)
	session_id:=verification_header.Cookie

	user_id,err:=database.Get_user_id_from_session(session_id)
	if err==""{
		fmt.Println("Session Does not exist")
		util.Message(w,1003)
		return
	}
	verified:= database.IsUserVerified(user_id)
	if !verified{
		_,user_phone:=database.Get_user_detail(user_id)

		if user_phone!=""{
			send_verification_otp(w,user_id,user_phone)
		} else{
			fmt.Println("Mobile no does not exist")
			util.Message(w,1301)
		}
	} else{
		fmt.Println("User already verified")
		util.Message(w,1004)
	}

}

func send_verification_otp(w http.ResponseWriter,id string,phone string){
	count,last_modified:=database.Get_otp_verification_count(id)
	if count<MAX_COUNT{
		otp:="1234"
		database.Enter_verification_otp(id,otp)
		sms.Send_sms(phone,otp)
		fmt.Println("OTP sent")
		util.Message(w,200)
	}

	time_elasped:=time.Since(last_modified)
	if time_elasped.Hours() < MAX_HOUR{
		fmt.Println("otp sent more than limit")
		util.Message(w,1302)
		return
	} else{
		otp:="1234"
		database.Enter_verification_otp(id,otp)
		sms.Send_sms(phone,otp)
		fmt.Println("OTP sent")
		util.Message(w,200)
	}
	
}