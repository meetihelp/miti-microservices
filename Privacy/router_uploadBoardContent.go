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

func UploadBoardContent(w http.ResponseWriter, r *http.Request){
	header:=UploadBoardContentHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie
	userId,dErr:=util.GetUserIdFromSession(sessionId)
	fmt.Print("UploadBoardContent Header:->")
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
	
	uploadBoardContentRequest:=UploadBoardContentRequest{}
	profileRequestErr:=json.Unmarshal(requestBody,&uploadBoardContentRequest)
	if profileRequestErr!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}

	fmt.Print("UploadBoardContent Body:->")
	fmt.Println(uploadBoardContentRequest)
	date:=uploadBoardContentRequest.Date
	boardId:=uploadBoardContentRequest.BoardId
	// boardId:=GetBoardId(userId,date)
	CreateBoard(userId,date,boardId)
	requestId:=uploadBoardContentRequest.RequestId
	text:=uploadBoardContentRequest.ContentText
	imageId:=uploadBoardContentRequest.ContentImageId
	if(text=="" && imageId==""){
		util.Message(w,1002)
		return
	}
	contentId:=util.GenerateToken()
	createdAt:=util.GetTime()
	createdAt,contentId=EnterBoardContent(userId,boardId,text,imageId,contentId,requestId,createdAt)

	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&UploadBoardContentResponse{Code:200,Message:msg,RequestId:requestId,
			CreatedAt:createdAt,BoardId:boardId,ContentId:contentId}
	fmt.Print("UploadBoardContent Response:->")
	fmt.Println(*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}