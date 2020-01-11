package Profile

import(
	"fmt"
	"net/http"
	"log"
	// "io/ioutil"
	// "strings"
	"encoding/json"
   util "miti-microservices/Util"
)

func GetCheckInterestRouter(w http.ResponseWriter, r *http.Request){
	header:=GetCheckInterestHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie
	userId,dErr:=util.GetUserIdFromSession(sessionId)
	if dErr=="Error"{
		fmt.Println("Session Does not exist")
		util.Message(w,1003)
		return
	}

	// requestBody,err:=ioutil.ReadAll(r.Body)
	// if err!=nil{
	// 	fmt.Println("Could not read body")
	// 	util.Message(w,1000)
	// 	return 
	// }
	
	
	response:=GetCheckInterestDB(userId)
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&GetCheckInterestResponse{Code:200,Message:msg,List:response}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}