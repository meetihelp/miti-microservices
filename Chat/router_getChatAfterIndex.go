package Chat 
import(
	"net/http"
	"fmt"
	// redis "miti-microservices/Model/Redis"
	util "miti-microservices/Util"
	"io/ioutil"
	"encoding/json"
)

func GetChatAfterIndex(w http.ResponseWriter,r *http.Request){
	getChatAfterIndexHeader:=GetChatAfterIndexHeader{}
	util.GetHeader(r,&getChatAfterIndexHeader)
	sessionId:=getChatAfterIndexHeader.Cookie
	userId,getChatStatus:=util.GetUserIdFromSession(sessionId)
	fmt.Println(userId)
	fmt.Print("GetChatAfterIndexHeader:")
	fmt.Println(getChatAfterIndexHeader)
	if getChatStatus=="Error"{
		fmt.Println("Session Error for GetChatAfterIndex")
		util.Message(w,1003)
		return
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body for GetChatAfterIndex")
		util.Message(w,1000)
		return 
	}

	chatData:=ChatAfterTime{}
	errUserData:=json.Unmarshal(requestBody,&chatData)
	if errUserData!=nil{
		fmt.Println("Could not Unmarshall user data for GetChatAfterIndex")
		util.Message(w,1001)
		return 
	}
	sanatizationStatus :=Sanatize(chatData)
	if sanatizationStatus =="Error"{
		fmt.Println("User data invalid for GetChatAfterIndex")
		util.Message(w,1002)
		return
	}

	status:=CheckCorrectChat(userId,chatData.ChatId)
	if status=="Error"{
		util.Message(w,1002)
		return
	}

	chat:=GetChatAfterTimeMessages(chatData.ChatId,chatData.NumOfChat,chatData.CreatedAt)

	SendChat(w,chat)
}