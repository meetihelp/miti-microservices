package Chat
import(
	"net/http"
	// "io"
	"encoding/json"	
	util "miti-microservices/Util"
	"log"
)
func SendChat(w http.ResponseWriter,chat []Chat){
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&SendChatContent{Code:200,Message:msg,Chat:chat}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}

func SendChatDetail(w http.ResponseWriter, chatDetail []ChatDetail,statusCode int){
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(statusCode)
	p:=&ChatDetailContent{ChatDetail:chatDetail,Code:statusCode,Message:msg}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}