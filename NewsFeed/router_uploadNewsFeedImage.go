package NewsFeed

import(
	"net/http"
	"log"
	"fmt"
	util "miti-microservices/Util"
	image "miti-microservices/Media/Image"
	database "miti-microservices/Database"
	"io/ioutil"
	"encoding/json"
)

func UploadNewsFeedImage(w http.ResponseWriter,r *http.Request){
	uploadNewsFeedImageHeader:=UploadNewsFeedImageHeader{}
	util.GetHeader(r,&uploadNewsFeedImageHeader)
	sessionId:=uploadNewsFeedImageHeader.Cookie
	db:=database.DBConnection()
	userId,status:=util.GetUserIdFromSession2(db,sessionId)
	fmt.Println(userId)
	if status=="Error"{
		util.Message(w,1003)
		return
	}

	//Read body data
	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}

	uploadNewsFeedImageData :=UploadNewsFeedImageRequest{}
	errGetNewsFeedArticleData:=json.Unmarshal(requestBody,&uploadNewsFeedImageData)
	if errGetNewsFeedArticleData!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}

	url:=uploadNewsFeedImageData.URL
	format:=uploadNewsFeedImageData.Format
	resp,err:=http.Get(url)
	buffer,err:=ioutil.ReadAll(resp.Body)
	filename:=util.GenerateToken()
	bucket:=image.GetPublicImageBucket()
	_,err=image.UploadToS3(buffer,filename,bucket,format)
	if(err!=nil){
		util.Message(w,1002)
		db.Close()
		return
	}

	imageURL:=image.GetPublicImageURL(filename)
	code:=200
	msg:=util.GetMessageDecode(code)
	w.Header().Set("Content-Type", "application/json")
	p:=&UploadNewsFeedImageResponse{Code:code,Message:msg,ImageURL:imageURL}
	fmt.Print("GetImageByIdResponse:")
	fmt.Println(*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}