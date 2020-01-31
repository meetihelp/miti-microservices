package Authentication
import (
	"net/http"
	util "miti-microservices/Util"
	database "miti-microservices/Database"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"log"
)


func Login(w http.ResponseWriter,r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	loginHeader:=LoginHeader{}

	//Initializing response variables
	content:=LoginResponse{}
	moveTo:=0

	//Intializing response header variables
	loginResponseHeader:=LoginResponseHeader{}
	var data map[string]string
	
	db:=database.DBConnection()
	//Session,TemporarySession,Body,Unmarshal,Sanatize,Database
	list:=[]bool{false,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	//Unmarshalling Header
	util.GetHeader(r,&loginHeader)
	sessionId:=loginHeader.Cookie
	util.APIHitLog("Login",ipAddress,sessionId)

	//Body Reading
	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		errorList.BodyReadError=true
	}

	//Unmarshalling Body
	userData :=LoginRequest{}
	if(!errorList.BodyReadError){
		errUserData:=json.Unmarshal(requestBody,&userData)
		if(errUserData!=nil){
			errorList.UnmarshallingError=true
		}
	}
	
	//Sanatizing Body
	if(!errorList.UnmarshallingError){
		util.BodyLog("Login",ipAddress,sessionId,userData)
		sanatizationStatus :=Sanatize(userData)
		if(sanatizationStatus=="Error"){
			errorList.SanatizationError=true
		}
	}

	if(!errorList.SanatizationError){
		userId,dbError:=EnterUserData(db,userData)
		errorList.DatabaseError=dbError
		var otpCode string
		if(!errorList.DatabaseError){
			sessionId,dbError=util.InsertTemporarySession(db,userId,ipAddress)
			errorList.DatabaseError=dbError	
		}
		if(!errorList.DatabaseError){
			otpCode,dbError=InsertOTP(db,userId,sessionId)
			errorList.DatabaseError=dbError	
		}
		if(!errorList.DatabaseError){
			_=SendOTP(userData.Phone,otpCode)	
		}      
	}

	code:=util.GetCode(errorList)
	if(code==200){
		moveTo=3
	}

	content.Code=code
	content.Message=util.GetMessageDecode(code)
	content.MoveTo=moveTo
	loginResponseHeader.MitiCookie=sessionId
	loginResponseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(loginResponseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
	p:=&content
	util.ResponseLog("Login",ipAddress,sessionId,content.Code,*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}
