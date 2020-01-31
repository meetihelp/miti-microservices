package Chat

import (
	"fmt"
	"net/http"
	database "miti-microservices/Database"
	util "miti-microservices/Util"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"log"
)

func ChatInsert(w http.ResponseWriter,r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	chatHeader:=ChatHeader{}

	content:=ChatResponse{}
	statusCode:=0

	responseHeader:=ChatResponseHeader{}
	var data map[string]string

	db:=database.DBConnection()
	//Session,TemporarySession,Body,Unmarshal,Sanatize,Database
	list:=[]bool{false,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	util.GetHeader(r,&chatHeader)
	sessionId:=chatHeader.Cookie
	userId,loginStatus,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	util.APIHitLog("ChatInsert",ipAddress,sessionId)

	if (loginStatus=="Error"){
		fmt.Println("ChatInsert line 36")
		errorList.SessionError=true
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	// errorStatus:=util.ErrorListStatus(errorList)
	if (err!=nil && !util.ErrorListStatus(errorList)){
		fmt.Println("ChatInsert line 43")
		errorList.BodyReadError=true
	}

	chatData :=ChatRequest{}
	errorStatus:=util.ErrorListStatus(errorList)
	if(!errorStatus){
		fmt.Println("ChatInsert line 50")
		errUserData:=json.Unmarshal(requestBody,&chatData)
		if(errUserData!=nil){
			errorList.UnmarshallingError=true
		}	
	}

	errorStatus=util.ErrorListStatus(errorList)
	if(!errorStatus){
		fmt.Println("ChatInsert line 59")
		util.BodyLog("ChatInsert",ipAddress,sessionId,chatData)
		sanatizationStatus :=Sanatize(chatData)
		if(sanatizationStatus=="Error"){
			errorList.SanatizationError=true
		}		
	}

	chat:=Chat{}
	chat.UserId=userId
	chat.MessageId=util.GenerateToken()
	chat.CreatedAt=util.GetTime()
	chat.ChatId=chatData.ChatId
	chat.MessageType=chatData.MessageType
	chat.MessageContent=chatData.MessageContent
	chat.RequestId=chatData.RequestId
	lastUpdate:=chatData.CreatedAt

	errorStatus=util.ErrorListStatus(errorList)
	var unSyncedChat []Chat
	var chatResponse Chat
	if(!errorStatus){
		fmt.Println("ChatInsert line 81")
		chatResponse,unSyncedChat,dbError=ChatInsertDB(db,chat,lastUpdate)
		errorList.DatabaseError=dbError
		if(chat.CreatedAt==chatResponse.CreatedAt && !dbError){
			fmt.Println("ChatInsert line 85")
			dbError:=UpdateChatTime(db,chatData.ChatId,chatData.CreatedAt)
			errorList.DatabaseError=dbError
		}
	}else{
		statusCode=1002
	}
	
	code:=util.GetCode(errorList)
	if(code==200){
		fmt.Println("ChatInsert line 95")
		content.Code=statusCode
	}else{
		fmt.Println("ChatInsert line 98")
		content.Code=code
	}
	content.Message=util.GetMessageDecode(code)
	content.MessageId=chatResponse.MessageId
	content.RequestId=chatResponse.RequestId
	content.CreatedAt=chatResponse.CreatedAt
	content.MessageType=chatResponse.MessageType
	content.Chat=unSyncedChat

	responseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(responseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
	p:=&content
	util.ResponseLog("ChatInsert",ipAddress,sessionId,content.Code,*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}
