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
	util.Message(w,code)
}