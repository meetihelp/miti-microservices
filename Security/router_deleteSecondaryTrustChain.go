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

func DeleteSecondaryTrustChain(w http.ResponseWriter, r *http.Request){
	header:=DeleteSecondaryTrustChainHeader{}
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
	
	deleteSecondaryTrustChainRequest:=DeleteSecondaryTrustChainRequest{}
	profileRequestErr:=json.Unmarshal(requestBody,&deleteSecondaryTrustChainRequest)
	if profileRequestErr!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}
	// primaryTrustChain.UserId=userId
	// DeletePrimaryTrustChain(primaryTrustChain)
	// util.Message(w,200)

	requestId:=deleteSecondaryTrustChainRequest.RequestId
	chatId:=deleteSecondaryTrustChainRequest.ChatId
	DeleteSecondaryTrustChainDB(userId,chatId)
	// updatedAt=DeletePrimaryTrustChainDB(userId,id,requestId,updatedAt)

	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&DeleteSecondaryTrustChainResponse{Code:200,Message:msg,RequestId:requestId}
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}