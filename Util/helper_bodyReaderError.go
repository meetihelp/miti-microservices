package Util

import(
	"net/http"
	"encoding/json"
	"bytes"
)


type BodyReadErrorDS struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	MoveTo int `json:"MoveTo"`
}

type BodyReadErrorHeader struct{
	ContentType string `json:"Content-Type"`
}
func GetBodyReadErrorContent(w http.ResponseWriter) (BodyReadErrorDS,http.ResponseWriter){
	statusCode:=1000
	moveTo:=0
	var data map[string]string

	
	bodyReadError:=BodyReadErrorDS{}
	bodyReadError.Code=statusCode
	bodyReadError.Message=GetMessageDecode(statusCode)
	bodyReadError.MoveTo=moveTo

	responseHeader:=BodyReadErrorHeader{}
	responseHeader.ContentType="application/json"
	headerBytes:=new(bytes.Buffer)
	json.NewEncoder(headerBytes).Encode(responseHeader)
	responseHeaderBytes:=headerBytes.Bytes()
	if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
    	panic(err)
	}
	w=GetResponseFormatHeader(w,data)
	return bodyReadError,w
}