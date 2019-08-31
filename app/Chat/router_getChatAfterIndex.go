package Chat 
import(
	"net/http"
	"fmt"
	// redis "app/Model/Redis"
	util "app/Util"
	"io/ioutil"
	"encoding/json"
)

func GetChatAfterIndex(w http.ResponseWriter,r *http.Request){
	getChatAfterIndex_header:=GetChatAfterIndex_header{}
	util.GetHeader(r,&getChatAfterIndex_header)
	session_id:=getChatAfterIndex_header.Cookie
	user_id,getChat_status:=util.Get_user_id_from_session(session_id)
	fmt.Println(user_id)
	if getChat_status=="ERROR"{
		util.Message(w,1003)
		return
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}

	chat_data:=ChatAfterIndex{}
	err_user_data:=json.Unmarshal(requestBody,&chat_data)
	if err_user_data!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}

	status:=Check_correct_chat(user_id,chat_data.Chat_id)
	if status=="ERROR"{
		util.Message(w,1002)
		return
	}

	chat:=GetChatAfterIndexMessages(chat_data.Chat_id,chat_data.Offset,chat_data.Num_of_chat,chat_data.Index)

	SendChat(w,chat)
}