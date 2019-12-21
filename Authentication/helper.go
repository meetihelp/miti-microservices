package Authentication

import(
	"net/http"
	// "io/ioutil"
	// "encoding/json"	
	util "miti-microservices/Util"
	sms "miti-microservices/Notification/SMS"
	// "log"
	"time"
	// "reflect"
	// "fmt"
)

const (
	MAXCOUNT = 5
	MAXMINUTE = 10
	MAXFAILCOUNT=5
	MAXRESENDCOUNT=5
	ONEDAY=1440
	NUM_OF_PREFERENCE=6
)

// func SendPreference(w http.ResponseWriter,preferenceCreationStatus int,code int){
// 	w.Header().Set("Content-Type", "application/json")
// 	msg:=util.GetMessageDecode(code)
// 	p:=&PreferenceContent{Code:code,Message:msg,Preference:preferenceCreationStatus}
// 	enc := json.NewEncoder(w)
// 	err:= enc.Encode(p)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func SendOTP(phone string,otp string)(*http.Response,error){
	return sms.SendSMS(phone,otp)
}

func OTPHelper(sessionId string) (string,int){
	userId,loginStatus:=util.GetUserIdFromTemporarySession(sessionId)
	if loginStatus=="Ok"{
		otp:=GetOTPDetails(userId)
		duration:=CalculateDuration(otp.CreatedAt)
		if(duration>ONEDAY){
			DeleteOTP(userId)
			return userId,200
			// return userId,3005
		}
		if(otp.FailCount>=MAXFAILCOUNT){
			// util.Message(w,3000)
			return userId,3000
		}
		if(otp.ResendCount>MAXRESENDCOUNT){
			// util.Message(w,3001)
			return userId,3001
		}
		
		deliveryCount:=otp.DeliverCount
		// if(duration<MAXMINUTE && deliveryCount!=0){
		// 	// util.Message(w,3002)
		// 	return userId,3002
		// }
		if(duration<MAXMINUTE || deliveryCount==0){
			// phone,_:=GetPhoneFromUserId(userId)
			// sms.ReSendSMSHelper(phone)
			// util.Message(w,200)
			return userId,200
			// return userId,3003
		}
		if(duration>MAXMINUTE){
			// otpCode:=InsertOTP(userId,sessionId)
	  //       resp,err:=SendOTP(phone,otpCode)
	  //       if(err==nil && resp.StatusCode==http.StatusOK){
	  //           util.Message(w,200)
	  //       } else{
	  //           log.Println(err)
	  //       }
	  		return userId,200
	  		// return userId,3004
		}
	}
	return userId,1003
}

func CalculateDuration(lastModified string) int{
	layout:="2006-01-02 15:04:05"
	t,_:=time.Parse(layout,lastModified)
	now:=time.Now()
	now,_=time.Parse(layout,now.Format("2006-01-02 15:04:05"))
	duration:=now.Sub(t)
	h:=duration.Minutes()
	h_int:=int(h)
	return h_int
}


