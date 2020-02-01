package Profile

import(
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
   util "miti-microservices/Util"
   database "miti-microservices/Database"
   "bytes"
   "fmt"
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
		fmt.Println("GetProfile Line 35")
		errorList.SessionError=true
	}


	

	requestBody,err:=ioutil.ReadAll(r.Body)
	if (err!=nil && !util.ErrorListStatus(errorList)){
		fmt.Println("GetProfile Line 44")
		errorList.BodyReadError=true
	}
	
	getProfileRequest:=GetProfileRequest{}
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("GetProfile Line 50")
		err:=json.Unmarshal(requestBody, &getProfileRequest)
		if(err!=nil){
			errorList.UnmarshallingError=true
		}
	}
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("GetProfile Line 57")
		sanatizationStatus:=Sanatize(getProfileRequest)
		if sanatizationStatus =="Error"{
			errorList.SanatizationError=true
		}
	}

	var profileViewAuthorization string
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("GetProfile Line 66")
		profileViewAuthorization,dbError=ProfileViewAuthorization(db,userId,getProfileRequest.UserId)	
		errorList.DatabaseError=dbError
	}
	
	var profileResponse ProfileResponse
	if (profileViewAuthorization=="UnAuthorized" && !util.ErrorListStatus(errorList)){
		fmt.Println("GetProfile Line 73")
		profileResponse,dbError=GetUnAuthorizedProfileDB(db,getProfileRequest.UserId)
		errorList.DatabaseError=dbError
	}else if(profileViewAuthorization=="Authorized" && !util.ErrorListStatus(errorList)){
		fmt.Println("GetProfile Line 77")
		profileResponse,dbError=GetAuthorizedProfileDB(db,getProfileRequest.UserId)
		errorList.DatabaseError=dbError
	}else{
		fmt.Println("GetProfile Line 81")
		statusCode=1002
	}

	if(!util.ErrorListStatus(errorList) && statusCode==0){
		fmt.Println("GetProfile Line 86")
		statusCode=200
	}

	//RETURN PROFILE
	code:=util.GetCode(errorList)
	if(code==200){
		fmt.Println("GetProfile Line 93")
		content.Code=statusCode
	}else{
		fmt.Println("GetProfile Line 96")
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