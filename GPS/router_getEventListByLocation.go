package GPS
// import(
// 	"net/http"
// 	"fmt"
// 	"strconv"
// 	// CD "miti-microservices/Model/CreateDatabase"
// 	util "miti-microservices/Util"
// 	"io/ioutil"
// 	"encoding/json"
// )

// func GetEventListByLocation(w http.ResponseWriter, r *http.Request){
// 	getEventListByLocationHeader:=GetUserListByLocationHeader{}
// 	util.GetHeader(r,&getEventListByLocationHeader)
// 	sessionId:=getEventListByLocationHeader.Cookie
// 	userId,getChatStatus:=util.GetUserIdFromSession(sessionId)
// 	fmt.Println(userId)
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

// 	getEventListByLocationData :=GetEventListByLocationDS{}
// 	errUserData:=json.Unmarshal(requestBody,&getEventListByLocationData)
// 	if errUserData!=nil{
// 		fmt.Println("Could not Unmarshall user data")
// 		util.Message(w,1001)
// 		return 
// 	}
// 	location:=Location{}
// 	location.Latitude=getEventListByLocationData.Latitude
// 	location.Longitude=getEventListByLocationData.Longitude
// 	distance,_:=strconv.ParseFloat(getEventListByLocationData.Distance,64)
// 	eventType:=getEventListByLocationData.EventType
// 	eventList:=GetEventListByLocationDB(eventType,location,distance)
// 	// fmt.Println(eventList)
// 	SendEventList(w,eventList)
// }