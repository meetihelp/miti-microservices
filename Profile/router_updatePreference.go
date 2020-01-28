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
   database "miti-microservices/Database"
)

func UpdatePreference(w http.ResponseWriter, r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	header:=UpdatePreferenceResponseHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie
	db:=database.DBConnection()
	userId,dErr:=util.GetUserIdFromSession2(db,sessionId)
	// if dErr=="Error"{
	// 	fmt.Println("Session Does not exist")
	// 	util.Message(w,1003)
	// 	return
	// }
	if dErr=="Error"{
		userId,dErr=util.GetUserIdFromTemporarySession2(db,sessionId)
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
	
	// var data map[string]string
	// data:=QuestionResponse{}
	data:=PreferenceRequest{}
	if err := json.Unmarshal(requestBody, &data); err != nil {
        panic(err)
    }

    preferenceStatus:=UpdatePreferecePResponseDB(userId,data)
	auth.UpdatePreferencetatus(userId,preferenceStatus)
    
	// preferenceStatus:=UpdatePreferecePResponseDB(userId,data)
	// auth.UpdatePreferencetatus(userId,preferenceStatus)
	// // UpdateIPIPScore(userId)
	if(preferenceStatus>=6){
		util.InsertSessionValue(db,sessionId,userId,ipAddress)
        util.DeleteTemporarySession(db,sessionId)
	}
	util.Message(w,200)
	db.Close()
}

func getDataInInterestForm(interest Interest,data PreferenceRequest) (int,Interest){
	preferenceStatus:=data.Page
	if(preferenceStatus==1){
		interest.InterestIndoorPassive1=data.I1
		interest.InterestIndoorPassive2=data.I2
	} else if(preferenceStatus==2){
		interest.InterestIndoorActive1=data.I1
		interest.InterestIndoorActive2=data.I2
	} else if(preferenceStatus==3){
		interest.InterestOutdoorPassive1=data.I1
		interest.InterestOutdoorPassive2=data.I2
	} else if(preferenceStatus==4){
		interest.InterestOutdoorActive1=data.I1
		interest.InterestOutdoorActive2=data.I2
	} else if(preferenceStatus==5){
		interest.InterestOthers1=data.I1
		interest.InterestOthers2=data.I2
	} else if(preferenceStatus==6){
		interest.InterestIdeology1=data.I1
		interest.InterestIdeology2=data.I2
	}
	return preferenceStatus,interest
}
// func getDataInInterestForm(interest Interest,data map[string]string) (int,Interest){
// 	preferenceStatus:=0
// 	for key,value:=range data{
// 		switch key{
// 		case "InterestIndoorPassive1":
// 			preferenceStatus=1
// 			interest.InterestIndoorPassive1=value
// 		case "InterestIndoorPassive2":
// 			preferenceStatus=1
// 			interest.InterestIndoorPassive2=value
// 		case "InterestIndoorActive1":
// 			preferenceStatus=2
// 			interest.InterestIndoorActive1=value
// 		case "InterestIndoorActive2":
// 			preferenceStatus=2
// 			interest.InterestIndoorActive2=value
// 		case "InterestOutdoorPassive1":
// 			preferenceStatus=3
// 			interest.InterestOutdoorPassive1=value
// 		case "InterestOutdoorPassive2":
// 			preferenceStatus=3
// 			interest.InterestOutdoorPassive2=value
// 		case "InterestOutdoorActive1":
// 			preferenceStatus=4
// 			interest.InterestOutdoorActive1=value
// 		case "InterestOutdoorActive2":
// 			preferenceStatus=4
// 			interest.InterestOutdoorActive2=value
// 		case "InterestOthers1":
// 			preferenceStatus=5
// 			interest.InterestOthers1=value
// 		case "InterestOthers2":
// 			preferenceStatus=5
// 			interest.InterestOthers2=value
// 		case "InterestIdeology1":
// 			preferenceStatus=6
// 			interest.InterestIdeology1=value
// 		case "InterestIdeology2":
// 			preferenceStatus=6
// 			interest.InterestIdeology2=value
// 		}
// 	}

// 	return preferenceStatus,interest
// }