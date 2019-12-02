package Profile

import(
	"net/http"
	// "io/ioutil"
	"encoding/json"	
	util "app/Util"
	"log"
	// "reflect"
	// "fmt"
)

func SendQuestion(w http.ResponseWriter,question []Question){
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&SendQuestionContent{Code:200,Message:msg,Question:question}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}

func SendProfile(w http.ResponseWriter,profileResponse ProfileResponse){
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=ProfileResponseContent{Code:200,Message:msg,ProfileResponse:profileResponse}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}
