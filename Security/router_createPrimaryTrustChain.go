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

func CreatePrimaryTrustChain(w http.ResponseWriter, r *http.Request){
	header:=CreatePrimaryTrustChainHeader{}
	util.GetHeader(r,&header)

	sessionId:=header.Cookie
	userId,dErr:=util.GetUserIdFromSession(sessionId)

	fmt.Print("CreatePrimaryTrustChainHeader:")
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
	
	primaryTrustChainRequest:=CreatePrimaryTrustChainRequest{}
	profileRequestErr:=json.Unmarshal(requestBody,&primaryTrustChainRequest)
	if profileRequestErr!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}

	fmt.Print("CreatePrimaryTrustChain Body:->")
	fmt.Println(primaryTrustChainRequest)
	// primaryTrustChainRequest.UserId=userId
	// if(primaryTrustChainRequest.Id>6 || primaryTrustChainRequest.Id<1){
	// 	util.Message(w,1002)
	// 	return
	// }
	requestId:=primaryTrustChainRequest.RequestId
	updatedAt:=util.GetTime()
	chainId:=primaryTrustChainRequest.ChainId
	phone:=primaryTrustChainRequest.Phone
	name:=primaryTrustChainRequest.Name
	updatedAt=UpdatePrimaryTrustChain(userId,chainId,phone,name,requestId,updatedAt)

	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&CreatePrimaryTrustChainResponse{Code:200,Message:msg,RequestId:requestId,UpdatedAt:updatedAt}
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}