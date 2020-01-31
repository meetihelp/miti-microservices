package Chat

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
   util "miti-microservices/Util"
   database "miti-microservices/Database"
   "bytes"
)

func GetMessageRequest(w http.ResponseWriter,r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	header:=GetMessageRequestHeader{}
	util.GetHeader(r,&header)

	content:=GetMessageRequestResponse{}
	statusCode:=0

	getMessageRequestResponseHeader:=GetMessageRequestResponseHeader{}
	var data map[string]string

	db:=database.DBConnection()
	//Session,TemporarySession,Body,Unmarshal,Sanatize,Database
	list:=[]bool{false,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	sessionId:=header.Cookie
	userId,dErr,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	util.APIHitLog("GetMessageRequest",ipAddress,sessionId)
	if dErr=="Error"{
		fmt.Println("GetMessageRequest line 34")
		errorList.SessionError=true
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if(err!=nil && !util.ErrorListStatus(errorList)){
		fmt.Println("GetMessageRequest line 40")
		errorList.BodyReadError=true 
	}
	
	getMessageRequestData:=GetMessageRequestDS{}
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("GetMessageRequest line 45")
		profileRequestErr:=json.Unmarshal(requestBody,&getMessageRequestData)
		if(profileRequestErr!=nil){
			errorList.UnmarshallingError=true
		}
	}

	createdAt:=getMessageRequestData.CreatedAt
	var requestList []MessageRequestDS
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("GetMessageRequest line 56")
		requestList,dbError=GetMessageRequestDB(db,userId,createdAt)
		errorList.DatabaseError=dbError
	}

	if(!util.ErrorListStatus(errorList)){
		fmt.Println("GetMessageRequest line 63")
		statusCode=200
	}

	code:=util.GetCode(errorList)
	if(code==200){
		fmt.Println("GetMessageRequest line 63")
		content.Code=statusCode
	}else{
		fmt.Println("GetMessageRequest line 66")
		content.Code=code
	}
	content.Message=util.GetMessageDecode(code)
	content.RequestList=requestList

	getMessageRequestResponseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(getMessageRequestResponseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
	p:=&content
	util.ResponseLog("GetMessageRequest",ipAddress,sessionId,content.Code,*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}