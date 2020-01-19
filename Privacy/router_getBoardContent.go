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

func GetBoardContent(w http.ResponseWriter, r *http.Request){
	header:=GetBoardContentHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie
	userId,dErr:=util.GetUserIdFromSession(sessionId)
	fmt.Print("GetBoardContent Header:->")
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
	
	getBoardContentRequest:=GetBoardContentRequest{}
	profileRequestErr:=json.Unmarshal(requestBody,&getBoardContentRequest)
	if profileRequestErr!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}

	fmt.Print("GetBoardContent Body:->")
	fmt.Println(getBoardContentRequest)
	
	createdAt:=getBoardContentRequest.CreatedAt

	boardContentList:=GetBoardContentDB(userId,createdAt)
	
	code:=200
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(code)
	p:=&GetBoardContentResponse{Code:code,Message:msg,BoardContentList:boardContentList}
	fmt.Print("UploadBoardContent Response:->")
	fmt.Println(*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}