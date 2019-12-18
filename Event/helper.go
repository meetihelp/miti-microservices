package Event

import(
	"net/http"
	// "fmt"
	"log"
	// "strconv"
	// CD "miti-microservices/Model/CreateDatabase"
	util "miti-microservices/Util"
	// "io/ioutil"
	"encoding/json"
)

func SendEvent(w http.ResponseWriter,event Event){
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	eventResponse:=EventResponse{}
	eventResponse.Code=200
	eventResponse.Message=msg
	eventResponse.EventId=event.EventId
	eventResponse.EventName=event.EventName
	eventResponse.EventPicURL=event.EventPicURL
	eventResponse.EventType=event.EventType
	eventResponse.Time=event.Time
	eventResponse.OrganiserId=event.OrganiserId
	eventResponse.Latitude=event.Latitude
	eventResponse.Longitude=event.Longitude
	p:=&eventResponse
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}

}