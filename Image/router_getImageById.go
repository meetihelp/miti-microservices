package Image

import(
	"net/http"
	"fmt"
	// CD "miti-microservices/Model/CreateDatabase"
	util "miti-microservices/Util"
	"io/ioutil"
	"encoding/json"
)

func GetImageById(w http.ResponseWriter, r *http.Request){
	getImageByIdHeader:=GetImageByIdHeader{}
	util.GetHeader(r,&getImageByIdHeader)
	sessionId:=getImageByIdHeader.Cookie
	_,getChatStatus:=util.GetUserIdFromSession(sessionId)
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
	path:=GetImagePath(imageId)
	status:=DoesImageExist(path)
	if status=="Ok"{
		SendImage(w,path)	
	}else {
		util.Message(w,1007)
	}
	
}