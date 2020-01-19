package Social

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	// "strings"
	"encoding/json"
   util "miti-microservices/Util"
   gps "miti-microservices/GPS"
   profile "miti-microservices/Profile"
)

func GroupPoolStatusRouter(w http.ResponseWriter, r *http.Request){
	header:=GroupPoolStatusHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie
	userId,dErr:=util.GetUserIdFromSession(sessionId)
	fmt.Print("GroupPoolStatusHeader:")
	fmt.Println(header)
	if dErr=="Error"{
		fmt.Println("Session Does not exist GroupPoolStatusRouter")
		util.Message(w,1003)
		return
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body GroupPoolStatusRouter")
		util.Message(w,1000)
		return 
	}

	groupPoolStatusRequest:=GroupPoolStatusRequest{}
	errQuestionData:=json.Unmarshal(requestBody,&groupPoolStatusRequest)
	if errQuestionData!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}
	fmt.Print("Get Group Pool Status Body:")
	fmt.Println(groupPoolStatusRequest)
	latitude:=groupPoolStatusRequest.Latitude
	longitude:=groupPoolStatusRequest.Longitude
	gps.UpdateUserCurrentLocation(userId,latitude,longitude)

	interest,groupPoolStatus:=GroupPoolStatusDB(userId)
	w.Header().Set("Content-Type", "application/json")
	// status:=groupPoolStatus.Status
	// createdAt:=groupPoolStatus.CreatedAt
	// chatId:=groupPoolStatus.ChatId
	ipip:=profile.CheckIPIPStatus(userId)
	code:=200
	if(ipip<5){
		code=2003
	}
	msg:=util.GetMessageDecode(code)
	p:=&GroupPoolStatusResponse{Code:code,Message:msg,Interest:interest,Status:groupPoolStatus,IPIP:ipip}
	fmt.Print("GroupPoolStatusResponse:")
	fmt.Println(*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}