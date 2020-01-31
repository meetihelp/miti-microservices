package Chat

import(
	"net/http"
	util "miti-microservices/Util"
	image "miti-microservices/Media/Image"
	database "miti-microservices/Database"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"os"
	"log"
	"bytes"
	"strings"

)

func SendChatImage(w http.ResponseWriter, r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	sendChatImageHeader:=SendChatImageHeader{}

	content:=SendChatImageResponse{}
	statusCode:=0

	sendChatImageResponseHeader:=SendChatImageResponseHeader{}
	var data map[string]string

	db:=database.DBConnection()
	//Session,TemporarySession,Body,Unmarshal,Sanatize,Database
	list:=[]bool{true,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	util.GetHeader(r,&sendChatImageHeader)
	sessionId:=sendChatImageHeader.Cookie
	accessType:=sendChatImageHeader.AccessType
	accessType=strings.ToLower(accessType)
	actualFileName:=sendChatImageHeader.ActualFileName
	format:=sendChatImageHeader.Format
	latitude:=sendChatImageHeader.Latitude
	longitude:=sendChatImageHeader.Longitude
	dimension:=sendChatImageHeader.Dimension
	requestId:=sendChatImageHeader.RequestId
	chatId:=sendChatImageHeader.ChatId
	lastUpdate:=sendChatImageHeader.CreatedAt

	userId,getChatStatus,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	util.APIHitLog("SendChatImage",ipAddress,sessionId)
	if (getChatStatus=="Error"){
		errorList.SessionError=true
	}

	file, _, err := r.FormFile("myFile")
	errorStatus:=util.ErrorListStatus(errorList)
    if(err!=nil && !errorStatus){
        errorList.BodyReadError=true
    }

	buffer, err := ioutil.ReadAll(file)
	errorStatus=util.ErrorListStatus(errorList)
    if(err!=nil && !errorStatus){
        errorList.BodyReadError=true
    }

    errorStatus=util.ErrorListStatus(errorList)
    if(!errorStatus){
    	sanatize:=Sanatize(sendChatImageHeader)
		if(sanatize=="Error"){
			errorList.SanatizationError=true
		}
    }
    

	url:=""
	chatResponse:=Chat{}
	imageUploadStatus:="Yes"
	errorStatus=util.ErrorListStatus(errorList)
	status:="Error"
	userImageData:=image.UserImage{}
	if(!errorStatus){
		userImageData,status,dbError=image.GetUserImageByRequestId(db,userId,requestId)	
		errorList.DatabaseError=dbError
	}

	errorStatus=util.ErrorListStatus(errorList)
	if(status=="Error" && !errorStatus){
		imageId:=util.GenerateToken()
		generatedName:=util.GenerateToken()
		filename:=generatedName+"."+format
		bucket:=""
		fmt.Println("AccessType:"+accessType)
		if(accessType=="private"){
			bucket=image.GetPrivateImageBucket()
			fmt.Println("Bucket:"+bucket)
		}else if(accessType=="public"){
			bucket=image.GetPublicImageBucket()
		}else{
			statusCode=1002
			errorList.LogicError=true
		}

		errorStatus=util.ErrorListStatus(errorList)
		if(!errorStatus){
			size,err:=image.UploadToS3(buffer,filename,bucket,format)	
			if(err!=nil){
				statusCode=3001
			}else{
				userImageData.UserId=userId
				userImageData.ImageId=imageId
				userImageData.AccessType=accessType
				userImageData.ActualFileName=actualFileName
				userImageData.Size=size
				userImageData.Format=format
				userImageData.Bucket=bucket
				userImageData.Dimension=dimension
				userImageData.Latitude=latitude
				userImageData.Longitude=longitude
				userImageData.GeneratedName=generatedName
				userImageData.RequestId=requestId
				userImageData.CreatedAt=util.GetTime()
				dbError:=image.EnterUserImage(db,userImageData)
				errorList.DatabaseError=dbError
				imageUploadStatus="Yes"	
			}
		}

		if(accessType=="public"){
			PublicCloudFront:=os.Getenv("publicImageCloudFront")
			url=PublicCloudFront+"/"+filename
		}

	}else{
		if(accessType=="public"){
			PublicCloudFront:=os.Getenv("publicImageCloudFront")
			filename:=userImageData.GeneratedName+"."+userImageData.Format
			url=PublicCloudFront+"/"+filename
		}
	}


	unSyncedChat:=[]Chat{}
	if(imageUploadStatus=="Yes"){
		chat:=Chat{}
		chat.UserId=userId
		chat.ChatId=chatId
		chat.MessageType="image"
		chat.MessageContent=userImageData.ImageId
		chat.RequestId=requestId
		messageId:=util.GenerateToken()
		chat.MessageId=messageId
		createdAt:=util.GetTime()
		chat.CreatedAt=createdAt
		chatResponse,unSyncedChat,dbError=ChatInsertDB(db,chat,lastUpdate)
		errorList.DatabaseError=dbError
		if(chat.CreatedAt==chatResponse.CreatedAt){
			dbError:=UpdateChatTime(db,chat.ChatId,chat.CreatedAt)
			errorList.DatabaseError=dbError
		}
		statusCode=200
	}else{
		statusCode=3001
	}
	


	
	code:=util.GetCode(errorList)
	if(code==200){
		content.Code=statusCode
	}else{
		content.Code=code
	}
	content.Message=util.GetMessageDecode(code)
	content.ImageId=userImageData.ImageId
	content.RequestId=requestId
	content.MessageId=chatResponse.MessageId
	content.CreatedAt=chatResponse.CreatedAt
	content.MessageType="image"
	content.URL=url
	content.Chat=unSyncedChat

	sendChatImageResponseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(sendChatImageResponseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
	p:=&content
	util.ResponseLog("SendChatImage",ipAddress,sessionId,content.Code,*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}