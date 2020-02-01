package Social

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	// "strings"
	"encoding/json"
   util "miti-microservices/Util"
   gps "miti-microservices/GPS"
   database "miti-microservices/Database"
   "bytes"
   // profile "miti-microservices/Profile"
)

func GroupPoolStatusRouter(w http.ResponseWriter, r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	header:=GroupPoolStatusHeader{}

	content:=GroupPoolStatusResponse{}
	statusCode:=0

	responseHeader:=GroupPoolStatusResponseHeader{}
	var data map[string]string


	db:=database.DBConnection()
	//Session,TemporarySession,Body,Unmarshal,Sanatize,Database
	list:=[]bool{false,false,false,false,false,false}
	errorList:=util.GetErrorList(list)


	util.GetHeader(r,&header)
	sessionId:=header.Cookie
	userId,dErr,dbError:=util.GetUserIdFromSession3(db,sessionId)
	util.APIHitLog("GroupPoolStatus",ipAddress,sessionId)
	if dErr=="Error"{
		fmt.Println("GroupPoolStatus Line 37")
		errorList.SessionError=true
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if (err!=nil && !util.ErrorListStatus(errorList)){
		fmt.Println("GroupPoolStatus Line 43")
		errorList.BodyReadError=true
	}

	groupPoolStatusData:=GroupPoolStatusRequest{}
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("GroupPoolStatus Line 49")
		errUserData:=json.Unmarshal(requestBody,&groupPoolStatusData)
		if(errUserData!=nil){
			errorList.UnmarshallingError=true
		}	
	}

	latitude:=groupPoolStatusData.Latitude
	longitude:=groupPoolStatusData.Longitude
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("GroupPoolStatus Line 59")
		dbError:=gps.UpdateUserCurrentLocation(db,userId,latitude,longitude)	
		errorList.DatabaseError=dbError
	}

	var interest []string
	var groupPoolStatus []GroupPoolStatusHelper
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("GroupPoolStatus Line 67")
		interest,groupPoolStatus,dbError=GroupPoolStatusDB(db,userId)
		errorList.DatabaseError=dbError
	}

	if(!util.ErrorListStatus(errorList)){
		statusCode=200
	}
	
	code:=util.GetCode(errorList)
	if(code==200){
		fmt.Println("GroupPoolStatus Line 82")
		content.Code=statusCode
	}else{
		fmt.Println("GroupPoolStatus Line 81")
		content.Code=code
	}
	content.Message=util.GetMessageDecode(code)
	content.Interest=interest
	content.Status=groupPoolStatus

	responseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(responseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
	p:=&content
	util.ResponseLog("PoolStatus",ipAddress,sessionId,content.Code,*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}