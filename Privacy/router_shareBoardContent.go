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

func ShareBoardContent(w http.ResponseWriter, r *http.Request){
	header:=ShareBoardContentHeader{}
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
	
	shareBoardContentRequest:=ShareBoardContentRequest{}
	profileRequestErr:=json.Unmarshal(requestBody,&shareBoardContentRequest)
	if profileRequestErr!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}
	
	boardId:=shareBoardContentRequest.BoardId
	accessType:=shareBoardContentRequest.AccessType
	accessRequestId:=shareBoardContentRequest.AccessRequestId
	contentId:=shareBoardContentRequest.ContentId

	accessUpdatedAt:=UpdateBoardContentSharePolicy(userId,contentId,boardId,accessType,accessRequestId)

	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&ShareBoardContentResponse{Code:200,Message:msg,AccessRequestId:accessRequestId,
				AccessUpdatedAt:accessUpdatedAt}
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}