package Chat 
import(
	"net/http"
	"fmt"
	// redis "miti-microservices/Model/Redis"
	util "miti-microservices/Util"
	database "miti-microservices/Database"
	"io/ioutil"
	"encoding/json"
)

func GetChatAfterIndex(w http.ResponseWriter,r *http.Request){
	getChatAfterIndexHeader:=GetChatAfterIndexHeader{}
	util.GetHeader(r,&getChatAfterIndexHeader)
	sessionId:=getChatAfterIndexHeader.Cookie
	db:=database.DBConnection()
	userId,getChatStatus:=util.GetUserIdFromSession2(db,sessionId)
	fmt.Println(userId)
	fmt.Print("GetChatAfterIndexHeader:")
	fmt.Println(getChatAfterIndexHeader)
	if getChatStatus=="Error"{
		fmt.Println("Session Error for GetChatAfterIndex")
		util.Message(w,1003)
		db.Close()
		return
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body for GetChatAfterIndex")
		util.Message(w,1000)
		db.Close()
		return 
	}

	chatData:=ChatAfterTime{}
	errUserData:=json.Unmarshal(requestBody,&chatData)
	if errUserData!=nil{
		fmt.Println("Could not Unmarshall user data for GetChatAfterIndex")
		util.Message(w,1001)
		db.Close()
		return 
	}
	sanatizationStatus :=Sanatize(chatData)
	if sanatizationStatus =="Error"{
		fmt.Println("User data invalid for GetChatAfterIndex")
		util.Message(w,1002)
		db.Close()
		return
	}

	status:=CheckCorrectChat(db,userId,chatData.ChatId)
	if status=="Error"{
		util.Message(w,1002)
		db.Close()
		return
	}

	chat:=GetChatAfterTimeMessages(db,chatData.ChatId,chatData.NumOfChat,chatData.CreatedAt)

	SendChat(w,chat)
	db.Close()
}