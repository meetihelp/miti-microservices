package Authentication

import(
	"net/http"
	util "miti-microservices/Util"
	"bytes"
	"encoding/json"
	"log"
)

func OTPStatus(w http.ResponseWriter,r *http.Request){
	otpStatusHeader:=OTPStatusHeader{}
	util.GetHeader(r,&otpStatusHeader)
	sessionId:=otpStatusHeader.Cookie
	_,code:=OTPHelper(sessionId)
	content:=OTPStatusResponseContent{}
	responseHeader:=OTPStatusResponseHeader{}
	var data map[string]string
	if(code==3005 || code==3003 || code==3004){
		statusCode:=200
		moveTo:=0
		content.Code=statusCode
		content.MoveTo=moveTo
		content.Message=util.GetMessageDecode(statusCode)
		responseHeader.ContentType="application/json"
		headerBytes:=new(bytes.Buffer)
		json.NewEncoder(headerBytes).Encode(responseHeader)
		responseHeaderBytes:=headerBytes.Bytes()
		if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        	panic(err)
    	}
		w=util.GetResponseFormatHeader(w,data)
		// util.Message(w,200)
	}else{
		statusCode:=code
		moveTo:=0
		content.Code=statusCode
		content.MoveTo=moveTo
		content.Message=util.GetMessageDecode(statusCode)
		responseHeader.ContentType="application/json"
		headerBytes:=new(bytes.Buffer)
		json.NewEncoder(headerBytes).Encode(responseHeader)
		responseHeaderBytes:=headerBytes.Bytes()
		if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        	panic(err)
    	}
		w=util.GetResponseFormatHeader(w,data)
		// util.Message(w,code)
	}

	p:=&content
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}