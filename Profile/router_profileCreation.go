package Profile

import(
	"fmt"
	"net/http"
	// "log"
	"io/ioutil"
	// "strings"
	"encoding/json"
   util "miti-microservices/Util"
   auth "miti-microservices/Authentication"
)


func ProfileCreation(w http.ResponseWriter, r *http.Request){
	header:=ProfileCreationHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie

	userId,dErr:=util.GetUserIdFromSession(sessionId)
	// if dErr=="Error"{
	// 	fmt.Println("Session Does not exist")
	// 	util.Message(w,1003)
	// 	return
	// }

	if dErr=="Error"{
		userId,dErr=util.GetUserIdFromTemporarySession(sessionId)
		if dErr=="Error"{
			util.Message(w,1003)
			return
		}
	}


	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}
	profileData:=Profile{}
	errProfileData:=json.Unmarshal(requestBody,&profileData)
	if errProfileData!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}
	fmt.Println(profileData.Name)
	profileData.UserId=userId
	sanatizationStatus:=Sanatize(profileData)
	if sanatizationStatus =="Error"{
		fmt.Println("profile creation data invalid")
		util.Message(w,1002)
		return
	}
	// profile_data_handle(w,profile_data)
	auth.UpdateProfileCreationStatus(userId)
	EnterProfileData(profileData)
	util.Message(w,200)

}