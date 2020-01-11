package Social

import(
	"fmt"
	"net/http"
	"log"
	// "io/ioutil"
	// "strings"
	"encoding/json"
   util "miti-microservices/Util"
)

func PoolStatusRouter(w http.ResponseWriter, r *http.Request){
	header:=PoolStatusHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie
	userId,dErr:=util.GetUserIdFromSession(sessionId)
	if dErr=="Error"{
		fmt.Println("Session Does not exist")
		util.Message(w,1003)
		return
	}

	poolStatus:=PoolStatusDB(userId)
	w.Header().Set("Content-Type", "application/json")
	matchUsedId:=poolStatus.MatchUserId
	status:=poolStatus.Status
	createdAt:=poolStatus.CreatedAt
	matchTime:=poolStatus.MatchTime

	code:=200
	if(status=="Matched"){
		code=200
	}
	msg:=util.GetMessageDecode(code)
	p:=&PoolStatusResponse{Code:code,Message:msg,MatchUserId:matchUsedId,
			Status:status,CreatedAt:createdAt,MatchTime:matchTime}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}