package Authentication
// import (
// 	"net/http"
// 	"fmt"	
// 	"time"
// 	util "miti-microservices/Util"
// 	sms "miti-microservices/Notification/SMS")

const (
	MAXCOUNT = 5
	MAXHOUR = 1
)



// func GenerateOtp(w http.ResponseWriter,r *http.Request){
// 	verificationHeader:=LoginHeader{}
// 	util.GetHeader(r,&verificationHeader)
// 	sessionId:=verificationHeader.Cookie
// 	userId,err:=util.GetUserIdFromUnverifiedUserSession(sessionId)
// 	if err=="Error"{
// 		fmt.Println("Session Does not exist")
// 		util.Message(w,1003)
// 		return
// 	}
// 	verified:= IsUserVerified(userId)
// 	if !verified{
// 		_,userPhone:=GetUserDetail(userId)

// 		if userPhone!=""{
// 			sendVerificationOtp(w,userId,userPhone)
// 		} else{
// 			fmt.Println("Mobile no does not exist")
// 			util.Message(w,1301)
// 		}
// 	} else{
// 		fmt.Println("User already verified")
// 		util.Message(w,1004)
// 	}

// }

// func sendVerificationOtp(w http.ResponseWriter,id string,phone string){
// 	count,lastModified:=GetOtpVerificationCount(id)
// 	fmt.Println("count and lastModified")
// 	fmt.Println(count)
// 	fmt.Println(lastModified)
// 	if count<MAXCOUNT{
// 		// otp:=util.GenerateOtpString()
// 		otp:=util.GenerateToken()
// 		EnterVerificationOtp(id,otp)
// 		sms.SendSMS(phone,otp)
// 		fmt.Println("OTP sent")
// 		util.Message(w,200)
// 		return
// 	}

// 	duration:=calculateDuration(lastModified)
// 	if duration < MAXHOUR{
// 		util.Message(w,1302)
// 		return
// 	}
	
// 	DeleteOtp(id)
// 	fmt.Println("Deleting")
// 	// otp:=util.GenerateOtpString()
// 	otp:=util.GenerateToken()
// 	EnterVerificationOtp(id,otp)
// 	sms.SendSMS(phone,otp)
// 	fmt.Println("OTP sent")
// 	util.Message(w,200)
// 	return
	
// }

// func calculateDuration(lastModified string) int{
// 	layout:="2009-11-10 23:00:00 +0000 UTC"
// 	t,_:=time.Parse(layout,lastModified)
// 	now:=time.Now()
// 	duration:=now.Sub(t)
// 	h:=duration.Hours()
// 	h_int:=int(h)
// 	print(h_int)
// 	return h_int
// }