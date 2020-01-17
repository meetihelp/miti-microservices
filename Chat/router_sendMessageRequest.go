package Chat

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	// "strings"
	"encoding/json"
   util "miti-microservices/Util"
   profile "miti-microservices/Profile"
   sms "miti-microservices/Notification/SMS"
)

func SendMessageRequest(w http.ResponseWriter,r *http.Request){
	header:=SendMessageRequestHeader{}
	util.GetHeader(r,&header)
	sessionId:=header.Cookie
	userId,dErr:=util.GetUserIdFromSession(sessionId)
	fmt.Print("SendMessageRequest Header:->")
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
	
	sendMessageRequestData:=SendMessageRequestDS{}
	profileRequestErr:=json.Unmarshal(requestBody,&sendMessageRequestData)
	if profileRequestErr!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}

	fmt.Print("SendMessageRequest Body:->")
	fmt.Println(sendMessageRequestData)

	requestId:=sendMessageRequestData.RequestId
	senderPhone:=GetUserPhone(userId)
	phone:=sendMessageRequestData.Phone
	messageType:=sendMessageRequestData.MessageType
	messageContent:=sendMessageRequestData.MessageContent
	senderName:=profile.GetUserName(userId)
	createdAt:=util.GetTime()
	createdAt=InsertMessageRequestDB(userId,senderName,senderPhone,phone,requestId,messageType,messageContent,createdAt)
	
	availability:=IsPhoneNumberExist(phone)
	if(availability=="No"){
		sms.MessageRequestNotificaton(senderName,senderPhone,phone)
	}
	
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&SendMessageRequestResponse{Code:200,Message:msg,RequestId:requestId,CreatedAt:createdAt}
	fmt.Print("SendMessageRequest Response:->")
	fmt.Println(*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}