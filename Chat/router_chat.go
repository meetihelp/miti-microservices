package Chat

import (
	"net/http"
	"fmt"
	// redis "miti-microservices/Model/Redis"
	// database "miti-microservices/Database"
	util "miti-microservices/Util"
	// auth "miti-microservices/Authentication"
	"io/ioutil"
	"encoding/json"
	"log"
	// "time"
)

func ChatInsert(w http.ResponseWriter,r *http.Request){
	chatHeader:=ChatHeader{}
	util.GetHeader(r,&chatHeader)
	sessionId:=chatHeader.Cookie
	userId,loginStatus:=util.GetUserIdFromSession(sessionId)

	if loginStatus=="Error"{
		// util.Message(w,1003)
		// return
		content,w:=util.GetSessionErrorContent(w)
		p:=&content
		enc := json.NewEncoder(w)
		err:= enc.Encode(p)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		// fmt.Println("Could not read body")
		// util.Message(w,1000)
		// return 
		content,w:=util.GetBodyReadErrorContent(w)
		p:=&content
		enc := json.NewEncoder(w)
		err:= enc.Encode(p)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	chatData :=SendChatRequest{}
	errUserData:=json.Unmarshal(requestBody,&chatData)
	if errUserData!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}

	// sanatizationStatus :=Sanatize(chatData)
	// if sanatizationStatus =="Error"{
	// 	fmt.Println("User data invalid")
	// 	util.Message(w,1002)
	// 	return
	// }
	// tempUserId:=GetTempUserIdFromChatId(userId,chatData.ChatId)
	// chatData.UserId=tempUserId
	chat:=Chat{}
	chat.UserId=userId
	chat.MessageId=util.GenerateToken()
	chat.CreatedAt=util.GetTime()
	chat.ChatId=chatData.ChatId
	chat.MessageType=chatData.MessageType
	chat.MessageContent=chatData.MessageContent
	chat.RequestId=chatData.RequestId
	lastUpdate:=chatData.CreatedAt
	// tempTime:=time.Now()
	// chatData.CreatedAt=tempTime.Format("2006-01-02 15:04:05")
	// index:=GetLastChatIndex(chatData.ChatId)
	// index=util.GetNextLexString(index)
	// index=index+1
	// chatData.Index=index
	// fmt.Println(chatData.CreatedAt)
	// db:=database.GetDB()
	if(chatData.MessageContent!=""){
		chatResponse,unSyncedChat:=ChatInsertDB(chat,lastUpdate)
	// db.Create(&chatData)
		if(chat.CreatedAt==chatResponse.CreatedAt){
			e:=UpdateChatTime(chatData.ChatId,chatData.CreatedAt)
			if e!=nil{
				return
			}
		}
		
		// userList:=GetUserListFromChatId(chatData.ChatId)
		// EnterReadBy(userList,chatData.MessageId)
		// util.Message(w,200)
		// fmt.Print("Unsynced Chat:")
		// fmt.Println(unSyncedChat)
		SendMessageResponse(w,chatResponse.RequestId,chatResponse.MessageId,chatResponse.CreatedAt,chatResponse.MessageType,unSyncedChat)
	}else{
		util.Message(w,1002)
	}
}
