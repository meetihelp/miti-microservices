package Chat
import(
	"net/http"
	"fmt"
	// CD "miti-microservices/Model/CreateDatabase"
	util "miti-microservices/Util"
	"io/ioutil"
	"encoding/json"
)
func GetChatDetailroute(w http.ResponseWriter, r *http.Request){
	getChatDetailHeader:=GetChatDetailHeader{}
	util.GetHeader(r,&getChatDetailHeader)
	sessionId:=getChatDetailHeader.Cookie
	userId,getChatStatus:=util.GetUserIdFromSession(sessionId)
	fmt.Println(userId)
	if getChatStatus=="Error"{
		util.Message(w,1003)
		return
	}


	//Read body data
	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}

	chatDetailDs :=ChatDetailDs{}
	errUserData:=json.Unmarshal(requestBody,&chatDetailDs)
	if errUserData!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}

	sanatizationStatus :=Sanatize(chatDetailDs)
	if sanatizationStatus =="Error"{
		fmt.Println("User data invalid")
		util.Message(w,1002)
		return
	}

	userId2,chatDetail,chatDetailErr:=GetChatDetail(userId,chatDetailDs.CreatedAt,chatDetailDs.NumOfChat)
	if chatDetailErr=="Error"{
		SendChatDetail(w,chatDetail,userId2,7000)
		return
	}else{
		SendChatDetail(w,chatDetail,userId2,200)
		return
	}
}