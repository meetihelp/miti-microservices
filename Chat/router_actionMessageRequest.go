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

func ActionMessageRequest(w http.ResponseWriter,r *http.Request){
	header:=AcceptMessageRequestHeader{}
	util.GetHeader(r,&header)
	sessionId:=header.Cookie
	userId,dErr:=util.GetUserIdFromSession(sessionId)

	fmt.Print("AcceptMessageRequest Header->")
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
	
	acceptMessageRequestData:=AcceptMessageRequestDS{}
	profileRequestErr:=json.Unmarshal(requestBody,&acceptMessageRequestData)
	if profileRequestErr!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}

	fmt.Print("ActionMessageRequest Body:->")
	fmt.Println(acceptMessageRequestData)

	actionRequestId:=acceptMessageRequestData.RequestId
	action:=acceptMessageRequestData.Action
	senderPhone:=acceptMessageRequestData.Phone
	phone:=GetUserPhone(userId)
	updatedAt:=util.GetTime()
	if(action=="Accept"){
		userId2,updatedAtTemp:=UpdateMessageRequestDB(phone,senderPhone,action,actionRequestId,updatedAt)
		updatedAt=updatedAtTemp
		InsertChatDetail(userId,userId2)
	}else if(action=="Reject"){
		_,updatedAt=UpdateMessageRequestDB(phone,senderPhone,action,actionRequestId,updatedAt)
	}else{
		util.Message(w,1002)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&ActionMessageRequestResponse{Code:200,Message:msg,RequestId:actionRequestId,CreatedAt:updatedAt}
	fmt.Print("ActionMessageRequest Response:->")
	fmt.Println(*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}