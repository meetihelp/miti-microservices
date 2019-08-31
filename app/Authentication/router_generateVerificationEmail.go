package Authentication
import (
	"net/http"
	"fmt"
	// "strings"
	"time"
	util "app/Util"
	mail "app/Notification/Mail"
	// "net/smtp"
)

const(
	domain="http://localhost:9000"
)

type Email_Verification_Header struct{
	Cookie string `header:"Miti-Cookie"`
}

func generate_verification_email(w http.ResponseWriter,r *http.Request){
	email_verification_header:=Email_Verification_Header{}
	util.GetHeader(r,&email_verification_header)
	session_id:=email_verification_header.Cookie

	user_id,err:=util.Get_user_id_from_session(session_id)
	if err=="ERROR"{
		fmt.Println("Session Does not exist")
		util.Message(w,1003)
		return
	}
	verified:= IsUserVerified(user_id)
	if !verified{
		user_email,_:=Get_user_detail(user_id)

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
	count,last_modified:=Get_Email_verification_count(id)
	if count < MAX_COUNT{
		token:=util.Generate_token()
		url:=domain+"/verify_email?token="+token
		Enter_email_verification(id,token)
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
		Delete_all_email_verification(id)
		token:=util.Generate_token()
		url:=domain+"/verify_email?token="+token
		Enter_email_verification(id,token)
		mail.Send_email(email,url)
		fmt.Println("Email sent for verification")
		util.Message(w,200)
		return
	}
}