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
	chatDetailResponse:=[]ChatDetailResponse{}
	for _,c:=range chatDetail{
		temp:=ChatDetailResponse{}
		temp.UserId=c.TempUserId
		temp.ChatId=c.ChatId
		temp.ChatType=c.ChatType
		temp.CreatedAt=c.CreatedAt
		temp.LastUpdate=c.LastUpdate
		chatDetailResponse=append(chatDetailResponse,temp)
	}
	p:=&ChatDetailContent{ChatDetailResponse:chatDetailResponse,Code:statusCode,Message:msg}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}

func SendMessageResponse(w http.ResponseWriter, messageId string, createdAt string){
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&ChatResponse{Code:200,Message:msg,MessageId:messageId,CreatedAt:createdAt}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}

}