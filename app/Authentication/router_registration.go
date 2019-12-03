package Authentication

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
   	util "app/Util"
)

func Register(w http.ResponseWriter, r *http.Request){
	//Get ip address of user
	ipAddress:=util.GetIPAddress(r)

	//GET HEADER 
	header:=RegisterHeader{}
	util.GetHeader(r,&header)

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

	//SANITIZE USER AND PROFILE DATA
	sanatizationStatus :=Sanatize(userData)
	if sanatizationStatus =="Error"{
		fmt.Println("User data invalid")
		util.Message(w,1002)
		return
	}

	//Check if user exists or not and take  action accordingly
	userId,ok:=userDataHandle(w,userData)
	if ok{
		userData.UserId=userId
		cookie:=util.InsertTemporarySession(userData.UserId,ipAddress)
		w.Header().Set("Miti-Cookie",cookie)
		util.Message(w,200)
	}
}

func userDataHandle(w http.ResponseWriter, userData User) (string,bool){
	userId,dbStatus:=EnterUserData(userData)
	if dbStatus ==1{
		log.Println("User Already exist")
		util.Message(w,1101)
		return userId,false
	} else{
		log.Println("User data entered successfully")
		return userId,true
	}
}



