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

func GetMessageRequest(w http.ResponseWriter,r *http.Request){
	header:=GetMessageRequestHeader{}
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
	
	getMessageRequestData:=GetMessageRequestDS{}
	profileRequestErr:=json.Unmarshal(requestBody,&getMessageRequestData)
	if profileRequestErr!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}

	fmt.Print("GetMessageRequest Body:->")
	fmt.Println(getMessageRequestData)
	createdAt:=getMessageRequestData.CreatedAt

	requestList:=GetMessageRequestDB(userId,createdAt)
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&GetMessageRequestResponse{Code:200,Message:msg,RequestList:requestList}
	fmt.Print("GetMessageRequestResponse:")
	fmt.Println(*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}