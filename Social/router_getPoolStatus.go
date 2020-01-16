package Social

import(
	"fmt"
	"net/http"
	"log"
	// "io/ioutil"
	// "strings"
	"encoding/json"
   util "miti-microservices/Util"
   profile "miti-microservices/Profile"
)

func PoolStatusRouter(w http.ResponseWriter, r *http.Request){
	header:=PoolStatusHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie
	userId,dErr:=util.GetUserIdFromSession(sessionId)
	fmt.Print("PoolStatusHeader:")
	fmt.Println(header)
	if dErr=="Error"{
		fmt.Println("Session Does not exist PoolStatusRouter")
		util.Message(w,1003)
		return
	}

	poolStatus:=PoolStatusDB(userId)
	w.Header().Set("Content-Type", "application/json")
	matchUsedId:=poolStatus.MatchUserId
	status:=poolStatus.Status
	createdAt:=poolStatus.CreatedAt
	matchTime:=poolStatus.MatchTime
	chatId:=poolStatus.ChatId
	ipip:=profile.CheckIPIPStatus(userId)
	code:=200
	if(ipip<5){
		code=2003
	}
	msg:=util.GetMessageDecode(code)
	p:=&PoolStatusResponse{Code:code,Message:msg,ChatId:chatId,MatchUserId:matchUsedId,
			Status:status,CreatedAt:createdAt,MatchTime:matchTime,IPIP:ipip}
	fmt.Print("PoolStatusResponse:")
	fmt.Println(*p)
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}