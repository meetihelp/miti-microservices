package Authentication
import (
	"net/http"
	"fmt"
	util "app/Util"
	"io/ioutil"
	"encoding/json"
)


func Login(w http.ResponseWriter,r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	loginHeader:=LoginHeader{}
	util.GetHeader(r,&loginHeader)
	sessionId:=loginHeader.Cookie
	fmt.Println(sessionId)
	userId,loginStatus:=util.GetUserIdFromSession(sessionId)
	fmt.Println("session "+loginStatus)
	if loginStatus=="Ok"{
		util.Message(w,200)
		return
	}
	userId,loginStatus=util.GetUserIdFromUserVerificationSession(sessionId)
	fmt.Println("session "+loginStatus)
	if loginStatus=="Ok"{
		cookie:=sessionId
		w.Header().Set("Miti-Cookie",cookie)
		util.Message(w,1005)
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
	userData :=User{}
	errUserData:=json.Unmarshal(requestBody,&userData)
	if errUserData!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}

	sanatizationStatus :=Sanatize(userData)
	if sanatizationStatus =="Error"{
		fmt.Println("User data invalid")
		util.Message(w,1002)
		return
	}

	userId,loginStatus=CheckUser(userData)
	if loginStatus=="WrongPassword"{
		util.Message(w,1501)
		return
	}
	if loginStatus=="NoUser"{
		util.Message(w,1501)
		return
	}
	if loginStatus=="Unverified"{
		cookie:=util.InsertUserVerificationSession(userId,ipAddress)
		// http.SetCookie(w,&cookie)
		w.Header().Set("Miti-Cookie",cookie)
		util.Message(w,1005)
		return
	} 
	if loginStatus=="Ok"{
		cookie:=util.InsertSession(userId,ipAddress)
		fmt.Println("gaurav1")
		// http.SetCookie(w,&cookie)
		w.Header().Set("Miti-Cookie",cookie)
		util.Message(w,200)
		return
	}
}
