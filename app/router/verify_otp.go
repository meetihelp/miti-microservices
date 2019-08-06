package main

import(
	"net/http"
	"fmt"
	// "log"
    CD "app/Model/CreateDatabase"
	database "app/Model/UseDatabase"
	util "app/Utility"
    "encoding/json"
    "io/ioutil"
)

type Verify_OTP_Header struct{
    Cookie string `header:"Miti-Cookie"`
}
func verify_otp(w http.ResponseWriter,r *http.Request){
    verify_otp_header:=Verify_OTP_Header{}
    util.GetHeader(r,&verify_otp_header)
    session_id:=verify_otp_header.Cookie
    user_id,session_err:=database.Get_user_id_from_session(session_id)
    if session_err=="ERROR"{
        fmt.Println("Session Does not exist")
        util.Message(w,1003)
        return
    }
    //Read body data
    requestBody,err:=ioutil.ReadAll(r.Body)
    if err!=nil{
        fmt.Println("Could not read body")
        util.Message(w,1000)
        return 
    }

    otp_verification:=CD.OTP_verification{}
    err_user_data:=json.Unmarshal(requestBody,&otp_verification)
    if err_user_data!=nil{
        fmt.Println("Could not Unmarshall user data")
        util.Message(w,1001)
        return 
    }

    otp_verification.User_id=user_id

    //SANITIZE USER AND PROFILE DATA
    sanatization_status :=CD.Sanatize(otp_verification)
    if sanatization_status =="ERROR"{
        fmt.Println("User data invalid")
        util.Message(w,1002)
        return
    }

    otp_verify:=database.Verify_OTP(otp_verification.User_id,otp_verification.Verification_otp)
    if otp_verify{
        //CHANGE STATUS OF USER TO VERIFIED
        database.Change_Verification_Status(user_id)
        util.Message(w,200)
    } else{
        util.Message(w,1401)
    }




}