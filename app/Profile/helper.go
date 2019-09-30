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
	msg:=util.Get_message_decode(200)
	p:=&SendQuestion_Content{Code:200,Message:msg,Question:question}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}

