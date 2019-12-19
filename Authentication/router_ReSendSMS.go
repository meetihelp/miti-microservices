package Authentication

import(
	"net/http"
	util "miti-microservices/Util"
	sms "miti-microservices/Notification/SMS"
	"log"
)
func ReSendOTP(w http.ResponseWriter,r *http.Request){
	smsHeader:=SMSHeader{}
	util.GetHeader(r,&smsHeader)
	sessionId:=smsHeader.Cookie
	userId,code:=OTPHelper(sessionId)
	if(code==3003){
		phone,_:=GetPhoneFromUserId(userId)
		sms.ReSendSMSHelper(phone)
		util.Message(w,200)
		return
	}else if(code==3004){
		phone,_:=GetPhoneFromUserId(userId)
		otpCode:=InsertOTP(userId,sessionId)
        resp,err:=SendOTP(phone,otpCode)
        if(err==nil && resp.StatusCode==http.StatusOK){
            util.Message(w,200)
        } else{
            log.Println(err)
        }
	} else {
		util.Message(w,code)
	} 
	
}