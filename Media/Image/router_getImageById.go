package Image

import(
	"net/http"
	util "miti-microservices/Util"
	database "miti-microservices/Database"
	"io/ioutil"
	"encoding/json"
	"log"
	"bytes"
)

func GetImageById(w http.ResponseWriter, r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	getImageByIdHeader:=GetImageByIdHeader{}

	content:=GetImageByIdResponse{}
	statusCode:=0

	responseHeader:=GetImageByIdResponseHeader{}
	var data map[string]string

	db:=database.DBConnection()
	list:=[]bool{false,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	util.GetHeader(r,&getImageByIdHeader)
	sessionId:=getImageByIdHeader.Cookie
	userId,getChatStatus,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	util.APIHitLog("GetImageById",ipAddress,sessionId)
	if getChatStatus=="Error"{
		errorList.SessionError=true
	}

	//Read body data
	requestBody,err:=ioutil.ReadAll(r.Body)
	if (err!=nil && !util.ErrorListStatus(errorList)){
		errorList.BodyReadError=true 
	}

	getImageByIdData :=GetImageByIdRequest{}
	if(!util.ErrorListStatus(errorList)){
		errUserData:=json.Unmarshal(requestBody,&getImageByIdData)
		if(errUserData!=nil){
			errorList.UnmarshallingError=true	
		}
		
	}

	if(!util.ErrorListStatus(errorList)){
		sanatize:=Sanatize(getImageByIdData)
		if(sanatize=="Error"){
			errorList.SanatizationError=true
		}
	}
	imageId:=getImageByIdData.ImageId
	//Check If user has permission to access this image
	var userId2 string
	var access string
	if(!util.ErrorListStatus(errorList)){
		userId2,access,dbError=IsUserPermittedToSeeImage(db,userId,imageId)	
		errorList.DatabaseError=dbError
		if(access=="Error"){
			errorList.LogicError=true
		}
	}

	var imageURL string
	if(!util.ErrorListStatus(errorList)){
		imageURL,dbError=GetImageURL(db,userId2,imageId)
		errorList.DatabaseError=dbError
	}
	
	code:=util.GetCode(errorList)
	if(code==200){
		content.Code=statusCode
	}else{
		content.Code=code
	}
	content.Message=util.GetMessageDecode(content.Code)
	content.ImageURL=imageURL
	responseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(responseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
	p:=&content
	util.ResponseLog("GetImageById",ipAddress,sessionId,content.Code,*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
	
}