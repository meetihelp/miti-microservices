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
func GetChatDetailroute(w http.ResponseWriter, r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	getChatDetailHeader:=ChatDetailHeader{}
	var data map[string]string
	content:=ChatDetailResponse{}
	statusCode:=0
	chatDetailResponseHeader:=GetChatDetailResponseHeader{}
	db:=database.DBConnection()
	//Session,TemporarySession,Body,Unmarshal,Sanatize,Database
	list:=[]bool{false,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	util.GetHeader(r,&getChatDetailHeader)
	sessionId:=getChatDetailHeader.Cookie
	util.APIHitLog("GetChatDetail",ipAddress,sessionId)

	userId,getChatStatus,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	if (getChatStatus=="Error"){
		fmt.Println("GetChatDetail line 30")
		errorList.SessionError=true
	}


	//Read body data
	requestBody,err:=ioutil.ReadAll(r.Body)
	if (err!=nil && !util.ErrorListStatus(errorList)){
		fmt.Println("GetChatDetail line 38")
		errorList.BodyReadError=true
	}

	chatDetailData :=ChatDetailRequest{}
	if (!util.ErrorListStatus(errorList)){
		fmt.Println("GetChatDetail line 44")
		errUserData:=json.Unmarshal(requestBody,&chatDetailData)
		if(errUserData!=nil){
			errorList.UnmarshallingError=true	
		}		
	}

	if(!util.ErrorListStatus(errorList)){
		fmt.Println("GetChatDetail line 52")
		util.BodyLog("GetChatDetail",ipAddress,sessionId,chatDetailData)	
		sanatizationStatus :=Sanatize(chatDetailData)
		if(sanatizationStatus=="Error"){
			errorList.SanatizationError=true
		}
	}
	
	var userId2 []string
	var chatDetail []ChatDetail
	var chatDetailErr string
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("GetChatDetail line 64")
		numOfChat:=chatDetailData.NumOfChat
		createdAt:=chatDetailData.CreatedAt
		userId2,chatDetail,chatDetailErr,dbError=GetChatDetail(db,userId,createdAt,numOfChat)	
		errorList.DatabaseError=dbError
		if(!util.ErrorListStatus(errorList) || chatDetailErr=="Error"){
			fmt.Println("GetChatDetail line 70")
			statusCode=7000
		}else{
			fmt.Println("GetChatDetail line 73")
			statusCode=200
		}
	}

	if(!util.ErrorListStatus(errorList)){
		statusCode=200
	}

	code:=util.GetCode(errorList)
	if(code==200){
		fmt.Println("GetChatDetail line 80")
		content.Code=statusCode
	}else{
		fmt.Println("GetChatDetail line 83")
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