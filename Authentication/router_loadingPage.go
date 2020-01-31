package Authentication
import (
	"net/http"
	util "miti-microservices/Util"
	database "miti-microservices/Database"
	"encoding/json"
	"bytes"
	"log"
)


func LoadingPage(w http.ResponseWriter,r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	loadingPageHeader:=LoadingPageHeader{}

	//Initializing response variables
	content:=LoadingResponse{}
	statusCode:=0
	moveTo:=0
	preference:=0

	//Initializing response header variables
	loadingPageResponseHeader:=LoadingPageResponseHeader{}
	var data map[string]string
	
	//Creating DB Connection for Router
	db:=database.DBConnection()
	//Session,TemporarySession,Body,Unmarshal,Sanatize,Database
	list:=[]bool{false,true,false,false,false,false}
	errorList:=util.GetErrorList(list)

	//Unmarshalling Header
	util.GetHeader(r,&loadingPageHeader)
	sessionId:=loadingPageHeader.Cookie
	util.APIHitLog("LoadingPage",ipAddress,sessionId)
	
	//Querying for Session Status
	userId,loginStatus,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	if (loginStatus=="Ok" && !errorList.DatabaseError){
		util.SessionLog("LoadingPage",ipAddress,sessionId,"Success")
		errorList.SessionError=false
		statusCode=200
		moveTo=6
	}

	//Querying for temporary session status
	if(errorList.SessionError && !errorList.DatabaseError){
		util.SessionLog("LoadingPage",ipAddress,sessionId,"Fail")
		userId,loginStatus,errorList.DatabaseError=util.GetUserIdFromTemporarySession3(db,sessionId)
		if(loginStatus=="Ok"){
			util.TemporarySessionLog("LoadingPage",ipAddress,sessionId,"Success")
			errorList.TemporarySessionError=false
		}else{
			moveTo=2
		}
	}

	if(errorList.TemporarySessionError){
		util.TemporarySessionLog("LoadingPage",ipAddress,sessionId,"Fail")

	}


	//Checking ProfileCreation Status
	if(!errorList.TemporarySessionError && !errorList.DatabaseError){
		IsUserVerified,IsProfileCreated,Preference,dbError:=LoadingPageQuery(db,userId)
		errorList.DatabaseError=dbError
		preference=Preference
		if(!IsUserVerified && !errorList.DatabaseError){
			statusCode=1004
			moveTo=3
		}else if(!IsProfileCreated && !errorList.DatabaseError){
			statusCode=1005
			moveTo=4
		}else if(Preference<NUM_OF_PREFERENCE && !errorList.DatabaseError){
			statusCode=1003
			moveTo=5
		}else{
			statusCode=1004
			moveTo=3
		}
	}

	//Setting Response
	code:=util.GetCode(errorList)
	if(code==200){
		content.Code=statusCode
	}else{
		content.Code=code
	}
	content.Message=util.GetMessageDecode(content.Code)
	content.MoveTo=moveTo
	content.Preference=preference

	//Setting Response Header
	loadingPageResponseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(loadingPageResponseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
	p:=&content
	util.ResponseLog("LoadingPage",ipAddress,sessionId,content.Code,*p)
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}

