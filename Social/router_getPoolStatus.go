package Social

import(
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
   util "miti-microservices/Util"
   gps "miti-microservices/GPS"
   profile "miti-microservices/Profile"
   database "miti-microservices/Database"
   "bytes"
)

func PoolStatusRouter(w http.ResponseWriter, r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	header:=PoolStatusHeader{}

	content:=PoolStatusResponse{}
	statusCode:=0

	responseHeader:=PoolStatusResponseHeader{}
	var data map[string]string

	db:=database.DBConnection()
	//Session,TemporarySession,Body,Unmarshal,Sanatize,Database
	list:=[]bool{false,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	util.GetHeader(r,&header)
	sessionId:=header.Cookie
	userId,dErr,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	util.APIHitLog("PoolStatus",ipAddress,sessionId)
	if dErr=="Error"{
		errorList.SessionError=true
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if (err!=nil && !util.ErrorListStatus(errorList)){
		errorList.BodyReadError=true
	}

	poolStatusData:=PoolStatusRequest{}
	if(!util.ErrorListStatus(errorList)){
		errUserData:=json.Unmarshal(requestBody,&poolStatusData)
		if(errUserData!=nil){
			errorList.UnmarshallingError=true
		}	
	}

	latitude:=poolStatusData.Latitude
	longitude:=poolStatusData.Longitude
	if(!util.ErrorListStatus(errorList)){
		dbError:=gps.UpdateUserCurrentLocation(db,userId,latitude,longitude)	
		errorList.DatabaseError=dbError
	}
	
	var poolStatus PoolStatus
	if(!util.ErrorListStatus(errorList)){
		poolStatus,dbError=PoolStatusDB(db,userId)
		errorList.DatabaseError=dbError
	}

	var ipip int
	if(!util.ErrorListStatus(errorList)){
		ipip,dbError=profile.CheckIPIPStatus(db,userId)
		errorList.DatabaseError=dbError
	}

	code:=util.GetCode(errorList)
	if(code==200){
		content.Code=statusCode
		if(ipip<5){
			content.Code=2003
		}
	}else{
		content.Code=code
	}
	content.Message=util.GetMessageDecode(code)
	content.ChatId=poolStatus.ChatId
	content.MatchUserId=poolStatus.MatchUserId
	content.Status=poolStatus.Status
	content.CreatedAt=poolStatus.CreatedAt
	content.MatchTime=poolStatus.MatchTime
	content.IPIP=ipip

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