package Authentication

import(
	"net/http"
	"log"
	// "io/ioutil"
	// "strings"
	"encoding/json"
   util "miti-microservices/Util"
   database "miti-microservices/Database"
   "bytes"
)

func GetTemporaryUserId(w http.ResponseWriter, r *http.Request){
	//CHECK IF USER IS AUTHORIZED TO SEE THE PROFILE
	ipAddress:=util.GetIPAddress(r)
	header:=GetTempUserIdHeader{}

	content:=GetTempUserIdResponse{}
	statusCode:=200

	responseHeader:=GetTempUserIdResponseHeader{}
	var data map[string]string

	db:=database.DBConnection()
	//Session,TemporarySession,Body,Unmarshal,Sanatize,Database
	list:=[]bool{false,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	util.GetHeader(r,&header)
	sessionId:=header.Cookie
	userId,dErr,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	if dErr=="Error"{
		errorList.SessionError=true
	}

	if(!util.ErrorListStatus(errorList)){
		statusCode=200
	}

	code:=util.GetCode(errorList)
	if(code==200){
		content.Code=statusCode
	}else{
		content.Code=code
	}
	content.Message=util.GetMessageDecode(content.Code)
	content.UserId=userId

	responseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(responseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
	p:=&content
	util.ResponseLog("GetTempUserId",ipAddress,sessionId,content.Code,*p)
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()

}