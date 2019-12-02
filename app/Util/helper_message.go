package Util

import(
	"net/http"
	// "io"
	// CD "app/Model/CreateDatabase"
	"encoding/json"	
	"log"
)

type MessageContent struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
}
func Message(w http.ResponseWriter,statusCode int){
	// io.WriteString(w,get_message_decode(status_code))
	msg:=GetMessageDecode(statusCode)
	w.Header().Set("Content-Type", "application/json")
	p := &MessageContent {Code:statusCode,Message:msg}
	// content,_:=json.Marshal(&p)
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}


type MatchListContent struct{
	Code int `json:"Code"`
	MatchList []string `json:"AnonymousUser"`
	ChatIdList []string `json:"ChatId"`
}
func Send_Match_list(w http.ResponseWriter,matchList []string,chatIdList []string){
	w.Header().Set("Content-Type", "application/json")
	p := &MatchListContent {Code:5000,MatchList:matchList,ChatIdList:chatIdList}	
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}

