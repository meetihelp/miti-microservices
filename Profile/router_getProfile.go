package Profile

import(
	"fmt"
	"net/http"
	// "log"
	"io/ioutil"
	// "strings"
	"encoding/json"
   util "miti-microservices/Util"
)

func GetProfile(w http.ResponseWriter, r *http.Request){
	//CHECK IF USER IS AUTHORIZED TO SEE THE PROFILE
	header:=GetProfileHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie
	userId,dErr:=util.GetUserIdFromSession(sessionId)
	if dErr=="Error"{
		fmt.Println("Session Does not exist")
		util.Message(w,1003)
		return
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}
	
	profileRequest:=ProfileRequest{}
	profileRequestErr:=json.Unmarshal(requestBody,&profileRequest)
	if profileRequestErr!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}
	profileViewAuthorization:=ProfileViewAuthorization(userId,profileRequest.UserId)
	if profileViewAuthorization=="Error"{
		profileResponse:=GetUnAuthorizedProfileDB(profileRequest.UserId)
		SendProfile(w,2001,profileResponse)
	}else{
		profileResponse:=GetAuthorizedProfileDB(profileRequest.UserId)
		SendProfile(w,200,profileResponse)
	}

	//RETURN PROFILE
	
	

}