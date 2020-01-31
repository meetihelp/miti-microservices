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
func GetChatDetailroute(w http.ResponseWriter, r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	getChatDetailHeader:=ChatDetailHeader{}
	var data map[string]string
	content:=ChatDetailResponse{}
	statusCode:=0
	chatDetailResponseHeader:=GetChatDetailResponseHeader{}
	db:=database.DBConnection()
	//Session,TemporarySession,Body,Unmarshal,Sanatize,Database
	list:=[]bool{true,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	util.GetHeader(r,&getChatDetailHeader)
	sessionId:=getChatDetailHeader.Cookie
	util.APIHitLog("GetChatDetail",ipAddress,sessionId)

	userId,getChatStatus,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	if (getChatStatus=="Error" && errorList.DatabaseError){
		errorList.SessionError=true
	}


	//Read body data
	requestBody,err:=ioutil.ReadAll(r.Body)
	if (err!=nil && !errorList.SessionError && !errorList.DatabaseError){
		errorList.BodyReadError=true
	}

	chatDetailData :=ChatDetailRequest{}
	if (!errorList.BodyReadError && !errorList.DatabaseError){
		errUserData:=json.Unmarshal(requestBody,&chatDetailData)
		if(errUserData!=nil){
			errorList.UnmarshallingError=true	
		}		
	}

	if(!errorList.UnmarshallingError){
		util.BodyLog("GetChatDetail",ipAddress,sessionId,chatDetailData)	
		sanatizationStatus :=Sanatize(chatDetailData)
		if(sanatizationStatus=="Error"){
			errorList.SanatizationError=true
		}
	}
	
	var userId2 []string
	var chatDetail []ChatDetail
	var chatDetailErr string
	if(!errorList.SanatizationError){
		numOfChat:=chatDetailData.NumOfChat
		createdAt:=chatDetailData.CreatedAt
		userId2,chatDetail,chatDetailErr,dbError=GetChatDetail(db,userId,createdAt,numOfChat)	
		errorList.DatabaseError=dbError
		if(!errorList.DatabaseError || chatDetailErr=="Error"){
			statusCode=7000
		}else{
			statusCode=200
		}
	}

	code:=util.GetCode(errorList)
	if(code==200){
		content.Code=statusCode
	}else{
		content.Code=code
	}
	content.Message=util.GetMessageDecode(code)
	chatDetailContent:=[]ChatDetailContent{}
	for i,c:=range chatDetail{
		temp:=ChatDetailContent{}
		// temp.UserId=c.TempUserId
		temp.UserId=c.ActualUserId
		temp.UserId2=userId2[i]
		temp.ChatId=c.ChatId
		temp.ChatType=c.ChatType
		temp.CreatedAt=c.CreatedAt
		temp.LastUpdate=c.LastUpdate
		temp.Name=c.Name
		chatDetailContent=append(chatDetailContent,temp)
	}
	content.ChatDetailContent=chatDetailContent
	chatDetailResponseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(chatDetailResponseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
	p:=&content
	util.ResponseLog("GetChatDetail",ipAddress,sessionId,content.Code,*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}