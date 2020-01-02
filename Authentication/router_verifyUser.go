package Authentication

import(
	"net/http"
	// "fmt"
    "log"
	util "miti-microservices/Util"
    "encoding/json"
    "bytes"
    // "encoding/json"
    // "io/ioutil"
)

func VerifyUser(w http.ResponseWriter,r *http.Request){
    // ipAddress:=util.GetIPAddress(r)
    verifyOtpHeader:=VerifyOTPHeader{}
    util.GetHeader(r,&verifyOtpHeader)
    sessionId:=verifyOtpHeader.Cookie
    userId,sessionErr:=util.GetUserIdFromTemporarySession(sessionId)
    // fmt.Println(sessionId)
    statusCode:=0
    moveTo:=0
    var data map[string]string
    content:=OTPResponse{}
    responseHeader:=OTPResponseHeader{}
    if sessionErr=="Error"{
        // statusCode=1003
        // moveTo=0
        // content.Code=statusCode
        // content.MoveTo=moveTo
        // content.Message=util.GetMessageDecode(statusCode)
        // headerBytes:=new(bytes.Buffer)
        // json.NewEncoder(headerBytes).Encode(responseHeader)
        // responseHeaderBytes:=headerBytes.Bytes()
        // if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        //     panic(err)
        // }
        // w=util.GetResponseFormatHeader(w,data)
        // p:=&content
        // enc := json.NewEncoder(w)
        // err:= enc.Encode(p)
        // if err != nil {
        //     log.Fatal(err)
        // }
        // fmt.Println("Session Does not exist")
        // util.Message(w,1003)
        content,w:=util.GetSessionErrorContent(w)
        p:=&content
        enc := json.NewEncoder(w)
        err:= enc.Encode(p)
        if err != nil {
            log.Fatal(err)
        }
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
    _,code:=OTPHelper(sessionId)
    if status=="Ok"{
        // if(code==3003 || code ==3004 || code ==3005){
        if(code==200){
            otpCode:=InsertOTP(userId,sessionId)
            err:=SendOTP(phone,otpCode)
            if err=="Ok"{
            // resp,err:=SendOTP(phone,otpCode)
            // if(err==nil && resp.StatusCode==http.StatusOK){
                // util.Message(w,200)
                statusCode=200
                moveTo=0
                content.Code=statusCode
                content.MoveTo=moveTo
                content.Message=util.GetMessageDecode(statusCode)
                headerBytes:=new(bytes.Buffer)
                json.NewEncoder(headerBytes).Encode(responseHeader)
                responseHeaderBytes:=headerBytes.Bytes()
                if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
                    // panic(err)
                    log.Println(err)
                }
            } else{
                // util.Message(w,200)
                // fmt.Println(err)
                statusCode=200
                moveTo=0
                content.Code=statusCode
                content.MoveTo=moveTo
                content.Message=util.GetMessageDecode(statusCode)
                headerBytes:=new(bytes.Buffer)
                json.NewEncoder(headerBytes).Encode(responseHeader)
                responseHeaderBytes:=headerBytes.Bytes()
                if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
                    // panic(err)
                    log.Println(err)
                }
            }
        } else {
            // util.Message(w,code)
            statusCode=code
            moveTo=0
            content.Code=statusCode
            content.MoveTo=moveTo
            content.Message=util.GetMessageDecode(statusCode)
            headerBytes:=new(bytes.Buffer)
            json.NewEncoder(headerBytes).Encode(responseHeader)
            responseHeaderBytes:=headerBytes.Bytes()
            if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
                // panic(err)
                log.Println(err)
            }
        } 
        w=util.GetResponseFormatHeader(w,data)
        p:=&content
        enc := json.NewEncoder(w)
        err:= enc.Encode(p)
        if err != nil {
            log.Fatal(err)
        }
        // otpCode:=InsertOTP(userId,sessionId)
        // resp,err:=SendOTP(phone,otpCode)
        // if(err==nil && resp.StatusCode==http.StatusOK){
        //     util.Message(w,200)
        // } else{
        //     util.Message(w,200)
        //     // fmt.Println(err)
        // }
    }
    
}
