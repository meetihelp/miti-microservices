package Authentication
import (
	"net/http"
	"fmt"
	util "miti-microservices/Util"
	// "io/ioutil"
	"encoding/json"
	"bytes"
	"log"
)


func LoadingPage(w http.ResponseWriter,r *http.Request){
	//Get IP Address of Client
	// ipAddress:=util.GetIPAddress(r)
	//Read header of the client packet
	loginHeader:=LoginHeader{}
	util.GetHeader(r,&loginHeader)
	sessionId:=loginHeader.Cookie
	//Check if the user is already logged in? Using session value
	statusCode:=0
	moveTo:=0
	var data map[string]string
	content:=LoadingResponse{}

	userId,loginStatus:=util.GetUserIdFromSession(sessionId)

	log.Println("session "+loginStatus)
	if loginStatus=="Ok"{
		fmt.Println("Loading Page Session Ok")
		statusCode=200
		moveTo:=6
		content.Code=statusCode
		content.MoveTo=moveTo
		content.Message=util.GetMessageDecode(statusCode)
		responseHeader:=LoadingToFeedHeader{}
		responseHeader.ContentType="application/json"
		headerBytes:=new(bytes.Buffer)
		json.NewEncoder(headerBytes).Encode(responseHeader)
		responseHeaderBytes:=headerBytes.Bytes()
		if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        	panic(err)
    	}
		w=util.GetResponseFormatHeader(w,data)
		// util.Message(w,200)
		// return
	}else{
		userId,loginStatus=util.GetUserIdFromTemporarySession(sessionId)
		if loginStatus=="Error"{
			fmt.Println("Loading Page Session Error")
			content,w:=util.GetSessionErrorContent(w)
			p:=&content
			enc := json.NewEncoder(w)
			err:= enc.Encode(p)
			if err != nil {
				log.Fatal(err)
			}
			return
			// util.Message(w,1003)
			// return
			// statusCode=1003
			// moveTo=2
			// content.Code=statusCode
			// content.MoveTo=moveTo
			// content.Message=util.GetMessageDecode(statusCode)
			// responseHeader:=LoadingToLoginHeader{}
			// responseHeader.ContentType="application/json"
			// headerBytes:=new(bytes.Buffer)
			// json.NewEncoder(headerBytes).Encode(responseHeader)
			// responseHeaderBytes:=headerBytes.Bytes()
			// if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
	  //       	panic(err)
	  //   	}
			// w=util.GetResponseFormatHeader(w,data)
		}else{
			// w=temporarySessionCase(w,userId)
			IsUserVerified,IsProfileCreated,Preference:=LoadingPageQuery(userId)
			if !IsUserVerified{
				// util.Message(w,1004)
				fmt.Println("Loading Page user not verified")
				statusCode=1004
				moveTo=3
				content.Code=statusCode
				content.MoveTo=moveTo
				content.Message=util.GetMessageDecode(statusCode)
				responseHeader:=LoadingToOTPHeader{}
				responseHeader.ContentType="application/json"
				headerBytes:=new(bytes.Buffer)
				json.NewEncoder(headerBytes).Encode(responseHeader)
				responseHeaderBytes:=headerBytes.Bytes()
				if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
		        	panic(err)
		    	}
				w=util.GetResponseFormatHeader(w,data)
			}else if !IsProfileCreated{
				// util.Message(w,2002)
				fmt.Println("Loading Page Profile not created")
				statusCode=1005
				moveTo=4
				content.Code=statusCode
				content.MoveTo=moveTo
				content.Message=util.GetMessageDecode(statusCode)
				responseHeader:=LoadingToProfileHeader{}
				responseHeader.ContentType="application/json"
				headerBytes:=new(bytes.Buffer)
				json.NewEncoder(headerBytes).Encode(responseHeader)
				responseHeaderBytes:=headerBytes.Bytes()
				if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
		        	panic(err)
		    	}
				w=util.GetResponseFormatHeader(w,data)
			} else if Preference<NUM_OF_PREFERENCE{
				// SendPreference(w,Preferece,1006)
				fmt.Println("Loading Page Preference not created")
				statusCode=1003
				moveTo=2
				content.Code=statusCode
				content.MoveTo=moveTo
				content.Preference=Preference
				content.Message=util.GetMessageDecode(statusCode)
				responseHeader:=LoadingToPreferenceHeader{}
				responseHeader.ContentType="application/json"
				headerBytes:=new(bytes.Buffer)
				json.NewEncoder(headerBytes).Encode(responseHeader)
				responseHeaderBytes:=headerBytes.Bytes()
				if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
		        	panic(err)
		    	}
				w=util.GetResponseFormatHeader(w,data)
				
			}else{
				statusCode=1004
				moveTo=3
				content.Code=statusCode
				content.MoveTo=moveTo
				content.Preference=Preference
				content.Message=util.GetMessageDecode(statusCode)
				responseHeader:=LoadingToPreferenceHeader{}
				responseHeader.ContentType="application/json"
				headerBytes:=new(bytes.Buffer)
				json.NewEncoder(headerBytes).Encode(responseHeader)
				responseHeaderBytes:=headerBytes.Bytes()
				if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
		        	panic(err)
		    	}
				w=util.GetResponseFormatHeader(w,data)
			}
		}
	}
	//Send Response
	p:=&content
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}

func temporarySessionCase(w http.ResponseWriter,userId string){
	// IsUserVerified,IsProfileCreated,Preferece:=LoadingPageQuery(userId)
	// if !IsUserVerified{
	// 	// util.Message(w,1004)
	// 	statusCode=1004
	// 	moveTo=3
	// 	content:=LoadingToOTP{}
	// 	content.Code=statusCode
	// 	content.moveTo=moveTo
	// 	content.Message=util.GetMessageDecode(statusCode)
	// 	responseHeader:=LoginToOTPHeader{}
	// 	responseHeader.ContentType="application/json"
	// 	w=GetResponseFormatHeader(responseHeader)
	// 	return w,content
	// }

	// if !IsProfileCreated{
	// 	// util.Message(w,2002)
	// 	statusCode=1005
	// 	moveTo=4
	// 	content:=LoadingToProfile{}
	// 	content.Code=statusCode
	// 	content.moveTo=moveTo
	// 	content.Message=util.GetMessageDecode(statusCode)
	// 	responseHeader:=LoginToProfileHeader{}
	// 	responseHeader.ContentType="application/json"
	// 	w=GetResponseFormatHeader(responseHeader)
	// 	return w,content
	// }

	// if Preferece<6{
	// 	// SendPreference(w,Preferece,1006)
	// 	statusCode=1003
	// 	moveTo=2
	// 	content:=LoadingToPreference{}
	// 	content.Code=statusCode
	// 	content.moveTo=moveTo
	// 	content.Preferece=Preferece
	// 	content.Message=util.GetMessageDecode(statusCode)
	// 	responseHeader:=LoginToPreferenceHeader{}
	// 	responseHeader.ContentType="application/json"
	// 	w=GetResponseFormatHeader(responseHeader)
	// 	return w,content
	// }

	// util.Message(w,1003)
}
