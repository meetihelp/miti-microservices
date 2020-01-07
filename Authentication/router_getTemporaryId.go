package Authentication

import(
	"fmt"
	"net/http"
	"log"
	// "io/ioutil"
	// "strings"
	"encoding/json"
   util "miti-microservices/Util"
)

func GetTemporaryUserId(w http.ResponseWriter, r *http.Request){
	//CHECK IF USER IS AUTHORIZED TO SEE THE PROFILE
	header:=GetTempUserIdHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie
	userId,dErr:=util.GetUserIdFromSession(sessionId)
	if dErr=="Error"{
		fmt.Println("Session Does not exist")
		util.Message(w,1003)
		return
	}

	// temporaryIdList:=GetTemporaryIdList(userId)
	// SendTemporaryIdList(temporaryIdList)
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	// p:=&TempUserResponse{Code:200,Message:msg,List:temporaryIdList}
	p:=&TempUserResponse{Code:200,Message:msg,UserId:userId}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}

}