package Profile

import(
	"net/http"
	"io/ioutil"
	"encoding/json"
   util "miti-microservices/Util"
   auth "miti-microservices/Authentication"
   database "miti-microservices/Database"
   "bytes"
   "log"
)


func ProfileCreation(w http.ResponseWriter, r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	header:=ProfileCreationHeader{}

	content:=ProfileCreationResponse{}
	statusCode:=0

	responseHeader:=ProfileCreationResponseHeader{}
	var data map[string]string

	db:=database.DBConnection()
	list:=[]bool{false,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	util.GetHeader(r,&header)
	sessionId:=header.Cookie

	userId,dErr,dbError:=util.GetUserIdFromTemporarySession3(db,sessionId)
	errorList.DatabaseError=dbError
	util.APIHitLog("ProfileCreation",ipAddress,sessionId)

	if dErr=="Error"{
		errorList.TemporarySessionError=true
	}


	requestBody,err:=ioutil.ReadAll(r.Body)
	if (err!=nil && !util.ErrorListStatus(errorList)){
		errorList.BodyReadError=true
	}

	profileData:=Profile{}
	if(!util.ErrorListStatus(errorList)){
		errProfileData:=json.Unmarshal(requestBody,&profileData)
		if errProfileData!=nil{
			errorList.UnmarshallingError=true
		}
	}

	profileData.UserId=userId
	if(!util.ErrorListStatus(errorList)){
		sanatizationStatus:=Sanatize(profileData)
		if sanatizationStatus =="Error"{
			errorList.SanatizationError=true
		}
	}
	
	if(!util.ErrorListStatus(errorList)){
		dbError:=auth.UpdateProfileCreationStatus(db,userId)
		errorList.DatabaseError=dbError	
	}

	if(!util.ErrorListStatus(errorList)){
		dbError:=EnterProfileData(db,profileData)
		errorList.DatabaseError=dbError
	}

	if(!util.ErrorListStatus(errorList)){
		statusCode=200
	}
	
	
	code:=util.GetCode(errorList)
	if(code==200){
		content.Code=statusCode
	}else{
		content.Code=code
	}
	content.Message=util.GetMessageDecode(content.Code)
	responseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(responseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
	p:=&content
	util.ResponseLog("ProfileCreation",ipAddress,sessionId,content.Code,*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()

}