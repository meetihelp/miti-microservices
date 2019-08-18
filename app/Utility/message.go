package Utility

import(
	"net/http"
	// "io"
	CD "app/Model/CreateDatabase"
	"encoding/json"	
	"log"
)

type Message_Content struct{
	Code int `json:"code"`
	Message string `json:"message"`
}
func Message(w http.ResponseWriter,status_code int){
	// io.WriteString(w,get_message_decode(status_code))
	msg:=get_message_decode(status_code)
	w.Header().Set("Content-Type", "application/json")
	p := &Message_Content {Code:status_code,Message:msg}
	// content,_:=json.Marshal(&p)
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}


type MatchList_Content struct{
	Code int `json:"code"`
	MatchList []string `json:"anonymousUser"`
	ChatIdList []string `json:"chatId"`
}
func Send_Match_list(w http.ResponseWriter,matchList []string,chatIdList []string){
	w.Header().Set("Content-Type", "application/json")
	p := &MatchList_Content {Code:5000,MatchList:matchList,ChatIdList:chatIdList}	
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}

type SendChat_Content struct{
	Code int `json:"code"`
	Chat []CD.Chat `json:"chat"`
}
func SendChat(w http.ResponseWriter,chat []CD.Chat){
	w.Header().Set("Content-Type", "application/json")
	p:=&SendChat_Content{Code:6000,Chat:chat}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}

}

type ChatDetail_Content struct{
	ChatDetail []CD.ChatDetail `json:chatdetail`
}
func Send_ChatDetail(w http.ResponseWriter, chatDetail []CD.ChatDetail){
	w.Header().Set("Content-Type", "application/json")
	p:=&ChatDetail_Content{ChatDetail:chatDetail}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}