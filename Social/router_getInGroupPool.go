package Social

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	// "strings"
	"encoding/json"
	gps "miti-microservices/GPS"
	profile "miti-microservices/Profile"
   util "miti-microservices/Util"
   database "miti-microservices/Database"
   "bytes"
)

func GetInGroupPool(w http.ResponseWriter, r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	header:=GetInGroupPoolHeader{}

	content:=GetInGroupPoolResponse{}
	statusCode:=0

	responseHeader:=GetInGroupPoolResponseHeader{}
	var data map[string]string

	db:=database.DBConnection()
	//Session,TemporarySession,Body,Unmarshal,Sanatize,Database
	list:=[]bool{false,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	util.GetHeader(r,&header)
	sessionId:=header.Cookie
	userId,dErr,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	util.APIHitLog("GetInGroupPool",ipAddress,sessionId)
	if dErr=="Error"{
		fmt.Println("GetInGroupPool Line 36")
		errorList.SessionError=true
	}


	requestBody,err:=ioutil.ReadAll(r.Body)
	if (err!=nil && !util.ErrorListStatus(errorList)){
		fmt.Println("GetInGroupPool Line 43")
		errorList.BodyReadError=true
	}

	getInGroupPoolRequest:=GetInGroupPoolRequest{}
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("GetInGroupPool Line 49")
		errQuestionData:=json.Unmarshal(requestBody,&getInGroupPoolRequest)
		if errQuestionData!=nil{
			errorList.UnmarshallingError=true
		}	
	}
	

	interest:=getInGroupPoolRequest.Interest
	requestId:=getInGroupPoolRequest.RequestId
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("GetInGroupPool Line 60")
		interestStatus:=util.CheckInterestAvailablity(interest)	
		if(interestStatus=="Error"){
			errorList.SanatizationError=true
		}
	}
	
	pincode:="Error"
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("GetInGroupPool Line 69")
		pincode,dbError=gps.GetUserCurrentPincode(db,userId)	
		errorList.DatabaseError=dbError
	}
	
	var gender string
	var sex string
	createdAt:=util.GetTime()

	if(!util.ErrorListStatus(errorList)){
		fmt.Println("GetInGroupPool Line 79")
		profileData,dbError:=profile.GetProfileDB(db,userId)
		errorList.DatabaseError=dbError
		gender=profileData.Gender
		sex=profileData.Sex
	}
	
	
	
	groupPoolStatus:=GroupPoolStatusHelper{}
	var chatId string
	var groupAvailabilty string
	if(!util.ErrorListStatus(errorList) && pincode!="Error"){
		fmt.Println("GetInGroupPool Line 92")
		chatId,groupAvailabilty,dbError=GetGroupAvailabilty(db,userId,pincode,interest,requestId)
		errorList.DatabaseError=dbError
	}
	
	if(!util.ErrorListStatus(errorList) && pincode!="Error"){
		fmt.Println("GetInGroupPool Line 98")
		if(groupAvailabilty=="already"){
			fmt.Println("GetInGroupPool Line 100")
			groupPoolStatus,dbError=GetGroupPoolStatus(db,userId,pincode,interest)
		}else if(groupAvailabilty=="None"){
			fmt.Println("GetInGroupPool Line 103")
			groupPoolStatus,dbError=EnterInGroupPooL(db,userId,pincode,interest,createdAt,gender,sex)
		}else if(groupAvailabilty=="permanent"){
			fmt.Println("GetInGroupPool Line 106")
			groupPoolStatus,dbError=InsertInGroup(db,chatId,pincode,userId,"permanent",interest,requestId)
		}else if(groupAvailabilty=="temporary"){
			fmt.Println("GetInGroupPool Line 109")
			groupPoolStatus,dbError=InsertInGroup(db,chatId,pincode,userId,"temporary",interest,requestId)
		}
		errorList.DatabaseError=dbError
	}

	if(!util.ErrorListStatus(errorList) && pincode=="Error"){
		groupPoolStatus,dbError=InsertInGroup(db,chatId,pincode,userId,"temporary",interest,requestId)
		errorList.DatabaseError=dbError
	}

	if(!util.ErrorListStatus(errorList)){
		fmt.Println("GetInGroupPool Line 116")
		statusCode=200
	}
	
	code:=util.GetCode(errorList)
	if(code==200){
		fmt.Println("GetInGroupPool Line 122")
		content.Code=statusCode
	}else{
		fmt.Println("GetInGroupPool Line 125")
		content.Code=code
	}
	content.Message=util.GetMessageDecode(code)
	content.Interest=interest
	content.CreatedAt=createdAt
	content.Status=groupPoolStatus
	content.RequestId=requestId
	
	responseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(responseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
	p:=&content
	util.ResponseLog("GetInGroupPool",ipAddress,sessionId,content.Code,*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}