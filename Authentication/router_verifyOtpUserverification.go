package Authentication

import(
	"net/http"
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

    content:=VerifyOTPResponse{}
    statusCode:=0
    moveTo:=0
    preference:=0

    responseHeader:=VerifyOTPResponseHeader{}
    var data map[string]string

    db:=database.DBConnection()
    list:=[]bool{false,false,true,true,true,true}
    errorList:=util.GetErrorList(list)

    util.GetHeader(r,&verifyOtpHeader)
    sessionId:=verifyOtpHeader.Cookie

    //Checking Session Status
    userId,sessionErr,dbError:=util.GetUserIdFromTemporarySession3(db,sessionId)
    errorList.DatabaseError=dbError
    util.APIHitLog("VerifyOTP",ipAddress,sessionId)
    if (sessionErr=="Error" && !errorList.DatabaseError){
        util.TemporarySessionLog("VerifyOTP",ipAddress,sessionId,"Fail")
        errorList.TemporarySessionError=true
    }

    //Read Body
    requestBody,err:=ioutil.ReadAll(r.Body)
    if(!errorList.TemporarySessionError){
        util.TemporarySessionLog("VerifyOTP",ipAddress,sessionId,"Success")
        if(err==nil){
            errorList.BodyReadError=false
        }
    }

    //Unmarshall Body
    verifyOTPData:=VerifyOTPRequest{}
    if(!errorList.BodyReadError){
        errUserData:=json.Unmarshal(requestBody,&verifyOTPData)
        if(errUserData==nil){
            errorList.UnmarshallingError=false
        }
    }
    
    if(!errorList.UnmarshallingError){
        util.BodyLog("VerifyOTP",ipAddress,sessionId,verifyOTPData)
        sanatizationStatus :=Sanatize(verifyOTPData)
        if(sanatizationStatus=="Ok"){
            errorList.SanatizationError=false
        }
    }

    otpVerify:=false
    failCount:=0
    if(!errorList.SanatizationError){
        otpVerify,failCount,errorList.DatabaseError=VerifyOTPDB(db,userId,verifyOTPData.OTP)
    }
    
    if (otpVerify && !errorList.DatabaseError){
        IsUserVerified,IsProfileCreated,Preference,dbError:=LoadingPageQuery(db,userId)
        errorList.DatabaseError=dbError
        if(!IsUserVerified && !errorList.DatabaseError){
            errorList.DatabaseError=ChangeVerificationStatus(db,userId)
            if(!errorList.DatabaseError){
                statusCode=1005
                moveTo=4
            }
        }else if(!IsProfileCreated && !errorList.DatabaseError){
            statusCode=1005
            moveTo=4
        }else if(Preference<NUM_OF_PREFERENCE && !errorList.DatabaseError){
            statusCode=1006
            moveTo=5
            preference=Preference
        }else if(!errorList.DatabaseError){
            errorList.DatabaseError=util.InsertSessionValue(db,sessionId,userId,ipAddress)
            if(!errorList.DatabaseError){
                errorList.DatabaseError=util.DeleteTemporarySession(db,sessionId)
            }
            if(!errorList.DatabaseError){
                statusCode=200
                moveTo=6
            }      
        }
    } else{
        errorList.DatabaseError=UpdateFailCount(db,userId,failCount)
        if(!errorList.DatabaseError){
           statusCode=1401
            moveTo=0 
        }               
    }
    

    code:=util.GetCode(errorList)
    if(code==200){
        content.Code=statusCode
    }else{
        content.Code=code
    }
    content.Message=util.GetMessageDecode(content.Code)
    content.MoveTo=moveTo
    content.Preference=preference
    responseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(responseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
    p:=&content
    util.ResponseLog("VerifyOTP",ipAddress,sessionId,content.Code,*p)
    enc := json.NewEncoder(w)
    err= enc.Encode(p)
    if err != nil {
        log.Fatal(err)
    }
    db.Close()
}