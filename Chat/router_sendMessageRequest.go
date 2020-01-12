package Chat

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	// "strings"
	"encoding/json"
   util "miti-microservices/Util"
)

func SendMessageRequest(w http.ResponseWriter,r *http.Request){
	header:=SendMessageRequestHeader{}
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
	
	sendMessageRequestData:=SendMessageRequestDS{}
	profileRequestErr:=json.Unmarshal(requestBody,&sendMessageRequestData)
	if profileRequestErr!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}

	requestId:=sendMessageRequestData.RequestId
	phone:=sendMessageRequestData.Phone
	messageType:=sendMessageRequestData.MessageType
	messageContent:=sendMessageRequestData.MessageContent
	createdAt:=util.GetTime()
	createdAt=InsertMessageRequestDB(userId,phone,requestId,messageType,messageContent,createdAt)

	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&SendMessageRequestResponse{Code:200,Message:msg,RequestId:requestId,CreatedAt:createdAt}
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}