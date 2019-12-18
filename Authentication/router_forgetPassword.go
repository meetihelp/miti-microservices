package Authentication

import(
	"fmt"
	"net/http"
	// "log"
	"io/ioutil"
	"encoding/json"
   	util "app/Util"
)

func ForgetPassword(w http.ResponseWriter, r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	//GET HEADER 
	header:=RegisterHeader{}
	util.GetHeader(r,&header)

	// Read body data
	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}

	forgetPasswordData :=ForgetPasswordDS{}
	errForgetPasswordData:=json.Unmarshal(requestBody,&forgetPasswordData)
	if errForgetPasswordData!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}
	// sanatizationStatus :=Sanatize(forgetPasswordData)
	// if sanatizationStatus =="Error"{
	// 	fmt.Println("User data invalid")
	// 	util.Message(w,1002)
	// 	return
	// }
	
	//Generate A session
	userId,status:=GetUserIdFromPhone(forgetPasswordData.Phone)
	if status=="Ok"{
		cookie:=util.InsertTemporarySession(userId,ipAddress)
		w.Header().Set("Miti-Cookie",cookie)
		otp:=InsertOTP(userId,cookie)
		InsertForgetPasswordStatus(cookie)
		SendOTP(forgetPasswordData.Phone,otp)
		util.Message(w,200)
	}

}