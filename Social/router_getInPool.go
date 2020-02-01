package Social

import(
	"net/http"
	"log"
	"encoding/json"
	gps "miti-microservices/GPS"
	profile "miti-microservices/Profile"
   	util "miti-microservices/Util"
   	database "miti-microservices/Database"
   	"bytes"
   	"fmt"
)

func GetInPool(w http.ResponseWriter, r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	header:=GetInPoolHeader{}

	content:=GetInPoolResponse{}
	statusCode:=0

	responseHeader:=GetInPoolResponseHeader{}
	var data map[string]string

	db:=database.DBConnection()
	//Session,TemporarySession,Body,Unmarshal,Sanatize,Database
	list:=[]bool{false,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	util.GetHeader(r,&header)
	sessionId:=header.Cookie
	userId,dErr,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	util.APIHitLog("GetInPool",ipAddress,sessionId)
	if dErr=="Error"{
		fmt.Println("GetInPoo; Line 35")
		errorList.SessionError=true
	}

	var pincode string
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("GetInPoo; Line 41")
		pincode,dbError=gps.GetUserCurrentPincode(db,userId)	
		errorList.DatabaseError=dbError
	}

	var profileData profile.Profile
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("GetInPool Line 48")
		profileData,dbError=profile.GetProfileDB(db,userId)	
		errorList.DatabaseError=dbError
	}
	
	
	createdAt:=util.GetTime()
	gender:=profileData.Gender
	sex:=profileData.Sex
	var ipip int
	if(!util.ErrorListStatus(errorList)){
		fmt.Println("GetInPool Line 59")
		ipip,dbError=profile.CheckIPIPStatus(db,userId)
		errorList.DatabaseError=dbError
	}

	poolStatus:=PoolStatusHelper{}
	if(ipip<5){
		fmt.Println("GetInPool Line 66")
		statusCode=2003
	}else{
		fmt.Println("GetInPool Line 69")
		poolStatus,dbError=EnterInPooL(db,userId,pincode,createdAt,gender,sex)
		errorList.DatabaseError=dbError
	}

	if(!util.ErrorListStatus(errorList) && ipip>=5){
		fmt.Println("GetInPool Line 75")
		statusCode=200
	}

	code:=util.GetCode(errorList)
	if(code==200){
		fmt.Println("GetInPool Line 81")
		content.Code=statusCode
	}else{
		fmt.Println("GetInPool Line 84")
		content.Code=code
	}
	content.Message=util.GetMessageDecode(code)
	content.PoolStatus=poolStatus
	content.IPIP=ipip

	responseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(responseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
	p:=&content
	util.ResponseLog("GetInPool",ipAddress,sessionId,content.Code,*p)
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()

}