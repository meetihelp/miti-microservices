package Authentication
import (
	"net/http"
	"fmt"	
	util "app/Util"
	sms "app/Notification/SMS")

const (
	MAXCOUNT = 5
	MAXHOUR = 4
)



func GenerateOtp(w http.ResponseWriter,r *http.Request){
	verificationHeader:=LoginHeader{}
	util.GetHeader(r,&verificationHeader)
	sessionId:=verificationHeader.Cookie
	userId,err:=util.GetUserIdFromUserVerificationSession(sessionId)
	if err=="Error"{
		fmt.Println("Session Does not exist")
		util.Message(w,1003)
		return
	}
	verified:= IsUserVerified(userId)
	if !verified{
		_,userPhone:=GetUserDetail(userId)

		if userPhone!=""{
			sendVerificationOtp(w,userId,userPhone)
		} else{
			fmt.Println("Mobile no does not exist")
			util.Message(w,1301)
		}
	} else{
		fmt.Println("User already verified")
		util.Message(w,1004)
	}

}

func sendVerificationOtp(w http.ResponseWriter,id string,phone string){
	count,lastModified:=GetOtpVerificationCount(id)
	if count<MAXCOUNT{
		otp:=util.GenerateOtpString()
		EnterVerificationOtp(id,otp)
		sms.SendSMS(phone,otp)
		fmt.Println("OTP sent")
		util.Message(w,200)
		return
	}

	// time_elasped:=time.Since(last_modified)
	// if time_elasped.Hours() < MAX_HOUR{
	// 	fmt.Println("otp sent more than limit")
	// 	util.Message(w,1302)
	// 	return
	// } else{
	// 	otp:=util.Generate_otp_string()
	// 	database.Enter_verification_otp(id,otp)
	// 	sms.Send_sms(phone,otp)
	// 	fmt.Println("OTP sent")
	// 	util.Message(w,200)
	// }
	fmt.Println(lastModified)
	fmt.Println("OTP sent")
		util.Message(w,200)
	
}