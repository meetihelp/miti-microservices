package Authentication

import(
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
    util "app/Util"
)


func UpdatePassword(w http.ResponseWriter,r *http.Request){
	updatePasswordHeader:=UpdatePasswordHeader{}
	util.GetHeader(r,&updatePasswordHeader)
	sessionId:=updatePasswordHeader.Cookie

	userId,status:=util.GetUserIdFromSession(sessionId)
	if status!="Ok"{
		fmt.Println("Session does not exist")
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

	//UNMARSHILING DATA
	passwordChangeData :=PasswordChange{}
	errUserData:=json.Unmarshal(requestBody,&passwordChangeData)
	if errUserData!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}

	sanatizationStatus :=Sanatize(passwordChangeData)
	if sanatizationStatus =="Error"{
		fmt.Println("User data invalid")
		util.Message(w,1002)
		return
	}

	status=CheckUserById(userId,passwordChangeData.OldPassword)

	if status=="Ok"{
		//UPDATE PASSWORD
		UpdatePasswordFunc(userId,passwordChangeData.NewPassword)
		util.Message(w,200)
	} else{
		//SEND ERROR MESSAGE

	}
}