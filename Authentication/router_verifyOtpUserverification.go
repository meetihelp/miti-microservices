package Authentication

import(
	"net/http"
	"fmt"
	util "miti-microservices/Util"
    database "miti-microservices/Database"
    "io/ioutil"
    "encoding/json"
    "bytes"
    "log"
)

func VerifyOTPUserverification(w http.ResponseWriter,r *http.Request){
    ipAddress:=util.GetIPAddress(r)
    verifyOtpHeader:=VerifyOTPHeader{}
    util.GetHeader(r,&verifyOtpHeader)

    fmt.Print("Verify OTP Header")
    fmt.Println(verifyOtpHeader)

    sessionId:=verifyOtpHeader.Cookie
    statusCode:=0
    moveTo:=0
    var data map[string]string
    content:=VerifyOTPResponse{}

    db:=database.DBConnection()
    responseHeader:=VerifyOTPResponseHeader{}
    userId,sessionErr:=util.GetUserIdFromTemporarySession2(db,sessionId)
    if sessionErr=="Error"{
        // fmt.Println("Session Does not exist")
        // util.Message(w,1003)
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
    requestBody,err:=ioutil.ReadAll(r.Body)
    if err!=nil{
        // fmt.Println("Could not read body")
        // util.Message(w,1000)
        statusCode=1000
        moveTo=0
        content.Code=statusCode
        content.MoveTo=moveTo
        content.Message=util.GetMessageDecode(statusCode)
        headerBytes:=new(bytes.Buffer)
        json.NewEncoder(headerBytes).Encode(responseHeader)
        responseHeaderBytes:=headerBytes.Bytes()
        if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
            panic(err)
        }
        w=util.GetResponseFormatHeader(w,data)
        p:=&content
        enc := json.NewEncoder(w)
        err:= enc.Encode(p)
        if err != nil {
            log.Fatal(err)
        }
        return 
    }

    otpVerification:=OTPVerification{}
    errUserData:=json.Unmarshal(requestBody,&otpVerification)
    if errUserData!=nil{
        // fmt.Println("Could not Unmarshall user data")
        // util.Message(w,1001)
        statusCode=1001
        moveTo=0
        content.Code=statusCode
        content.MoveTo=moveTo
        content.Message=util.GetMessageDecode(statusCode)
        headerBytes:=new(bytes.Buffer)
        json.NewEncoder(headerBytes).Encode(responseHeader)
        responseHeaderBytes:=headerBytes.Bytes()
        if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
            panic(err)
        }
        w=util.GetResponseFormatHeader(w,data)
        p:=&content
        enc := json.NewEncoder(w)
        err:= enc.Encode(p)
        if err != nil {
            log.Fatal(err)
        }
        return 
    }

    otpVerification.UserId=userId
    otpVerification.SessionId=sessionId
    //SANITIZE USER AND PROFILE DATA
    sanatizationStatus :=Sanatize(otpVerification)
    if sanatizationStatus =="Error"{
        // fmt.Println("User data invalid")
        // util.Message(w,1002)
        statusCode=1002
        moveTo=0
        content.Code=statusCode
        content.MoveTo=moveTo
        content.Message=util.GetMessageDecode(statusCode)
        headerBytes:=new(bytes.Buffer)
        json.NewEncoder(headerBytes).Encode(responseHeader)
        responseHeaderBytes:=headerBytes.Bytes()
        if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
            panic(err)
        }
        w=util.GetResponseFormatHeader(w,data)
        p:=&content
        enc := json.NewEncoder(w)
        err:= enc.Encode(p)
        if err != nil {
            log.Fatal(err)
        }
        return
    }

    fmt.Print("Verify otp Body:")
    fmt.Println(otpVerification)

    otpVerify,otpVerificationDB:=VerifyOTPDB(db,otpVerification.UserId,otpVerification.OTP)
    if otpVerify{
        //CHANGE STATUS OF USER TO VERIFIED
        // ChangeVerificationStatus(userId)
        // util.InsertSessionValue(sessionId,userId,ipAddress)
        // util.DeleteTemporarySession(sessionId)
        IsUserVerified,IsProfileCreated,Preference:=LoadingPageQuery(db,userId)
        if(!IsUserVerified){
            ChangeVerificationStatus(db,userId)
            statusCode=1005
            moveTo=4
            content.Code=statusCode
            content.MoveTo=moveTo
            content.Message=util.GetMessageDecode(statusCode)
            headerBytes:=new(bytes.Buffer)
            json.NewEncoder(headerBytes).Encode(responseHeader)
            responseHeaderBytes:=headerBytes.Bytes()
            if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
                panic(err)
            }
        }else if(!IsProfileCreated){
            statusCode=1005
            moveTo=4
            content.Code=statusCode
            content.MoveTo=moveTo
            content.Message=util.GetMessageDecode(statusCode)
            headerBytes:=new(bytes.Buffer)
            json.NewEncoder(headerBytes).Encode(responseHeader)
            responseHeaderBytes:=headerBytes.Bytes()
            if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
                panic(err)
            }
        }else if(Preference<NUM_OF_PREFERENCE){
            statusCode=1006
            moveTo=5
            content.Code=statusCode
            content.MoveTo=moveTo
            content.Message=util.GetMessageDecode(statusCode)
            content.Preference=Preference
            headerBytes:=new(bytes.Buffer)
            json.NewEncoder(headerBytes).Encode(responseHeader)
            responseHeaderBytes:=headerBytes.Bytes()
            if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
                panic(err)
            }
        }else{
            util.InsertSessionValue(db,sessionId,userId,ipAddress)
            util.DeleteTemporarySession(db,sessionId)
            statusCode=200
            moveTo=6
            content.Code=statusCode
            content.MoveTo=moveTo
            content.Message=util.GetMessageDecode(statusCode)
            headerBytes:=new(bytes.Buffer)
            json.NewEncoder(headerBytes).Encode(responseHeader)
            responseHeaderBytes:=headerBytes.Bytes()
            if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
                panic(err)
            }
        }
    } else{
        // fmt.Println(otpVerificationDB.FailCount)
        UpdateFailCount(db,userId,otpVerificationDB.FailCount)
        // util.Message(w,1401)
        statusCode=1401
        moveTo=0
        content.Code=statusCode
        content.MoveTo=moveTo
        content.Message=util.GetMessageDecode(statusCode)
        headerBytes:=new(bytes.Buffer)
        json.NewEncoder(headerBytes).Encode(responseHeader)
        responseHeaderBytes:=headerBytes.Bytes()
        if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
            panic(err)
        }
    }
    w=util.GetResponseFormatHeader(w,data)
    p:=&content
    fmt.Print("Verify OTP Response:")
    fmt.Println(*p)     
    enc := json.NewEncoder(w)
    err= enc.Encode(p)
    if err != nil {
        log.Fatal(err)
    }

    db.Close()
}