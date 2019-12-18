package Authentication
import (
	"net/http"
	"fmt"
	util "miti-microservices/Util"
	mail "miti-microservices/Notification/Mail"
)

const(
	domain="http://localhost:9000"
)

type EmailVerificationHeader struct{
	Cookie string `header:"MitiCookie"`
}

func GenerateVerificationEmail(w http.ResponseWriter,r *http.Request){
	emailVerificationHeader:=EmailVerificationHeader{}
	util.GetHeader(r,&emailVerificationHeader)
	sessionId:=emailVerificationHeader.Cookie

	userId,err:=util.GetUserIdFromSession(sessionId)
	if err=="Error"{
		fmt.Println("Session Does not exist")
		util.Message(w,1003)
		return
	}
	verified:= IsUserVerified(userId)
	if !verified{
		userEmail,_:=GetUserDetail(userId)

		if userEmail!=""{
			sendVerificationEmail(w,userId,userEmail)
		} else{
			fmt.Println("Email id does not exist")
			util.Message(w,1201)
		}
	} else{
		fmt.Println("User already verified")
		util.Message(w,1004)
	}

}

func sendVerificationEmail(w http.ResponseWriter,id string,email string){
	count,lastModified:=GetEmailVerificationCount(id)
	if count < MAXCOUNT{
		token:=util.GenerateToken()
		url:=domain+"/verify_email?token="+token
		EnterEmailVerification(id,token)
		mail.SendEmail(email,url)
		fmt.Println("Email sent for verification")
		util.Message(w,200)
		return
	}
	fmt.Println(lastModified)
	// timeElasped:=time.Since(lastModified)
	// if timeElasped.Hours() < MAXHOUR{
	// 	fmt.Println("Link sent more than limit")
	// 	util.Message(w,1202)
	// 	return
	// } else{
	// 	DeleteAllEmailVerification(id)
	// 	token:=util.GenerateToken()
	// 	url:=domain+"/verify_email?token="+token
	// 	EnterEmailVerification(id,token)
	// 	mail.SendEmail(email,url)
	// 	fmt.Println("Email sent for verification")
	// 	util.Message(w,200)
	// 	return
	// }
}