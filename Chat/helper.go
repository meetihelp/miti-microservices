package Chat
import(
	"net/http"
	// "io"
	"encoding/json"	
	util "miti-microservices/Util"
	"log"
	"fmt"
)
// func SendChat(w http.ResponseWriter,chat []Chat){
// 	w.Header().Set("Content-Type", "application/json")
// 	msg:=util.GetMessageDecode(200)
// 	p:=&SendChatContent{Code:200,Message:msg,Chat:chat}
// 	fmt.Print("GetChatResponse:")
// 	fmt.Println(*p)
// 	enc := json.NewEncoder(w)
// 	err:= enc.Encode(p)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func SendChatDetail(w http.ResponseWriter, chatDetail []ChatDetail,userId2 []string,statusCode int){
// 	w.Header().Set("Content-Type", "application/json")
// 	msg:=util.GetMessageDecode(statusCode)
// 	chatDetailResponse:=[]ChatDetailResponse{}
// 	for i,c:=range chatDetail{
// 		temp:=ChatDetailResponse{}
// 		// temp.UserId=c.TempUserId
// 		temp.UserId=c.ActualUserId
// 		temp.UserId2=userId2[i]
// 		temp.ChatId=c.ChatId
// 		temp.ChatType=c.ChatType
// 		temp.CreatedAt=c.CreatedAt
// 		temp.LastUpdate=c.LastUpdate
// 		temp.Name=c.Name
// 		chatDetailResponse=append(chatDetailResponse,temp)
// 	}
// 	p:=&ChatDetailContent{ChatDetailResponse:chatDetailResponse,Code:statusCode,Message:msg}
// 	fmt.Print("GetChatDetail Response:")
// 	fmt.Println(*p)
// 	enc := json.NewEncoder(w)
// 	err:= enc.Encode(p)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func SendMessageResponse(w http.ResponseWriter,code int,requestId string, messageId string, createdAt string,messageType string,chat []Chat){
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(code)
	p:=&ChatResponse{Code:code,Message:msg,RequestId:requestId,MessageId:messageId,CreatedAt:createdAt,MessageType:messageType,Chat:chat}
	fmt.Print("ChatResponse:")
	fmt.Println(*p)
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

type User struct{
	UserId string `json:"UserId"`
	Phone string `json:"Phone"`
	Email string  `json:"Email"`
	// Password string `gorm:"type:varchar(100)" validate:"required" json:"Password"`
	Status string `json:"Status"`  //Verified/Unverified/Deleted
	ProfileCreationStatus string `json:"ProfileCreationStatus"`
	PreferenceCreationStatus int `son:"PreferenceCreationStatus"`
	IPIPStatus int `json:"IPIPStatus"`
	CreatedAt string `json:"CreatedAt"`
}