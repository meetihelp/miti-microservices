package Authentication
import (
	"net/http"
	"fmt"
	util "app/Util"
	"io/ioutil"
	"encoding/json"
)


func Login(w http.ResponseWriter,r *http.Request){
	//Get IP Address of Client
	ipAddress:=util.GetIPAddress(r)
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
	//Check if User is verified or not
	//session of Unverified user is stored separately to reduce the risk.... 
	//...of accesing the unauthorized data without verification 
	userId,loginStatus=util.GetUserIdFromTemporarySession(sessionId)
	if loginStatus=="Ok"{
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
	fmt.Println(userData);
	//Check if the user data is proper or not
	sanatizationStatus :=Sanatize(userData)
	if sanatizationStatus =="Error"{
		fmt.Println("User data invalid")
		util.Message(w,1002)
		return
	}

	//Check if the credentials given by user is Proper or not
	userId,loginStatus=CheckUserCredentials(userData)
	// if loginStatus=="WrongPassword"{
	// 	util.Message(w,1502)
	// 	return
	// }
	if loginStatus=="NoUser"{
		util.Message(w,1501)
		return
	}
	if loginStatus=="Unverified"{
		cookie:=util.InsertTemporarySession(userId,ipAddress)
		w.Header().Set("Miti-Cookie",cookie)
		util.Message(w,1005)
		return
	} 
	if loginStatus=="Ok"{
		cookie:=util.InsertTemporarySession(userId,ipAddress)
		w.Header().Set("Miti-Cookie",cookie)
		util.Message(w,200)
		return
	}
}
