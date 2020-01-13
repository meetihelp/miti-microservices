package Chat

import(
	"net/http"
	util "miti-microservices/Util"
	image "miti-microservices/Media/Image"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"os"
	"log"

)

func SendChatImage(w http.ResponseWriter, r *http.Request){
	//get buffer of image from user
	sendChatImageHeader:=SendChatImageHeader{}
	util.GetHeader(r,&sendChatImageHeader)
	sessionId:=sendChatImageHeader.Cookie
	log.Println("upload Image Cookie:"+sessionId)
	accessType:=sendChatImageHeader.AccessType
	log.Println("upload Image AccessType:"+accessType)
	actualFileName:=sendChatImageHeader.ActualFileName
	format:=sendChatImageHeader.Format
	latitude:=sendChatImageHeader.Latitude
	longitude:=sendChatImageHeader.Longitude
	dimension:=sendChatImageHeader.Dimension
	requestId:=sendChatImageHeader.RequestId
	chatId:=sendChatImageHeader.ChatId
	fmt.Println(sendChatImageHeader)
	userId,getChatStatus:=util.GetUserIdFromSession(sessionId)
	// fmt.Println(userId)
	if getChatStatus=="Error"{
		util.Message(w,1003)
		return
	}

	url:=""
	chatResponse:=Chat{}
	imageUploadStatus:=""
	userImageData,status:=image.GetUserImageByRequestId(userId,requestId)
	if(status=="Error"){
		file, _, err := r.FormFile("myFile")
	    if err != nil {
	        fmt.Println("Error Retrieving the File")
	        fmt.Println(err)
	        return
	    }

		buffer, err := ioutil.ReadAll(file)
	    if err != nil {
	        fmt.Println(err)
	    }


		imageId:=util.GenerateToken()
		generatedName:=util.GenerateToken()
		filename:=generatedName+"."+format
		bucket:=""
		fmt.Println("AccessType:"+accessType)
		if(accessType=="Private"){
			bucket=image.GetPrivateImageBucket()
			fmt.Println("Bucket:"+bucket)
		}else{
			bucket=image.GetPublicImageBucket()
		}
		size,err:=image.UploadToS3(buffer,filename,bucket,format)
		if(err!=nil){
			//Could Not Upload Image
			fmt.Println(err)
			util.Message(w,3001)
		}else{
			//Uploaded image Successfully
			// userImageData=image.UserImage{}
			userImageData.UserId=userId
			userImageData.ImageId=imageId
			// fmt.Println("Send Chat ImageID:"+userImageData.ImageId)
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
			image.EnterUserImage(userImageData)
			imageUploadStatus="Yes"
		}
	

		// signedURL:=""
		// url:=""
		if(accessType=="Public"){
			PublicCloudFront:=os.Getenv("publicImageCloudFront")
			url=PublicCloudFront+"/"+filename
		}

	}else{
		// chatResponse=GetChatByRequestId(userId,requestId)
		if(accessType=="Public"){
			PublicCloudFront:=os.Getenv("publicImageCloudFront")
			filename:=userImageData.GeneratedName+"."+userImageData.Format
			url=PublicCloudFront+"/"+filename
		}
	}

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
		chatResponse=ChatInsertDB(chat)
		// db.Create(&chatData)
		if(chat.CreatedAt==chatResponse.CreatedAt){
			e:=UpdateChatTime(chat.ChatId,chat.CreatedAt)
			if e!=nil{
				return
			}
		}
	}
	


	code:=200
	msg:=util.GetMessageDecode(code)
	w.Header().Set("Content-Type", "application/json")
	// p:=&UploadImageResponse{Code:code,Message:msg,ImageId:imageId,URL:url}
	p:=&SendChatImageResponse{Code:code,Message:msg,ImageId:userImageData.ImageId,
				RequestId:requestId,MessageId:chatResponse.MessageId,
				CreatedAt:chatResponse.CreatedAt,MessageType:"image",URL:url}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}