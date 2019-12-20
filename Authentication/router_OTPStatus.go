package Authentication

import(
	"net/http"
	util "miti-microservices/Util"
)

func OTPStatus(w http.ResponseWriter,r *http.Request){
	otpStatusHeader:=OTPStatusHeader{}
	util.GetHeader(r,&otpStatusHeader)
	sessionId:=otpStatusHeader.Cookie
	_,code:=OTPHelper(sessionId)
	if(code==3005 || code==3003 || code==3004){
		util.Message(w,200)
	}else{
		util.Message(w,code)
	}
}