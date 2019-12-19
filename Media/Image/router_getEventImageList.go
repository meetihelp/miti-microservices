package Image

import(
	"net/http"
	"fmt"
	// CD "miti-microservices/Model/CreateDatabase"
	util "miti-microservices/Util"
	"io/ioutil"
	"encoding/json"
)

func GetEventImageList(w http.ResponseWriter, r *http.Request){
	getEventImageListHeader:=GetEventImageListHeader{}
	util.GetHeader(r,&getEventImageListHeader)
	sessionId:=getEventImageListHeader.Cookie
	_,getChatStatus:=util.GetUserIdFromSession(sessionId)
	// fmt.Println(EventId)
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

	getEventImageListData :=GetEventImageListDS{}
	errEventData:=json.Unmarshal(requestBody,&getEventImageListData)
	if errEventData!=nil{
		fmt.Println("Could not Unmarshall Event data")
		util.Message(w,1001)
		return 
	}

	imageList:=GetEventImageListDB(getEventImageListData.EventId)
	SendImageList(w,imageList)
}