package Authentication

import(
	"net/http"
    "log"
	util "miti-microservices/Util"
    database "miti-microservices/Database"
    "encoding/json"
    "bytes"
	sms "miti-microservices/Notification/SMS"
)
func ReSendOTP(w http.ResponseWriter,r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	resendOTPHeader:=ResendOTPHeader{}

    content:=ResendOTPResponse{}
	statusCode:=0
    moveTo:=0

    resendOTPResponseHeader:=ResendOTPResponseHeader{}
    var data map[string]string
    
    db:=database.DBConnection()
    list:=[]bool{false,false,false,false,false,false}
    errorList:=util.GetErrorList(list)
    

    util.GetHeader(r,&resendOTPHeader)
	sessionId:=resendOTPHeader.Cookie

    userId,sessionErr,dbError:=util.GetUserIdFromTemporarySession3(db,sessionId)
    errorList.DatabaseError=dbError
    util.APIHitLog("ReSendOTP",ipAddress,sessionId)
    if (sessionErr=="Error" && !errorList.DatabaseError){
        errorList.TemporarySessionError=true
    }

    var phone string
    // var status string
    code:=0
    if(!errorList.TemporarySessionError && !errorList.DatabaseError){
    	phone,_,dbError=GetPhoneFromUserId(db,userId)
        errorList.DatabaseError=dbError
        if(!errorList.DatabaseError){
            code,dbError=OTPHelper(db,userId)
            errorList.DatabaseError=dbError
        }
    }

	if(!errorList.TemporarySessionError && !errorList.DatabaseError && (code==3003 ||code ==3005)){
		sms.ReSendSMSHelper(phone)
		statusCode=200
		moveTo=0
	}else if(!errorList.TemporarySessionError && !errorList.DatabaseError && code==3004){
		otpCode,dbError:=InsertOTP(db,userId,sessionId)
        errorList.DatabaseError=dbError
        if(!errorList.DatabaseError){
            err:=SendOTP(phone,otpCode)
            if(err=="Ok"){
                statusCode=200
                moveTo=0
            }else{
                //Error in sending otp
            }
        }
	} else {
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
    resendOTPResponseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(resendOTPResponseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
    p:=&content
    util.ResponseLog("ReSendOTP",ipAddress,sessionId,content.Code,*p)
    enc := json.NewEncoder(w)
    err:= enc.Encode(p)
    if err != nil {
        log.Fatal(err)
    }
    db.Close()
	
}