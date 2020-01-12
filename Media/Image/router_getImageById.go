package Image

import(
	"net/http"
	"fmt"
	// CD "miti-microservices/Model/CreateDatabase"
	util "miti-microservices/Util"
	"io/ioutil"
	"encoding/json"
	"log"
)

func GetImageById(w http.ResponseWriter, r *http.Request){
	getImageByIdHeader:=GetImageByIdHeader{}
	util.GetHeader(r,&getImageByIdHeader)
	sessionId:=getImageByIdHeader.Cookie
	userId,getChatStatus:=util.GetUserIdFromSession(sessionId)
	// fmt.Println(userId)
	if getChatStatus=="Error"{
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

	getImageByIdData :=GetImageByIdDS{}
	errUserData:=json.Unmarshal(requestBody,&getImageByIdData)
	if errUserData!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}

	// imageId:=GetImageIdFromId(getImageByIdData.Id)
	imageId:=getImageByIdData.ImageId
	//Check If user has permission to access this image
	userId2,access:=IsUserPermittedToSeeImage(userId,imageId)
	if(access=="Error"){
		util.Message(w,5000)
	}
	imageURL:=GetImageURL(userId2,imageId)
	code:=200
	msg:=util.GetMessageDecode(code)
	w.Header().Set("Content-Type", "application/json")
	p:=&GetImageByIdResponse{Code:code,Message:msg,ImageURL:imageURL}
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}



	// path:=GetImagePath(imageId)
	// status:=DoesImageExist(path)
	// if status=="Ok"{
	// 	SendImage(w,path)	
	// }else {
	// 	util.Message(w,1007)
	// }
	
}