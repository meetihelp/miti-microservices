package Profile

import(
	"fmt"
	"net/http"
	// "log"
	"io/ioutil"
	// "strings"
	"encoding/json"
   util "miti-microservices/Util"
)

func CheckInterestRouter(w http.ResponseWriter, r *http.Request){
	header:=CheckInterestHeader{}
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
	
	checkInterestRequest:=CheckInterestRequest{}
	profileRequestErr:=json.Unmarshal(requestBody,&checkInterestRequest)
	if profileRequestErr!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}
	
	checkInterest:=CheckInterest{}
	checkInterest.UserId1=userId
	checkInterest.UserId2=checkInterestRequest.UserId
	checkInterest.Interest=checkInterestRequest.Interest
	checkInterest.CreatedAt=util.GetTime()
	EnterCheckInterestDB(checkInterest)
	util.Message(w,200)
	// w.Header().Set("Content-Type", "application/json")
	// msg:=util.GetMessageDecode(200)
	// p:=&CreateStatusResponse{Code:200,Message:msg,RequestId:response.RequestId,CreatedAt:response.CreatedAt}
	// enc := json.NewEncoder(w)
	// err= enc.Encode(p)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}