package Util

import(
	"net/http"
	"encoding/json"
	"bytes"
)


type SessionErrorDS struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	MoveTo int `json:"MoveTo"`
}

type SessionErrorHeader struct{
	ContentType string `json:"Content-Type"`
}
func GetSessionErrorContent(w http.ResponseWriter) (SessionErrorDS,http.ResponseWriter){
	statusCode:=1003
	moveTo:=2
	var data map[string]string

	
	sessionError:=SessionErrorDS{}
	sessionError.Code=statusCode
	sessionError.Message=GetMessageDecode(statusCode)
	sessionError.MoveTo=moveTo

	responseHeader:=SessionErrorHeader{}
	responseHeader.ContentType="application/json"
	headerBytes:=new(bytes.Buffer)
	json.NewEncoder(headerBytes).Encode(responseHeader)
	responseHeaderBytes:=headerBytes.Bytes()
	if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
    	panic(err)
	}
	w=GetResponseFormatHeader(w,data)
	return sessionError,w
}