package Authentication
import (
	"net/http"
	"fmt"
	util "miti-microservices/Util"
	// "io/ioutil"
	// "encoding/json"
)


func LoadingPage(w http.ResponseWriter,r *http.Request){
	//Get IP Address of Client
	// ipAddress:=util.GetIPAddress(r)
	//Read header of the client packet
	loginHeader:=LoginHeader{}
	util.GetHeader(r,&loginHeader)
	sessionId:=loginHeader.Cookie
	//Check if the user is already logged in? Using session value
	userId,loginStatus:=util.GetUserIdFromSession(sessionId)
	fmt.Println("session "+loginStatus)
	if loginStatus=="Ok"{
		util.Message(w,300)
		return
	}
	userId,loginStatus=util.GetUserIdFromTemporarySession(sessionId)
	if loginStatus=="Error"{
		util.Message(w,2000)
		return
	}else{
		 temporarySessionCase(w,userId)
	}
}

func temporarySessionCase(w http.ResponseWriter,userId string){
	IsUserVerified,IsProfileCreated,Preferece:=LoadingPageQuery(userId)
	if !IsUserVerified{
		util.Message(w,2001)
		return
	}

	if !IsProfileCreated{
		util.Message(w,2002)
		return
	}

	if Preferece<6{
		SendPreference(w,Preferece,2003)
		return
	}

	util.Message(w,200)
}
