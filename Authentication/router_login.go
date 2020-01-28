package Authentication
import (
	"net/http"
	// "fmt"
	util "miti-microservices/Util"
	database "miti-microservices/Database"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"log"
)


func Login(w http.ResponseWriter,r *http.Request){
	//Get IP Address of Client
	ipAddress:=util.GetIPAddress(r)
	//Read header of the client packet
	loginHeader:=LoginHeader{}
	util.GetHeader(r,&loginHeader)
	// sessionId:=loginHeader.Cookie
	statusCode:=0
	moveTo:=0
	var data map[string]string
	content:=LoginResponse{}
	responseHeader:=LoginToOTPHeader{}
	db:=database.DBConnection()
	//Check if the user is already logged in? Using session value
	// userId,loginStatus:=util.GetUserIdFromSession(sessionId)
	// fmt.Println("session "+loginStatus)
	// if loginStatus=="Ok"{
	// 	util.Message(w,200)
	// 	return
	// }
	//Check if User is verified or not
	//session of Unverified user is stored separately to reduce the risk.... 
	//...of accesing the unauthorized data without verification 
	// userId,loginStatus=util.GetUserIdFromTemporarySession(sessionId)
	// if loginStatus=="Ok"{
	// 	util.Message(w,1005)
	// 	return
	// }

	//Read body data
	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		statusCode=1002
		moveTo=0
		content.Code=statusCode
		content.MoveTo=moveTo
		content.Message=util.GetMessageDecode(statusCode)
		headerBytes:=new(bytes.Buffer)
		json.NewEncoder(headerBytes).Encode(responseHeader)
		responseHeaderBytes:=headerBytes.Bytes()
		if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        	panic(err)
    	}
		w=util.GetResponseFormatHeader(w,data)
		p:=&content
		enc := json.NewEncoder(w)
		err:= enc.Encode(p)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println("Could not read body")
		// util.Message(w,1000)
		db.Close()
		return 
	}

	//UNMARSHILING DATA
	userData :=User{}
	errUserData:=json.Unmarshal(requestBody,&userData)
	if errUserData!=nil{
		statusCode=1001
		moveTo=0
		content.Code=statusCode
		content.MoveTo=moveTo
		content.Message=util.GetMessageDecode(statusCode)
		headerBytes:=new(bytes.Buffer)
		json.NewEncoder(headerBytes).Encode(responseHeader)
		responseHeaderBytes:=headerBytes.Bytes()
		if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        	panic(err)
    	}
		w=util.GetResponseFormatHeader(w,data)
		p:=&content
		enc := json.NewEncoder(w)
		err:= enc.Encode(p)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println("Could not Unmarshall user data")
		// util.Message(w,1001)
		db.Close()
		return 
	}
	// fmt.Println(userData);
	//Check if the user data is proper or not
	sanatizationStatus :=Sanatize(userData)
	if sanatizationStatus =="Error"{
		statusCode=1002
		moveTo=0
		content.Code=statusCode
		content.MoveTo=moveTo
		content.Message=util.GetMessageDecode(statusCode)
		headerBytes:=new(bytes.Buffer)
		json.NewEncoder(headerBytes).Encode(responseHeader)
		responseHeaderBytes:=headerBytes.Bytes()
		if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        	panic(err)
    	}
		w=util.GetResponseFormatHeader(w,data)
		p:=&content
		enc := json.NewEncoder(w)
		err:= enc.Encode(p)
		if err != nil {
			log.Fatal(err)
		}
		db.Close()
		return
	}

	//Check if the credentials given by user is Proper or not
	_,loginStatus:=CheckUserCredentials(userData)
	// if loginStatus=="WrongPassword"{
	// 	util.Message(w,1502)
	// 	return
	// }
	if (loginStatus=="NoUser" || loginStatus=="Unverified" || loginStatus=="Ok"){
		userId,dbStatus:=EnterUserData(userData)
		// fmt.Println(userId)
		// if(dbStatus==1){
		// 	cookie:=util.InsertTemporarySession(userId,ipAddress)
		// 	// w.Header().Set("Miti-Cookie",cookie)
		// 	// util.Message(w,1501)
		// }
		if(dbStatus==1){
			userId,_=GetUserIdFromPhone(db,userData.Phone)
		}
		cookie:=util.InsertTemporarySession(db,userId,ipAddress)
		statusCode=200
		moveTo=3
		content.Code=statusCode
		content.MoveTo=moveTo
		content.Message=util.GetMessageDecode(statusCode)
		responseHeader.MitiCookie=cookie
		headerBytes:=new(bytes.Buffer)
		json.NewEncoder(headerBytes).Encode(responseHeader)
		responseHeaderBytes:=headerBytes.Bytes()
		if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        	panic(err)
    	}
		w=util.GetResponseFormatHeader(w,data)
		p:=&content
		enc := json.NewEncoder(w)
		err= enc.Encode(p)
		if err != nil {
			log.Fatal(err)
		}
		db.Close()
		return
	}
	// else if loginStatus=="Unverified"{
	// 	cookie:=util.InsertTemporarySession(userId,ipAddress)
	// 	statusCode=200
	// 	moveTo=3
	// 	content:=LoginToOTP{}
	// 	content.Status=statusCode
	// 	content.MoveTo=moveTo
	// 	content.Message=util.GetMessageDecode(statusCode)
	// 	responseHeader:=LoginToOTPHeader{}
	// 	responseHeader.MitiCookie=cookie
	// 	headerBytes:=new(bytes.Buffer)
	// 	json.NewEncoder(headerBytes).Encode(responseHeader)
	// 	responseHeaderBytes:=headerBytes.Bytes()
	// 	if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
 //        	panic(err)
 //    	}
	// 	w=util.GetResponseFormatHeader(w,data)
	// 	//Send Response
	// 	p:=&content
	// 	enc := json.NewEncoder(w)
	// 	err:= enc.Encode(p)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	// cookie:=util.InsertTemporarySession(userId,ipAddress)
	// 	// w.Header().Set("Miti-Cookie",cookie)
	// 	// util.Message(w,1005)
	// 	// return
	// } 
	// if loginStatus=="Ok"{
	// 	cookie:=util.InsertTemporarySession(userId,ipAddress)
	// 	w.Header().Set("Miti-Cookie",cookie)
	// 	util.Message(w,200)
	// 	return
	// }


}
