package main 

import(
	"net/http"
	"fmt"
	// CD "app/Model/CreateDatabase"
	database "app/Model/UseDatabase"
	util "app/Utility"
	// "io/ioutil"
	// "encoding/json"
)


type GetChatDetail_header struct{
	Cookie string `header:"Miti-Cookie"`
}
func getChatDetail(w http.ResponseWriter, r *http.Request){
	getChatDetail_header:=GetChatDetail_header{}
	util.GetHeader(r,&getChatDetail_header)
	session_id:=getChatDetail_header.Cookie
	user_id,getChat_status:=database.Get_user_id_from_session(session_id)
	fmt.Println(user_id)
	if getChat_status=="ERROR"{
		util.Message(w,1003)
		return
	}

	chatDetail,err:=database.GetChatDetail(user_id)
	if err=="ERROR"{
		return
	}else{
		util.Send_ChatDetail(w,chatDetail)
		return
	}
}