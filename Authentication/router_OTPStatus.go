package Authentication

import(
	"net/http"
	util "miti-microservices/Util"
	database "miti-microservices/Database"
	"bytes"
	"encoding/json"
	"log"
)

func OTPStatus(w http.ResponseWriter,r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	otpStatusHeader:=OTPStatusHeader{}

	content:=OTPStatusResponse{}
	statusCode:=0
	moveTo:=0

	otpStatusResponseHeader:=OTPStatusResponseHeader{}
	var data map[string]string

	db:=database.DBConnection()
	list:=[]bool{false,false,false,false,false,false}
	errorList:=util.GetErrorList(list)
	
	util.GetHeader(r,&otpStatusHeader)
	sessionId:=otpStatusHeader.Cookie

	userId,loginStatus,dbError:=util.GetUserIdFromTemporarySession3(db,sessionId)
	errorList.DatabaseError=dbError
	util.APIHitLog("OTPStatus",ipAddress,sessionId)
	if(loginStatus=="Error" && !errorList.DatabaseError){
		errorList.TemporarySessionError=true
	}

	code:=0
	if(!errorList.TemporarySessionError && !errorList.DatabaseError){
		code,dbError=OTPHelper(db,userId)
		errorList.DatabaseError=dbError
	}
	
	util.APIHitLog("OTPStatus",ipAddress,sessionId)
	
	
	if((code==3005 || code==3003 || code==3004) && !errorList.TemporarySessionError && !errorList.DatabaseError){
		statusCode=200
		moveTo=0
	}else{
		statusCode=code
		moveTo=0
	}

	code=util.GetCode(errorList)
	if(code==200){
		content.Code=statusCode
	}else{
		content.Code=code
	}
	content.Message=util.GetMessageDecode(content.Code)
	content.MoveTo=moveTo
	otpStatusResponseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(otpStatusResponseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
	p:=&content
	util.ResponseLog("OTPStatus",ipAddress,sessionId,content.Code,*p)
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}