package Privacy

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	// "strings"
	"encoding/json"
   util "miti-microservices/Util"
)

func ShareBoard(w http.ResponseWriter, r *http.Request){
	header:=ShareBoardHeader{}
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
	
	shareBoardRequest:=ShareBoardRequest{}
	profileRequestErr:=json.Unmarshal(requestBody,&shareBoardRequest)
	if profileRequestErr!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}
	
	boardId:=shareBoardRequest.BoardId
	accessType:=shareBoardRequest.AccessType
	requestId:=shareBoardRequest.RequestId

	updatedAt:=UpdateBoardSharePolicy(userId,boardId,accessType,requestId)

	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&ShareBoardResponse{Code:200,Message:msg,RequestId:requestId,UpdatedAt:updatedAt}
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}