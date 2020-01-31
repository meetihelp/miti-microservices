package Chat

import(
	"fmt"
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
	list:=[]bool{false,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	sessionId:=header.Cookie
	userId,dErr,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	util.APIHitLog("SendMessageRequest",ipAddress,sessionId)
	if dErr=="Error"{
		fmt.Println("SendMessageRequest line 37")
		errorList.SessionError=true
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if (err!=nil && !util.ErrorListStatus(errorList)){
		fmt.Println("SendMessageRequest line 43")
		errorList.BodyReadError=true
	}
	
	sendMessageRequestData:=SendMessageRequestDS{}
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("SendMessageRequest line 49")
		profileRequestErr:=json.Unmarshal(requestBody,&sendMessageRequestData)
		if(profileRequestErr!=nil){
			errorList.UnmarshallingError=true
		}
	}

	if(!util.ErrorListStatus(errorList)){
		fmt.Println("SendMessageRequest line 57")
		sanatizationErr:=Sanatize(sendMessageRequestData)
		if(sanatizationErr=="Error"){
			errorList.SanatizationError=true
		}
	}
	var phone string
	var phoneErr string
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("SendMessageRequest line 66")
		phoneErr,phone=util.GetPhoneInFormat(sendMessageRequestData.Phone)
		if(phoneErr=="Error"){
			fmt.Println("SendMessageRequest line 69")
			errorList.LogicError=true
			statusCode=1002
		}
	}
	requestId:=sendMessageRequestData.RequestId
	var senderPhone string
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("SendMessageRequest line 77")
		senderPhone,dbError=GetUserPhone(db,userId)
		errorList.DatabaseError=dbError
	}
	
	messageType:=sendMessageRequestData.MessageType
	messageContent:=sendMessageRequestData.MessageContent
	var senderName string
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("SendMessageRequest line 86")
		senderName,dbError=profile.GetUserName(db,userId)
		errorList.DatabaseError=dbError
	}
	createdAt:=util.GetTime()
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("SendMessageRequest line 92")
		createdAt,dbError=InsertMessageRequestDB(db,userId,senderName,senderPhone,phone,requestId,messageType,messageContent,createdAt)
		errorList.DatabaseError=dbError	
	}
	
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("SendMessageRequest line 98")
		availability,dbError:=IsPhoneNumberExist(db,phone)
		errorList.DatabaseError=dbError
		if(availability=="No"){
			fmt.Println("SendMessageRequest line 102")
			_,err:=sms.MessageRequestNotificaton(senderName,senderPhone,phone)
			if(err!=nil){
				fmt.Println("SendMessageRequest line 105")
				statusCode=1007
			}
		}	
	}
	
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("SendMessageRequest line 112")
		statusCode=200
	}
	
	code:=util.GetCode(errorList)
	if(code==200){
		fmt.Println("SendMessageRequest line 118")
		content.Code=statusCode
	}else{
		fmt.Println("SendMessageRequest line 121")
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