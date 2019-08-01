package Utility

import(
	"net/http"
	"io"
)

func Message(w http.ResponseWriter,status_code int){
	io.WriteString(w,get_message_decode(status_code))
}