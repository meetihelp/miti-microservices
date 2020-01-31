package Chat

import(
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
   util "miti-microservices/Util"
   profile "miti-microservices/Profile"
   database "miti-microservices/Database"
   sms "miti-microservices/Notification/SMS"
   "bytes"
)

func SendMessageRequest(w http.ResponseWriter,r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	header:=SendMessageRequestHeader{}
	util.GetHeader(r,&header)
	
	content:=SendMessageRequestResponse{}
	statusCode:=0

	sendMessageRequestResponseHeader:=SendMessageRequestResponseHeader{}
	var data map[string]string

	db:=database.DBConnection()
	//Session,TemporarySession,Body,Unmarshal,Sanatize,Database
	list:=[]bool{true,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	sessionId:=header.Cookie
	userId,dErr,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	util.APIHitLog("SendMessageRequest",ipAddress,sessionId)
	if dErr=="Error"{
		errorList.SessionError=true
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if (err!=nil && !util.ErrorListStatus(errorList)){
		errorList.BodyReadError=true
	}
	
	sendMessageRequestData:=SendMessageRequestDS{}
	if(!util.ErrorListStatus(errorList)){
		profileRequestErr:=json.Unmarshal(requestBody,&sendMessageRequestData)
		if(profileRequestErr!=nil){
			errorList.UnmarshallingError=true
		}
	}

	if(!util.ErrorListStatus(errorList)){
		sanatizationErr:=Sanatize(sendMessageRequestData)
		if(sanatizationErr=="Error"){
			errorList.SanatizationError=true
		}
	}
	var phone string
	var phoneErr string
	if(!util.ErrorListStatus(errorList)){
		phoneErr,phone=util.GetPhoneInFormat(sendMessageRequestData.Phone)
		if(phoneErr=="Error"){
			errorList.LogicError=true
			statusCode=1002
		}
	}
	requestId:=sendMessageRequestData.RequestId
	var senderPhone string
	if(!util.ErrorListStatus(errorList)){
		senderPhone,dbError=GetUserPhone(db,userId)
		errorList.DatabaseError=dbError
	}
	
	messageType:=sendMessageRequestData.MessageType
	messageContent:=sendMessageRequestData.MessageContent
	var senderName string
	if(!util.ErrorListStatus(errorList)){
		senderName,dbError=profile.GetUserName(db,userId)
		errorList.DatabaseError=dbError
	}
	createdAt:=util.GetTime()
	if(!util.ErrorListStatus(errorList)){
		createdAt,dbError=InsertMessageRequestDB(db,userId,senderName,senderPhone,phone,requestId,messageType,messageContent,createdAt)
		errorList.DatabaseError=dbError	
	}
	
	if(!util.ErrorListStatus(errorList)){
		availability,dbError:=IsPhoneNumberExist(db,phone)
		errorList.DatabaseError=dbError
		if(availability=="No"){
			sms.MessageRequestNotificaton(senderName,senderPhone,phone)
		}	
	}
	
	
	code:=util.GetCode(errorList)
	if(code==200){
		content.Code=statusCode
	}else{
		content.Code=code
	}
	content.Message=util.GetMessageDecode(code)
	content.RequestId=requestId
	content.CreatedAt=createdAt

	sendMessageRequestResponseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(sendMessageRequestResponseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
	p:=&content
	util.ResponseLog("SendMessageRequest",ipAddress,sessionId,content.Code,*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()

}