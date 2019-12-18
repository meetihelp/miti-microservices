package Chat
import(
	"net/http"
	"fmt"
	// redis "app/Model/Redis"
	util "app/Util"
	"io/ioutil"
	"encoding/json"
)

func GetChat(w http.ResponseWriter, r *http.Request){
	getChatHeader:=GetChatHeader{}
	util.GetHeader(r,&getChatHeader)
	sessionId:=getChatHeader.Cookie
	userId,getChatStatus:=util.GetUserIdFromSession(sessionId)
	fmt.Println(userId)
	if getChatStatus=="Error"{
		util.Message(w,1003)
		return
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}

	chatData:=ChatRequest{}
	errUserData:=json.Unmarshal(requestBody,&chatData)
	if errUserData!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}

	sanatizationStatus :=Sanatize(chatData)
	if sanatizationStatus =="Error"{
		fmt.Println("User data invalid")
		util.Message(w,1002)
		return
	}

	status:=CheckCorrectChat(userId,chatData.ChatId)
	if status=="Error"{
		util.Message(w,1002)
		return
	}

	chat:=GetChatMessages(chatData.ChatId,chatData.Offset,chatData.NumOfChat)

	SendChat(w,chat)
}