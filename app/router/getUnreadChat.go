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

type GetUnreadChat_header struct{
	Cookie string `header:"Miti-Cookie"`
}


func getUnreadChat(w http.ResponseWriter, r *http.Request){
	getUnreadChat_header:=GetUnreadChat_header{}
	util.GetHeader(r,&getUnreadChat_header)
	session_id:=getUnreadChat_header.Cookie
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

	message_id:=database.GetUnreadMessage(user_id)
	chat:=database.GetChatByMessageId(message_id)

	util.SendChat(w,chat)
}