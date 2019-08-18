package main 
import(
	"net/http"
	"fmt"
	// CD "app/Model/CreateDatabase"
	database "app/Model/UseDatabase"
	util "app/Utility"
	"io/ioutil"
	"encoding/json"
)


type GetChat_header struct{
	Cookie string `header:"Miti-Cookie"`
}

type Chat struct{
	Chat_id string `json:"chat_id"`
	Offset int `json:"offset"`
	Num_of_chat int `json:"num_of_chat"`
}
func getAllChat(w http.ResponseWriter, r *http.Request){
	getChat_header:=GetChat_header{}
	util.GetHeader(r,&getChat_header)
	session_id:=getChat_header.Cookie
	user_id,getChat_status:=database.Get_user_id_from_session(session_id)
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

	chat_data:=Chat{}
	err_user_data:=json.Unmarshal(requestBody,&chat_data)
	if err_user_data!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}

	status:=database.Check_correct_chat(user_id,chat_data.Chat_id)
	if status=="ERROR"{
		util.Message(w,1002)
		return
	}

	chat:=database.GetChatMessages(chat_data.Chat_id,chat_data.Offset,chat_data.Num_of_chat)

	util.SendChat(w,chat)
}