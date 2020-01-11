package Security

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	// "strings"
	"encoding/json"
   util "miti-microservices/Util"
)

func CreateSecondaryTrustChain(w http.ResponseWriter, r *http.Request){
	header:=CreateSecondaryTrustChainHeader{}
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
	
	secondaryTrustChainRequest:=CreateSecondaryTrustChainRequest{}
	profileRequestErr:=json.Unmarshal(requestBody,&secondaryTrustChainRequest)
	if profileRequestErr!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}
	
	requestId:=secondaryTrustChainRequest.RequestId
	chatId:=secondaryTrustChainRequest.ChatId
	createdAt:=util.GetTime()
	createdAt=InsertSecondaryTrustChain(userId,chatId,requestId,createdAt)
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&CreateSecondaryTrustChainResponse{Code:200,Message:msg,RequestId:requestId,CreatedAt:createdAt}
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}