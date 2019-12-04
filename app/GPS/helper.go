package GPS

import(
	"net/http"
	// "io"
	"encoding/json"	
	util "app/Util"
	"log"
	"strconv"
	"math"
)

func CalculateDistance(location1 Location,location2 Location) float64{
	lat1,_:=strconv.ParseFloat(location1.Latitude,64)
	long1,_:=strconv.ParseFloat(location1.Longitude,64)
	lat2,_:=strconv.ParseFloat(location2.Latitude,64)
	long2,_:=strconv.ParseFloat(location2.Longitude,64)
	cos:=math.Cos
	p:= 0.017453292519943295;
	distance:= 0.5- cos((lat2 - lat1) * p)/2 + cos(lat1 * p) * cos(lat2 * p) * (1 - cos((long2 - long1) * p))/2;
	distance=12742 * math.Asin(math.Sqrt(distance));
	return distance
}

func GetCity(location Location) string{
	return "Patna"
}	

func SendUserList(w http.ResponseWriter,userList []string){
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&SendUserListContent{Code:200,Message:msg,UserList:userList}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}

func SendEventList(w http.ResponseWriter,eventList []string){
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&SendEventListContent{Code:200,Message:msg,EventList:eventList}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}


// func GetEventList(location Location,eventType string,distance float64) ([]string){
// 	db:=database.GetDB()
// 	profile:=[]Event{}
// 	db.Where("event_type=?",eventType).Find(&event)
// 	var eventList []string
// 	for _,e:=range event{
// 		eventList=append(eventList,event.EventId)
// 	}
// 	return eventList
// }