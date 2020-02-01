package Image


import(
	"net/http"
	util "miti-microservices/Util"
	database "miti-microservices/Database"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"os"
	"log"
	"strings"
	"bytes"

)

func UploadImage(w http.ResponseWriter, r *http.Request){
	//get buffer of image from user
	ipAddress:=util.GetIPAddress(r)
	uploadImageHeader:=UploadImageHeader{}

	content:=UploadImageResponse{}
	statusCode:=0

	responseHeader:=UploadImageResponseHeader{}
	var data map[string]string

	db:=database.DBConnection()
	list:=[]bool{false,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	util.GetHeader(r,&uploadImageHeader)
	sessionId:=uploadImageHeader.Cookie
	accessType:=uploadImageHeader.AccessType
	accessType=strings.ToLower(accessType)
	actualFileName:=uploadImageHeader.ActualFileName
	format:=uploadImageHeader.Format
	latitude:=uploadImageHeader.Latitude
	longitude:=uploadImageHeader.Longitude
	dimension:=uploadImageHeader.Dimension
	requestId:=uploadImageHeader.RequestId

	userId,getChatStatus,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	util.APIHitLog("UploadImage",ipAddress,sessionId)
	if getChatStatus=="Error"{
		errorList.SessionError=true
		fmt.Println("Upload Image line 49")
	}

	file, _, err := r.FormFile("myFile")
    if(err!=nil && !util.ErrorListStatus(errorList)){
    	fmt.Println("Upload Image line 54")
        errorList.BodyReadError=true
    }

	buffer, err := ioutil.ReadAll(file)
    if(err!=nil && !util.ErrorListStatus(errorList)){
    	fmt.Println("Upload Image line 60")
        errorList.BodyReadError=true
    }

    if(!util.ErrorListStatus(errorList)){
    	fmt.Println("Upload Image line 65")
    	sanatize:=Sanatize(uploadImageHeader)
		if(sanatize=="Error"){
			errorList.SanatizationError=true
		}
    }

    url:=""
	// imageUploadStatus:="Yes"
	status:="Error"
	userImageData:=UserImage{}
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("Upload Image line 77")
		userImageData,status,dbError=GetUserImageByRequestId(db,userId,requestId)	
		errorList.DatabaseError=dbError
	}

	var imageId string
	createdAt:=util.GetTime()
	if(status=="Error" && !util.ErrorListStatus(errorList)){
		fmt.Println("Upload Image line 85")
		imageId=util.GenerateToken()
		generatedName:=util.GenerateToken()
		filename:=generatedName+"."+format
		bucket:=""
		fmt.Println("AccessType:"+accessType)
		if(accessType=="private"){
			bucket=GetPrivateImageBucket()
			fmt.Println("Bucket:"+bucket)
		}else if(accessType=="public"){
			bucket=GetPublicImageBucket()
			fmt.Println("Upload Image line 96")
		}else{
			statusCode=1002
			errorList.LogicError=true
		}

		
		if(!util.ErrorListStatus(errorList)){
			size,err:=UploadToS3(buffer,filename,bucket,format)	
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
				userImageData.CreatedAt=createdAt
				dbError:=EnterUserImage(db,userImageData)
				errorList.DatabaseError=dbError
				// imageUploadStatus="Yes"	
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

	if(!util.ErrorListStatus(errorList)){
		statusCode=200
	}
	
	code:=util.GetCode(errorList)
	if(code==200){
		fmt.Println("Upload Image line 141")
		content.Code=statusCode
	}else{
		fmt.Println("Upload Image line 145")
		content.Code=code
	}
	content.Message=util.GetMessageDecode(content.Code)
	content.ImageId=imageId
	content.URL=url
	content.RequestId=requestId
	content.CreatedAt=createdAt
	responseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(responseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
	p:=&content
	util.ResponseLog("UploadImage",ipAddress,sessionId,content.Code,*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}