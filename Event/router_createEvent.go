package Event

// import(
// 	"net/http"
// 	"fmt"
// 	// "strconv"
// 	// CD "miti-microservices/Model/CreateDatabase"
// 	util "miti-microservices/Util"
// 	"io/ioutil"
// 	"encoding/json"
// 	gps "miti-microservices/GPS"
// )

// func CreateEvent(w http.ResponseWriter, r *http.Request){
// 	createEventHeader:=CreateEventHeader{}
// 	util.GetHeader(r,&createEventHeader)
// 	sessionId:=createEventHeader.Cookie
// 	organiserId,getChatStatus:=util.GetUserIdFromSession(sessionId)
// 	if getChatStatus=="Error"{
// 		util.Message(w,1003)
// 		return
// 	}

// 	//Read body data
// 	requestBody,err:=ioutil.ReadAll(r.Body)
// 	if err!=nil{
// 		fmt.Println("Could not read body")
// 		util.Message(w,1000)
// 		return 
// 	}

// 	createEventData :=CreateEventDS{}
// 	errUserData:=json.Unmarshal(requestBody,&createEventData)
// 	if errUserData!=nil{
// 		fmt.Println("Could not Unmarshall user data")
// 		util.Message(w,1001)
// 		return 
// 	}
// 	event:=getEvent(createEventData)
// 	event.OrganiserId=organiserId
// 	eventId:=InsertEvent(event)
// 	status:=gps.InsertEventLocation(eventId,createEventData.EventType,createEventData.Latitude,createEventData.Longitude)
// 	if status=="Ok"{
// 		util.Message(w,200)	
// 	}else{
// 		util.Message(w,1006)
// 	}
	

// }

// func getEvent(data CreateEventDS) Event{
// 	event:=Event{}
// 	event.EventName=data.EventName
// 	event.EventType=data.EventType
// 	event.Time=data.Time
// 	return event
// }