package Chat
import(
	"net/http"
	// "io"
	"encoding/json"	
	util "app/Util"
	"log"
)
func SendChat(w http.ResponseWriter,chat []Chat){
	w.Header().Set("Content-Type", "application/json")
	msg:=util.Get_message_decode(200)
	p:=&SendChat_Content{Code:200,Message:msg,Chat:chat}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}

func Send_ChatDetail(w http.ResponseWriter, chatDetail []ChatDetail,status_code int){
	w.Header().Set("Content-Type", "application/json")
	msg:=util.Get_message_decode(status_code)
	p:=&ChatDetail_Content{ChatDetail:chatDetail,Code:status_code,Message:msg}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}