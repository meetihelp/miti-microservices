package Profile

import(
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
   util "miti-microservices/Util"
   database "miti-microservices/Database"
   "bytes"
)

func GetProfile(w http.ResponseWriter, r *http.Request){
	//CHECK IF USER IS AUTHORIZED TO SEE THE PROFILE
	ipAddress:=util.GetIPAddress(r)
	header:=GetProfileHeader{}

	content:=GetProfileResponse{}
	statusCode:=0

	responseHeader:=GetProfileResponseHeader{}
	var data map[string]string

	db:=database.DBConnection()
	list:=[]bool{false,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	util.GetHeader(r,&header)
	sessionId:=header.Cookie

	userId,dErr,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	util.APIHitLog("GetProfile",ipAddress,sessionId)
	if dErr=="Error"{
		errorList.SessionError=true
	}


	

	requestBody,err:=ioutil.ReadAll(r.Body)
	if (err!=nil && !util.ErrorListStatus(errorList)){
		errorList.BodyReadError=true
	}
	
	getProfileRequest:=GetProfileRequest{}
	if(!util.ErrorListStatus(errorList)){
		err:=json.Unmarshal(requestBody, &getProfileRequest)
		if(err!=nil){
			errorList.UnmarshallingError=true
		}
	}
	if(!util.ErrorListStatus(errorList)){
		sanatizationStatus:=Sanatize(getProfileRequest)
		if sanatizationStatus =="Error"{
			errorList.SanatizationError=true
		}
	}

	var profileViewAuthorization string
	if(!util.ErrorListStatus(errorList)){
		profileViewAuthorization,dbError=ProfileViewAuthorization(db,userId,getProfileRequest.UserId)	
		errorList.DatabaseError=dbError
	}
	
	var profileResponse ProfileResponse
	if (profileViewAuthorization=="UnAuthorized" && !util.ErrorListStatus(errorList)){
		profileResponse,dbError=GetUnAuthorizedProfileDB(db,getProfileRequest.UserId)
		errorList.DatabaseError=dbError
	}else if(profileViewAuthorization=="AUTHORIZED" && !util.ErrorListStatus(errorList)){
		profileResponse,dbError=GetAuthorizedProfileDB(db,getProfileRequest.UserId)
		errorList.DatabaseError=dbError
	}else{
		statusCode=1002
	}

	//RETURN PROFILE
	code:=util.GetCode(errorList)
	if(code==200){
		content.Code=statusCode
	}else{
		content.Code=code
	}
	content.Message=util.GetMessageDecode(content.Code)
	content.ProfileResponse=profileResponse
	responseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(responseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
	p:=&content
	util.ResponseLog("GetProfile",ipAddress,sessionId,content.Code,*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
	

}