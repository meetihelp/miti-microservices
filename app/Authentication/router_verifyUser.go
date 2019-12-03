package Authentication

import(
	"net/http"
	"fmt"
	util "app/Util"
    // "encoding/json"
    // "io/ioutil"
)

func VerifyUser(w http.ResponseWriter,r *http.Request){
    // ipAddress:=util.GetIPAddress(r)
    verifyOtpHeader:=VerifyOTPHeader{}
    util.GetHeader(r,&verifyOtpHeader)
    sessionId:=verifyOtpHeader.Cookie
    userId,sessionErr:=util.GetUserIdFromTemporarySession(sessionId)
    if sessionErr=="Error"{
        fmt.Println("Session Does not exist")
        util.Message(w,1003)
        return
    }
    //Read body data
    // requestBody,err:=ioutil.ReadAll(r.Body)
    // if err!=nil{
    //     fmt.Println("Could not read body")
    //     util.Message(w,1000)
    //     return 
    // }

    // otpVerification:=OTPVerification{}
    // errUserData:=json.Unmarshal(requestBody,&otpVerification)
    // if errUserData!=nil{
    //     fmt.Println("Could not Unmarshall user data")
    //     util.Message(w,1001)
    //     return 
    // }

    // otpVerification.UserId=userId
    // otpVerification.SessionId=sessionId
    // //SANITIZE USER AND PROFILE DATA
    // sanatizationStatus :=Sanatize(otpVerification)
    // if sanatizationStatus =="Error"{
    //     fmt.Println("User data invalid")
    //     util.Message(w,1002)
    //     return
    // }

    phone,status:=GetPhoneFromUserId(userId)
    if status=="Ok"{
        otp:=InsertOTP(userId,sessionId)
        SendOTP(phone,otp)
        util.Message(w,200)
    }
}