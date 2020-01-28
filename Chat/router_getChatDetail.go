package Chat
import(
	"net/http"
	"fmt"
	// CD "miti-microservices/Model/CreateDatabase"
	util "miti-microservices/Util"
	database "miti-microservices/Database"
	"io/ioutil"
	"encoding/json"
)
func GetChatDetailroute(w http.ResponseWriter, r *http.Request){
	getChatDetailHeader:=GetChatDetailHeader{}
	util.GetHeader(r,&getChatDetailHeader)
	fmt.Print("GetChatDetailHeader:")
	fmt.Println(getChatDetailHeader)
	sessionId:=getChatDetailHeader.Cookie
	db:=database.DBConnection()
	userId,getChatStatus:=util.GetUserIdFromSession2(db,sessionId)
	fmt.Println(userId)
	if getChatStatus=="Error"{
		util.Message(w,1003)
		db.Close()
		return
	}


	//Read body data
	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body for GetChatDetail")
		util.Message(w,1000)
		db.Close()
		return 
	}

	chatDetailDs :=ChatDetailDs{}
	errUserData:=json.Unmarshal(requestBody,&chatDetailDs)
	if errUserData!=nil{
		fmt.Println("Could not Unmarshall user data for GetChatDetail")
		util.Message(w,1001)
		db.Close()
		return 
	}

	sanatizationStatus :=Sanatize(chatDetailDs)
	if sanatizationStatus =="Error"{
		fmt.Println("User data invalid for GetChatDetail")
		util.Message(w,1002)
		db.Close()
		return
	}

	userId2,chatDetail,chatDetailErr:=GetChatDetail(db,userId,chatDetailDs.CreatedAt,chatDetailDs.NumOfChat)
	if chatDetailErr=="Error"{
		SendChatDetail(w,chatDetail,userId2,7000)
		db.Close()
		return
	}else{
		SendChatDetail(w,chatDetail,userId2,200)
		db.Close()
		return
	}
}