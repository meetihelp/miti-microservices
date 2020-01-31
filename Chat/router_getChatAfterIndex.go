package Chat 
import(
	"fmt"
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
	list:=[]bool{false,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	util.GetHeader(r,&getChatAfterIndexHeader)
	sessionId:=getChatAfterIndexHeader.Cookie

	userId,getChatStatus,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	util.APIHitLog("GetChatAfterIndex",ipAddress,sessionId)
	if (getChatStatus=="Error"){
		fmt.Println("GetChatAfterIndex line 36")
		errorList.SessionError=true
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if (err!=nil && !errorList.SessionError){
		fmt.Println("GetChatAfterIndex line 42")
		errorList.BodyReadError=true
	}

	chatData:=GetChatRequest{}
	if(!errorList.BodyReadError){
		fmt.Println("GetChatAfterIndex line 48")
		errUserData:=json.Unmarshal(requestBody,&chatData)
		if errUserData!=nil{
			errorList.UnmarshallingError=true		
		}
	}

	if(!errorList.UnmarshallingError){
		fmt.Println("GetChatAfterIndex line 56")
		util.BodyLog("GetChatAfterIndex",ipAddress,sessionId,chatData)	
		sanatizationStatus :=Sanatize(chatData)
		if(sanatizationStatus=="Error"){
			errorList.SanatizationError=true
		}
	}
	
	var chat []Chat
	if(!errorList.SanatizationError){
		fmt.Println("GetChatAfterIndex line 66")
		status,dbError:=CheckCorrectChat(db,userId,chatData.ChatId)
		errorList.DatabaseError=dbError
		if(status=="Error"){
			fmt.Println("GetChatAfterIndex line 70")
			statusCode=1002
		}else if(status=="Ok" && !errorList.DatabaseError){
			fmt.Println("GetChatAfterIndex line 73")
			chat,dbError=GetChatAfterTimeMessages(db,chatData.ChatId,chatData.NumOfChat,chatData.CreatedAt)
			errorList.DatabaseError=dbError
		}
	}

	if(!util.ErrorListStatus(errorList)){
		fmt.Println("GetChatAfterIndex line 80")
		statusCode=200
	}

	
	code:=util.GetCode(errorList)
	if(code==200){
		fmt.Println("GetChatAfterIndex line 87")
		content.Code=statusCode
	}else{
		fmt.Println("GetChatAfterIndex line 90")
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