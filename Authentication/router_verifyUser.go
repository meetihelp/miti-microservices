package Authentication

import(
	"net/http"
    "log"
	util "miti-microservices/Util"
    database "miti-microservices/Database"
    "encoding/json"
    "bytes"
    "fmt"
)

func VerifyUser(w http.ResponseWriter,r *http.Request){
    ipAddress:=util.GetIPAddress(r)
    verifyUserHeader:=VerifyUserHeader{}

    content:=VerifyUserResponse{}
    statusCode:=0
    moveTo:=0

    verifyUserResponseHeader:=VerifyUserResponseHeader{}
    var data map[string]string

    db:=database.DBConnection()
    list:=[]bool{false,false,false,false,false,false}
    errorList:=util.GetErrorList(list)

    util.GetHeader(r,&verifyUserHeader)
    sessionId:=verifyUserHeader.Cookie
    
    userId,sessionErr,dbError:=util.GetUserIdFromTemporarySession3(db,sessionId)
    errorList.DatabaseError=dbError
    util.APIHitLog("VerifyUser",ipAddress,sessionId)
    if (sessionErr=="Error"){
        errorList.TemporarySessionError=true
    }

    var phone string
    var status string
    if(!errorList.TemporarySessionError && !errorList.DatabaseError){
        fmt.Println("Verify User Line 40")
        phone,status,dbError=GetPhoneFromUserId(db,userId)
        errorList.DatabaseError=dbError
        var code int
        if(!errorList.DatabaseError){
            fmt.Println("Verify User Line 45")
            code,dbError=OTPHelper(db,userId)
            errorList.DatabaseError=dbError
        }

        if(status=="Ok" && !errorList.DatabaseError){
            fmt.Println("Verify User Line 51")
            if(code==200){
                fmt.Println("Verify User Line 53")
                otpCode,dbError:=InsertOTP(db,userId,sessionId)
                errorList.DatabaseError=dbError
                if(!errorList.DatabaseError){
                    fmt.Println("Verify User Line 57")
                    _,err:=SendOTP(phone,otpCode)
                    if(err==nil){
                        fmt.Println("Verify User Line 60")
                        statusCode=200
                        moveTo=0
                    }else{
                        fmt.Println("Verify User Line 64")
                        //Error in sending otp
                        statusCode=1007
                        moveTo=0
                    }
                }
            }else{
                fmt.Println("Verify User Line 71")
                statusCode=code
                moveTo=0
            }

        }
    }

    code:=util.GetCode(errorList)
    if(code==200){
        fmt.Println("Verify User Line 80")
        content.Code=statusCode
    }else{
        fmt.Println("Verify User Line 84")
        content.Code=code
    }
    content.Message=util.GetMessageDecode(content.Code)
    content.MoveTo=moveTo
    verifyUserResponseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(verifyUserResponseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
    p:=&content
    util.ResponseLog("VerifyUser",ipAddress,sessionId,content.Code,*p)
    enc := json.NewEncoder(w)
    err:= enc.Encode(p)
    if err != nil {
        log.Fatal(err)
    }
    db.Close()
    
}
