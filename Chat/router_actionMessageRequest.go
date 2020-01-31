package Chat

import(
	"net/http"
	"log"
	"io/ioutil"
	"strings"
	"encoding/json"
   util "miti-microservices/Util"
   database "miti-microservices/Database"
   profile "miti-microservices/Profile"
   "bytes"
)

func ActionMessageRequest(w http.ResponseWriter,r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	header:=ActionMessageRequestHeader{}
	util.GetHeader(r,&header)

	content:=ActionMessageRequestResponse{}
	statusCode:=0

	actionMessageRequestResponseHeader:=ActionMessageRequestResponseHeader{}
	var data map[string]string

	db:=database.DBConnection()
	//Session,TemporarySession,Body,Unmarshal,Sanatize,Database
	list:=[]bool{true,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	sessionId:=header.Cookie
	userId,dErr,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	util.APIHitLog("ActionMessageRequest",ipAddress,sessionId)
	if dErr=="Error"{
		errorList.SessionError=true
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if (err!=nil && !util.ErrorListStatus(errorList)){
		errorList.BodyReadError=true
	}
	
	actionMessageRequestData:=ActionMessageRequestDS{}
	if(!util.ErrorListStatus(errorList)){
		profileRequestErr:=json.Unmarshal(requestBody,&actionMessageRequestData)
		if profileRequestErr!=nil{
			errorList.UnmarshallingError=true
		}
	}

	if(!util.ErrorListStatus(errorList)){
		util.BodyLog("ActionMessageRequest",ipAddress,sessionId,actionMessageRequestData)
		sanatizationStatus :=Sanatize(actionMessageRequestData)
		if(sanatizationStatus=="Error"){
			errorList.SanatizationError=true
		}
	}

	actionRequestId:=actionMessageRequestData.RequestId
	action:=actionMessageRequestData.Action
	action=strings.ToLower(action)
	senderPhone:=actionMessageRequestData.Phone

	var phone string
	if(!util.ErrorListStatus(errorList)){
		phone,dbError=GetUserPhone(db,userId)
		errorList.DatabaseError=dbError
	}
	
	updatedAt:=util.GetTime()
	if(action=="accept" && !util.ErrorListStatus(errorList)){
		userId2,updatedAtTemp,messageRequest,dbError:=UpdateMessageRequestDB(db,phone,senderPhone,action,actionRequestId,updatedAt)
		errorList.DatabaseError=dbError
		updatedAt=updatedAtTemp
		var chatId string
		if(!util.ErrorListStatus(errorList)){
			chatId,dbError=InsertChatDetail(db,userId,userId2,actionRequestId)
			errorList.DatabaseError=dbError
		}
		if(!util.ErrorListStatus(errorList)){
			dbError=InsertIntoChatFromMessageRequest(db,chatId,actionRequestId,messageRequest)	
			errorList.DatabaseError=dbError
		}

		if(!util.ErrorListStatus(errorList)){
			dbError=profile.InsertIntoMatch(db,userId,userId2)
			errorList.DatabaseError=dbError
		}
		
	}else if(action=="reject" && !util.ErrorListStatus(errorList)){
		_,updatedAt,_,dbError=UpdateMessageRequestDB(db,phone,senderPhone,action,actionRequestId,updatedAt)
		errorList.DatabaseError=dbError
	}else{
		statusCode=1002
	}


	code:=util.GetCode(errorList)
	if(code==200){
		content.Code=statusCode
	}else{
		content.Code=code
	}
	content.Message=util.GetMessageDecode(code)
	content.RequestId=actionRequestId
	content.CreatedAt=updatedAt

	actionMessageRequestResponseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(actionMessageRequestResponseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
	p:=&content
	util.ResponseLog("ActionMessageRequest",ipAddress,sessionId,content.Code,*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}