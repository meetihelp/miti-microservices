package Image


import(
	"net/http"
	util "miti-microservices/Util"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"os"
	"log"
	"strings"

)

func UploadImage(w http.ResponseWriter, r *http.Request){
	//get buffer of image from user
	uploadImageHeader:=UploadImageHeader{}
	util.GetHeader(r,&uploadImageHeader)
	sessionId:=uploadImageHeader.Cookie
	log.Println("upload Image Cookie:"+sessionId)
	accessType:=uploadImageHeader.AccessType
	accessType=strings.ToLower(accessType)
	log.Println("upload Image AccessType:"+accessType)
	actualFileName:=uploadImageHeader.ActualFileName
	format:=uploadImageHeader.Format
	latitude:=uploadImageHeader.Latitude
	longitude:=uploadImageHeader.Longitude
	dimension:=uploadImageHeader.Dimension
	requestId:=uploadImageHeader.RequestId
	fmt.Println(uploadImageHeader)
	userId,getChatStatus:=util.GetUserIdFromSession(sessionId)
	fmt.Println("Enter upload Image:"+userId)
	// fmt.Println(userId)
	if getChatStatus=="Error"{
		util.Message(w,1003)
		return
	}

	url:=""
	userImageData,status:=GetUserImageByRequestId(userId,requestId)
	fmt.Println("Status:"+status)
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

	    if(accessType!="private" && accessType!="public"){
	    	util.Message(w,1002)
	    	return
	    }

		imageId:=util.GenerateToken()
		generatedName:=util.GenerateToken()
		filename:=generatedName+"."+format
		bucket:=""
		fmt.Println("AccessType:"+accessType)
		if(accessType=="private"){
			bucket=GetPrivateImageBucket()
			fmt.Println("Bucket:"+bucket)
		}else{
			bucket=GetPublicImageBucket()
		}
		size,err:=UploadToS3(buffer,filename,bucket,format)
		if(err!=nil){
			//Could Not Upload Image
			fmt.Println(err)
			util.Message(w,3001)
		}else{
			//Uploaded image Successfully
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
			EnterUserImage(userImageData)

			// signedURL:=""
			// url:=""
			if(accessType=="public"){
				PublicCloudFront:=os.Getenv("publicImageCloudFront")
				url=PublicCloudFront+"/"+filename
			}
	
		}
		// uploadImageResponse.Code=200
		// uploadImageResponse.Message=util.GetMessageDecode(200)
		// uploadImageResponse.URL=signedURL
		// util.Message(w,200)
	}else{
		if(accessType=="public"){
			PublicCloudFront:=os.Getenv("publicImageCloudFront")
			filename:=userImageData.GeneratedName+"."+userImageData.Format
			url=PublicCloudFront+"/"+filename
		}
	}
	code:=200
	msg:=util.GetMessageDecode(code)
	w.Header().Set("Content-Type", "application/json")
	p:=&UploadImageResponse{Code:code,Message:msg,ImageId:userImageData.ImageId,URL:url,RequestId:userImageData.RequestId,CreatedAt:userImageData.CreatedAt}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}