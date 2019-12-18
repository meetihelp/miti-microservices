package Event
import(
	"net/http"
	"fmt"
	// "strconv"
	// CD "miti-microservices/Model/CreateDatabase"
	util "miti-microservices/Util"
	"io/ioutil"
	"encoding/json"
	// gps "miti-microservices/GPS"
)

func GetEventById(w http.ResponseWriter, r *http.Request){
	getEventByIdHeader:=GetEventByIdHeader{}
	util.GetHeader(r,&getEventByIdHeader)
	sessionId:=getEventByIdHeader.Cookie
	_,getChatStatus:=util.GetUserIdFromSession(sessionId)
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

	getEventByIdData :=GetEventByIdDS{}
	errUserData:=json.Unmarshal(requestBody,&getEventByIdData)
	if errUserData!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}

	event,status:=GetEventByIdDB(getEventByIdData.EventId)
	if status=="Ok"{
		SendEvent(w,event)
	}else{
		util.Message(w,8001)
	}

}