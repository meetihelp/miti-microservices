package Authentication
import (
	"fmt"
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
	list:=[]bool{false,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	//Unmarshalling Header
	util.GetHeader(r,&loadingPageHeader)
	sessionId:=loadingPageHeader.Cookie
	util.APIHitLog("LoadingPage",ipAddress,sessionId)
	
	//Querying for Session Status
	userId,loginStatus,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	sessionFlag:=0
	if (loginStatus=="Ok" && !util.ErrorListStatus(errorList)){
		fmt.Println("LoadingPage line 40")
		util.SessionLog("LoadingPage",ipAddress,sessionId,"Success")
		errorList.SessionError=false
		statusCode=200
		moveTo=6
		sessionFlag=1
	}

	//Querying for temporary session status
	if(!util.ErrorListStatus(errorList) && sessionFlag==0){
		fmt.Println("LoadingPage line 50")
		util.SessionLog("LoadingPage",ipAddress,sessionId,"Fail")
		userId,loginStatus,errorList.DatabaseError=util.GetUserIdFromTemporarySession3(db,sessionId)
		if(loginStatus=="Ok"){
			fmt.Println("LoadingPage line 54")
			util.TemporarySessionLog("LoadingPage",ipAddress,sessionId,"Success")
			errorList.TemporarySessionError=false
		}else{
			fmt.Println("LoadingPage line 58")
			errorList.TemporarySessionError=true
			moveTo=2
		}
	}

	if(errorList.TemporarySessionError && sessionFlag==0){
		fmt.Println("LoadingPage line 65")
		util.TemporarySessionLog("LoadingPage",ipAddress,sessionId,"Fail")

	}


	//Checking ProfileCreation Status
	if(!util.ErrorListStatus(errorList) && sessionFlag==0){
		fmt.Println("LoadingPage line 73")
		IsUserVerified,IsProfileCreated,Preference,dbError:=LoadingPageQuery(db,userId)
		errorList.DatabaseError=dbError
		preference=Preference
		if(!IsUserVerified && !util.ErrorListStatus(errorList)){
			fmt.Println("LoadingPage line 78")
			statusCode=1004
			moveTo=3
		}else if(!IsProfileCreated && !util.ErrorListStatus(errorList)){
			fmt.Println("LoadingPage line 82")
			statusCode=1005
			moveTo=4
		}else if(Preference<NUM_OF_PREFERENCE && !util.ErrorListStatus(errorList)){
			fmt.Println("LoadingPage line 86")
			statusCode=1003
			moveTo=5
		}else{
			fmt.Println("LoadingPage line 90")
			statusCode=1004
			moveTo=3
		}
	}

	//Setting Response
	code:=util.GetCode(errorList)
	if(code==200){
		fmt.Println("LoadingPage line 99")
		content.Code=statusCode
	}else{
		fmt.Println("LoadingPage line 102")
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

