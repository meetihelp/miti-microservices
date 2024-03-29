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
	"strings"

)

func SendChatImage(w http.ResponseWriter, r *http.Request){
	//get buffer of image from user
	sendChatImageHeader:=SendChatImageHeader{}
	util.GetHeader(r,&sendChatImageHeader)
	sessionId:=sendChatImageHeader.Cookie
	log.Println("upload Image Cookie:"+sessionId)
	accessType:=sendChatImageHeader.AccessType
	accessType=strings.ToLower(accessType)
	log.Println("upload Image AccessType:"+accessType)
	actualFileName:=sendChatImageHeader.ActualFileName
	format:=sendChatImageHeader.Format
	latitude:=sendChatImageHeader.Latitude
	longitude:=sendChatImageHeader.Longitude
	dimension:=sendChatImageHeader.Dimension
	requestId:=sendChatImageHeader.RequestId
	chatId:=sendChatImageHeader.ChatId
	lastUpdate:=sendChatImageHeader.CreatedAt
	fmt.Print("SendChatImageHeader")
	fmt.Println(sendChatImageHeader)
	db:=database.DBConnection()

	userId,getChatStatus:=util.GetUserIdFromSession2(db,sessionId)
	// fmt.Println(userId)
	if getChatStatus=="Error"{
		fmt.Println("Session Error for SendChatImage")
		util.Message(w,1003)
		db.Close()
		return
	}

	file, _, err := r.FormFile("myFile")
    if err != nil {
        fmt.Println("Error Retrieving the File for SendChatImage")
        fmt.Println(err)
        util.Message(w,1002)
        db.Close()
        return
    }

	buffer, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
        util.Message(w,1002)
        db.Close()
        return
    }
    sanatize:=Sanatize(sendChatImageHeader)
	if(sanatize=="Error"){
		util.Message(w,1002)
		db.Close()
		return
	}

	url:=""
	chatResponse:=Chat{}
	imageUploadStatus:="Yes"
	userImageData,status:=image.GetUserImageByRequestId(db,userId,requestId)
	if(status=="Error"){
		// file, _, err := r.FormFile("myFile")
	 //    if err != nil {
	 //        fmt.Println("Error Retrieving the File")
	 //        fmt.Println(err)
	 //        return
	 //    }

		// buffer, err := ioutil.ReadAll(file)
	 //    if err != nil {
	 //        fmt.Println(err)
	 //    }


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
			util.Message(w,1002)
			db.Close()
			return
		}
		size,err:=image.UploadToS3(buffer,filename,bucket,format)
		if(err!=nil){
			//Could Not Upload Image
			fmt.Println(err)
			util.Message(w,3001)
			imageUploadStatus="No"
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
			image.EnterUserImage(db,userImageData)
			imageUploadStatus="Yes"
		}
	

		// signedURL:=""
		// url:=""
		if(accessType=="public"){
			PublicCloudFront:=os.Getenv("publicImageCloudFront")
			url=PublicCloudFront+"/"+filename
		}

	}else{
		// chatResponse=GetChatByRequestId(userId,requestId)
		fmt.Println("Already with this request id")
		if(accessType=="public"){
			PublicCloudFront:=os.Getenv("publicImageCloudFront")
			filename:=userImageData.GeneratedName+"."+userImageData.Format
			url=PublicCloudFront+"/"+filename
		}
	}


	code:=200
	unSyncedChat:=[]Chat{}
	if(imageUploadStatus=="Yes"){
		fmt.Println("Image uploaded")
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
		chatResponse,unSyncedChat,code=ChatInsertDB(db,chat,lastUpdate)
		// db.Create(&chatData)
		if(chat.CreatedAt==chatResponse.CreatedAt){
			e:=UpdateChatTime(db,chat.ChatId,chat.CreatedAt)
			if e!=nil{
				db.Close()
				return
			}
		}
	}else{
		fmt.Println("Image not uploaded")
		code=3001
	}
	


	
	msg:=util.GetMessageDecode(code)
	w.Header().Set("Content-Type", "application/json")
	// p:=&UploadImageResponse{Code:code,Message:msg,ImageId:imageId,URL:url}
	p:=&SendChatImageResponse{Code:code,Message:msg,ImageId:userImageData.ImageId,
				RequestId:requestId,MessageId:chatResponse.MessageId,
				CreatedAt:chatResponse.CreatedAt,MessageType:"image",URL:url,Chat:unSyncedChat}
	fmt.Print("SendChatImageResponse:")
	fmt.Println(*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	fmt.Print("Send Chat image response Error:")
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}