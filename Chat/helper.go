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
		// temp.UserId=c.TempUserId
		temp.UserId=c.ActualUserId
		temp.ChatId=c.ChatId
		temp.ChatType=c.ChatType
		temp.CreatedAt=c.CreatedAt
		temp.LastUpdate=c.LastUpdate
		temp.Name=c.Name
		chatDetailResponse=append(chatDetailResponse,temp)
	}
	p:=&ChatDetailContent{ChatDetailResponse:chatDetailResponse,Code:statusCode,Message:msg}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}

func SendMessageResponse(w http.ResponseWriter,requestId string, messageId string, createdAt string,messageType string){
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&ChatResponse{Code:200,Message:msg,RequestId:requestId,MessageId:messageId,CreatedAt:createdAt,MessageType:messageType}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}

}

type AnonymousUserHelper struct{
	UserId string `gorm:"primary_key;varchar(100)"  json:"UserId"`
	AnonymousId string `gorm:"primary_key;unique;varchar(100)"  json:"AnonymousId"`
	ChatId string `gorm:"primary_key;varchar(100)"  json:"ChatId"`
	Status string `gorm:"type:varchar(6)" json:"Status"`  // status for Liked/not liked/ none
	CreatedAt string `gorm:"type:varchar(100)" json:"CreatedAt"`
}