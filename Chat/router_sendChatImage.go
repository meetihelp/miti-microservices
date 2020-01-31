package Chat

import(
	"fmt"
	"net/http"
	util "miti-microservices/Util"
	image "miti-microservices/Media/Image"
	database "miti-microservices/Database"
	"io/ioutil"
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
	list:=[]bool{false,false,false,false,false,false}
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
		fmt.Println("SendChatImage line 51")
		errorList.SessionError=true
	}

	file, _, err := r.FormFile("myFile")
	errorStatus:=util.ErrorListStatus(errorList)
    if(err!=nil && !errorStatus){
    	fmt.Println("SendChatImage line 58")
        errorList.BodyReadError=true
    }

	buffer, err := ioutil.ReadAll(file)
	errorStatus=util.ErrorListStatus(errorList)
    if(err!=nil && !errorStatus){
    	fmt.Println("SendChatImage line 65")
        errorList.BodyReadError=true
    }

    errorStatus=util.ErrorListStatus(errorList)
    if(!errorStatus){
    	fmt.Println("SendChatImage line 71")
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
		fmt.Println("SendChatImage line 86")
		userImageData,status,dbError=image.GetUserImageByRequestId(db,userId,requestId)	
		errorList.DatabaseError=dbError
	}

	errorStatus=util.ErrorListStatus(errorList)
	if(status=="Error" && !errorStatus){
		fmt.Println("SendChatImage line 93")
		imageId:=util.GenerateToken()
		generatedName:=util.GenerateToken()
		filename:=generatedName+"."+format
		bucket:=""
		fmt.Println("AccessType:"+accessType)
		if(accessType=="private"){
			fmt.Println("SendChatImage line 100")
			bucket=image.GetPrivateImageBucket()
			fmt.Println("Bucket:"+bucket)
		}else if(accessType=="public"){
			fmt.Println("SendChatImage line 104")
			bucket=image.GetPublicImageBucket()
		}else{
			fmt.Println("SendChatImage line 107")
			statusCode=1002
			errorList.LogicError=true
		}

		errorStatus=util.ErrorListStatus(errorList)
		if(!errorStatus){
			fmt.Println("SendChatImage line 114")
			size,err:=image.UploadToS3(buffer,filename,bucket,format)	
			if(err!=nil){
				fmt.Println("SendChatImage line 117")
				statusCode=3001
			}else{
				fmt.Println("SendChatImage line 120")
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
			fmt.Println("SendChatImage line 141")
			PublicCloudFront:=os.Getenv("publicImageCloudFront")
			url=PublicCloudFront+"/"+filename
		}

	}else{
		fmt.Println("SendChatImage line 147")
		if(accessType=="public"){
			fmt.Println("SendChatImage line 149")
			PublicCloudFront:=os.Getenv("publicImageCloudFront")
			filename:=userImageData.GeneratedName+"."+userImageData.Format
			url=PublicCloudFront+"/"+filename
		}
	}


	unSyncedChat:=[]Chat{}
	if(imageUploadStatus=="Yes"){
		fmt.Println("SendChatImage line 159")
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
			fmt.Println("SendChatImage line 173")
			dbError:=UpdateChatTime(db,chat.ChatId,chat.CreatedAt)
			errorList.DatabaseError=dbError
		}
		statusCode=200
	}else{
		fmt.Println("SendChatImage line 179")
		statusCode=3001
	}
	
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("SendChatImage line 184")
		statusCode=200
	}

	
	code:=util.GetCode(errorList)
	if(code==200){
		fmt.Println("SendChatImage line 191")
		content.Code=statusCode
	}else{
		fmt.Println("SendChatImage line 194")
		content.Code=code
	}
	content.Message=util.GetMessageDecode(content.Code)
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