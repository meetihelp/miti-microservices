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


type GetChatDetail_header struct{
	Cookie string `header:"Miti-Cookie"`
}

type ChatDetailDs struct{
	Offset int `json:"offset"`
	Num_of_chat int `json:"num_of_chat"`
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


	//Read body data
	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}

	chatDetailDs :=ChatDetailDs{}
	err_user_data:=json.Unmarshal(requestBody,&chatDetailDs)
	if err_user_data!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}

	chatDetail,chatDetail_err:=database.GetChatDetail(user_id,chatDetailDs.Offset,chatDetailDs.Num_of_chat)
	if chatDetail_err=="ERROR"{
		util.Send_ChatDetail(w,chatDetail,7000)
		return
	}else{
		util.Send_ChatDetail(w,chatDetail,200)
		return
	}
}