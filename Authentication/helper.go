package Authentication

import(
	"net/http"
	// "io/ioutil"
	"encoding/json"	
	util "app/Util"
	"log"
	// "reflect"
	// "fmt"
)

func SendPreference(w http.ResponseWriter,preferenceCreationStatus int,code int){
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(code)
	p:=&PreferenceContent{Code:code,Message:msg,Preference:preferenceCreationStatus}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}