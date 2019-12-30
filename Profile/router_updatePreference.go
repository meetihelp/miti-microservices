package Profile

import(
	"fmt"
	"net/http"
	// "log"
	"io/ioutil"
	// "strings"
	"encoding/json"
   util "miti-microservices/Util"
   auth "miti-microservices/Authentication"
)

func UpdatePreference(w http.ResponseWriter, r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	header:=UpdatePreferenceResponseHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie

	userId,dErr:=util.GetUserIdFromSession(sessionId)
	// if dErr=="Error"{
	// 	fmt.Println("Session Does not exist")
	// 	util.Message(w,1003)
	// 	return
	// }
	if dErr=="Error"{
		userId,dErr=util.GetUserIdFromTemporarySession(sessionId)
		if dErr=="Error"{
			util.Message(w,1003)
			return
		}
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}
	
	var data map[string]string
	// data:=QuestionResponse{}
	if err := json.Unmarshal(requestBody, &data); err != nil {
        panic(err)
    }
    
	preferenceStatus:=UpdatePreferecePResponseDB(userId,data)
	auth.UpdatePreferencetatus(userId,preferenceStatus)
	// UpdateIPIPScore(userId)
	if(preferenceStatus>=6){
		util.InsertSessionValue(sessionId,userId,ipAddress)
        util.DeleteTemporarySession(sessionId)
	}
	util.Message(w,200)
}

func getDataInInterestForm(interest Interest,data map[string]string) (int,Interest){
	preferenceStatus:=0
	for key,value:=range data{
		switch key{
		case "InterestIndoorPassive1":
			preferenceStatus=1
			interest.InterestIndoorPassive1=value
		case "InterestIndoorPassive2":
			preferenceStatus=1
			interest.InterestIndoorPassive2=value
		case "InterestIndoorActive1":
			preferenceStatus=2
			interest.InterestIndoorActive1=value
		case "InterestIndoorActive2":
			preferenceStatus=2
			interest.InterestIndoorActive2=value
		case "InterestOutdoorPassive1":
			preferenceStatus=3
			interest.InterestOutdoorPassive1=value
		case "InterestOutdoorPassive2":
			preferenceStatus=3
			interest.InterestOutdoorPassive2=value
		case "InterestOutdoorActive1":
			preferenceStatus=4
			interest.InterestOutdoorActive1=value
		case "InterestOutdoorActive2":
			preferenceStatus=4
			interest.InterestOutdoorActive2=value
		case "InterestOthers1":
			preferenceStatus=5
			interest.InterestOthers1=value
		case "InterestOthers2":
			preferenceStatus=5
			interest.InterestOthers2=value
		case "InterestIdeology1":
			preferenceStatus=6
			interest.InterestIdeology1=value
		case "InterestIdeology2":
			preferenceStatus=6
			interest.InterestIdeology2=value
		}
	}

	return preferenceStatus,interest
}