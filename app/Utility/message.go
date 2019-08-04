package Utility

import(
	"net/http"
	// "io"
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