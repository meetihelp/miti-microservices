package Authentication

import(
	"net/http"
	util "miti-microservices/Util"
	sms "miti-microservices/Notification/SMS"
	database "miti-microservices/Database"
	// "log"
)
func ReSendOTP(w http.ResponseWriter,r *http.Request){
	smsHeader:=SMSHeader{}
	util.GetHeader(r,&smsHeader)
	sessionId:=smsHeader.Cookie
	db:=database.DBConnection()
	userId,code:=OTPHelper(db,sessionId)
	if(code==3003){
		phone,_:=GetPhoneFromUserId(db,userId)
		sms.ReSendSMSHelper(phone)
		util.Message(w,200)
		return
	}else if(code==3005){
		// DeleteOtp(userId)
		phone,_:=GetPhoneFromUserId(db,userId)
		sms.ReSendSMSHelper(phone)
		util.Message(w,200)
		return
	}else if(code==3004){
		phone,_:=GetPhoneFromUserId(db,userId)
		otpCode:=InsertOTP(db,userId,sessionId)
		err:=SendOTP(phone,otpCode)
		if(err=="Ok"){
        // resp,err:=SendOTP(phone,otpCode)
        // if(err==nil && resp.StatusCode==http.StatusOK){
            util.Message(w,200)
        } else{
            // log.Println(err)
            util.Message(w,200)
        }
	} else {
		util.Message(w,code)
	} 

	db.Close()
	
}