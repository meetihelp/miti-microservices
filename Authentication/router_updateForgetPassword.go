package Authentication

import(
	"fmt"
	"net/http"
	// "log"
	"io/ioutil"
	"encoding/json"
   	util "miti-microservices/Util"
)

func UpdateForgetPassword(w http.ResponseWriter, r *http.Request){
	// ipAddress:=util.GetIPAddress(r)
	//GET HEADER 
	header:=VerificationHeader{}
	util.GetHeader(r,&header)
	sessionId:=header.Cookie
	userId,_:=util.GetUserIdFromTemporarySession(sessionId)
	// Read body data
	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}

	updateForgetPasswordData :=UpdateForgetPasswordDS{}
	errUpdateForgetPasswordData:=json.Unmarshal(requestBody,&updateForgetPasswordData)
	if errUpdateForgetPasswordData!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}
	// sanatizationStatus :=Sanatize(updateForgetPasswordData)
	// if sanatizationStatus =="Error"{
	// 	fmt.Println("User data invalid")
	// 	util.Message(w,1002)
	// 	return
	// }
	status:=CanUserUpdatePassword(sessionId)
	if status=="Ok"{
		UpdatePasswordFunc(userId,updateForgetPasswordData.Password)
		util.DeleteTemporarySession(sessionId)
		DeleteForgetPasswordSession(sessionId)
		util.Message(w,200)	
	}
	

}