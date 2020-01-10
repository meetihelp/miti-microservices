package Authentication

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	// "strings"
	"encoding/json"
   util "miti-microservices/Util"
)

func GetPhoneStatus(w http.ResponseWriter, r *http.Request){
	header:=GetPhoneStatusHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie
	_,dErr:=util.GetUserIdFromSession(sessionId)
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

	getPhoneStatusRequest:=GetPhoneStatusRequest{}
	profileRequestErr:=json.Unmarshal(requestBody,&getPhoneStatusRequest)
	if profileRequestErr!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}

	statusList:=CheckPhoneNumberStatusDB(getPhoneStatusRequest)
	w.Header().Set("Content-Type", "application/json")
	statusCode:=200
	msg:=util.GetMessageDecode(statusCode)
	p:=&GetPhoneStatusResponse{PhoneStatus:statusList,Code:statusCode,Message:msg}
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}