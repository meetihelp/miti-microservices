package Profile

import(
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
   util "miti-microservices/Util"
   auth "miti-microservices/Authentication"
   database "miti-microservices/Database"
   "bytes"
)

func UpdatePreference(w http.ResponseWriter, r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	header:=UpdatePreferenceHeader{}

	content:=UpdatePreferenceResponse{}
	statusCode:=0

	responseHeader:=UpdatePreferenceResponseHeader{}
	var data map[string]string

	db:=database.DBConnection()
	list:=[]bool{false,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	util.GetHeader(r,&header)
	sessionId:=header.Cookie

	userId,dErr,dbError:=util.GetUserIdFromTemporarySession3(db,sessionId)
	errorList.DatabaseError=dbError
	util.APIHitLog("UpdatePreference",ipAddress,sessionId)
	if dErr=="Error"{
		errorList.TemporarySessionError=true
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if (err!=nil && !util.ErrorListStatus(errorList)){
		errorList.BodyReadError=true
	}

	preferenceRequestData:=UpdatePreferenceRequest{}
	if(!util.ErrorListStatus(errorList)){
		err:=json.Unmarshal(requestBody, &preferenceRequestData)
		if(err!=nil){
			errorList.UnmarshallingError=true
		}
	}

	if(!util.ErrorListStatus(errorList)){
		sanatizationStatus:=Sanatize(preferenceRequestData)
		if sanatizationStatus =="Error"{
			errorList.SanatizationError=true
		}
	}

	var preferenceStatus int
	if(!util.ErrorListStatus(errorList)){
		preferenceStatus,dbError=UpdatePreferecePResponseDB(db,userId,preferenceRequestData)		
		errorList.DatabaseError=dbError
	}

	if(!util.ErrorListStatus(errorList)){
		dbError=auth.UpdatePreferencetatus(db,userId,preferenceStatus)	
		errorList.DatabaseError=dbError
	}

	if(preferenceStatus>=6){
		if(!util.ErrorListStatus(errorList)){
			dbError=util.InsertSessionValue(db,sessionId,userId,ipAddress)
			errorList.DatabaseError=dbError
		}
		if(!util.ErrorListStatus(errorList)){
			dbError=util.DeleteTemporarySession(db,sessionId)
			errorList.DatabaseError=dbError
		}
		
        
	}
	
	code:=util.GetCode(errorList)
	if(code==200){
		content.Code=statusCode
	}else{
		content.Code=code
	}
	responseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(responseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
	p:=&content
	util.ResponseLog("UpdateIPIP",ipAddress,sessionId,content.Code,*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}

func getDataInInterestForm(interest Interest,data UpdatePreferenceRequest) (int,Interest){
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