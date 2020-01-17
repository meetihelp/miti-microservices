package Security

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	// "strings"
	"encoding/json"
   util "miti-microservices/Util"
   // profile "miti-microservices/Profile"
   sms "miti-microservices/Notification/SMS"
)

func AlertMessage(w http.ResponseWriter, r *http.Request){
	header:=AlertMessageHeader{}
	util.GetHeader(r,&header)

	sessionId:=header.Cookie
	userId,dErr:=util.GetUserIdFromSession(sessionId)

	fmt.Print("AlertMessageHeader:")
	fmt.Println(header)
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
	alertMessageRequest:=AlertMessageRequest{}
	profileRequestErr:=json.Unmarshal(requestBody,&alertMessageRequest)
	if profileRequestErr!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}

	fmt.Print("CreatePrimaryTrustChain Body:->")
	fmt.Println(alertMessageRequest)

	latitude:=alertMessageRequest.Latitude
	longitude:=alertMessageRequest.Longitude

	if(latitude=="" || longitude==""){
		util.Message(w,1002)
		return
	}
	phoneList:=GetPrimaryTrustPhoneList(userId)
	name:=GetUserName(userId)

	status:=sms.AlertNotification(phoneList,name,latitude,longitude)

	code:=200
	if(status=="Ok"){
		code=200
	}else{
		code=2005
	}

	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(code)
	p:=&AlertMessageResponse{Code:code,Message:msg}
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}