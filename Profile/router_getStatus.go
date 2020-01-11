package Profile

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	// "strings"
	"encoding/json"
   util "miti-microservices/Util"
)

func GetStatus(w http.ResponseWriter, r *http.Request){
	header:=GetStatusHeader{}
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
	getStatusRequest:=GetStatusRequest{}
	profileRequestErr:=json.Unmarshal(requestBody,&getStatusRequest)
	if profileRequestErr!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}
	chatId:=getStatusRequest.ChatId
	authorization:=IsChatIdOfUser(userId,chatId)
	if(authorization=="Error"){
		util.Message(w,5001)
		return
	}

	statusResponse:=GetStatusDB(chatId)
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&GetStatusResponse{Code:200,Message:msg,ChatId:chatId,StatusList:statusResponse}
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}