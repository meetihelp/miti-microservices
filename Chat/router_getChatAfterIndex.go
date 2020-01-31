package Chat 
import(
	"net/http"
	util "miti-microservices/Util"
	database "miti-microservices/Database"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"log"
)

func GetChatAfterIndex(w http.ResponseWriter,r *http.Request){
	ipAddress:=util.GetIPAddress(r)

	getChatAfterIndexHeader:=GetChatAfterIndexHeader{}

	content:=GetChatResponse{}
	statusCode:=0

	responseHeader:=GetChatResponseHeader{}
	var data map[string]string

	db:=database.DBConnection()
	//Session,TemporarySession,Body,Unmarshal,Sanatize,Database
	list:=[]bool{true,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	util.GetHeader(r,&getChatAfterIndexHeader)
	sessionId:=getChatAfterIndexHeader.Cookie

	userId,getChatStatus,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	util.APIHitLog("GetChatAfterIndex",ipAddress,sessionId)
	if (getChatStatus=="Error"){
		errorList.SessionError=true
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if (err!=nil && !errorList.SessionError){
		errorList.BodyReadError=true
	}

	chatData:=GetChatRequest{}
	if(!errorList.BodyReadError){
		errUserData:=json.Unmarshal(requestBody,&chatData)
		if errUserData!=nil{
			errorList.UnmarshallingError=true		
		}
	}

	if(!errorList.UnmarshallingError){
		util.BodyLog("GetChatAfterIndex",ipAddress,sessionId,chatData)	
		sanatizationStatus :=Sanatize(chatData)
		if(sanatizationStatus=="Error"){
			errorList.SanatizationError=true
		}
	}
	
	var chat []Chat
	if(!errorList.SanatizationError){
		status,dbError:=CheckCorrectChat(db,userId,chatData.ChatId)
		errorList.DatabaseError=dbError
		if(status=="Error"){
			statusCode=1002
		}else if(status=="Ok" && !errorList.DatabaseError){
			chat,dbError=GetChatAfterTimeMessages(db,chatData.ChatId,chatData.NumOfChat,chatData.CreatedAt)
			errorList.DatabaseError=dbError
		}
	}

	code:=util.GetCode(errorList)
	if(code==200){
		content.Code=statusCode
	}else{
		content.Code=code
	}
	content.Message=util.GetMessageDecode(code)
	content.Chat=chat
	responseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(responseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
	p:=&content
	util.ResponseLog("GetChatAfterIndex",ipAddress,sessionId,content.Code,*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}