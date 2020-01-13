package Social

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	// "strings"
	"encoding/json"
	profile "miti-microservices/Profile"
   util "miti-microservices/Util"
)

func GetInGroupPool(w http.ResponseWriter, r *http.Request){
	header:=GetInGroupPoolHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie
	userId,dErr:=util.GetUserIdFromSession(sessionId)
	if dErr=="Error"{
		fmt.Println("Session Does not exist")
		util.Message(w,1003)
		return
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}

	getInGroupPoolRequest:=GetInGroupPoolRequest{}
	errQuestionData:=json.Unmarshal(requestBody,&getInGroupPoolRequest)
	if errQuestionData!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}

	profileData:=profile.GetProfileDB(userId)
	pincode:=profileData.Pincode
	createdAt:=util.GetTime()
	gender:=profileData.Gender
	sex:=profileData.Sex
	ipip:=profile.CheckIPIPStatus(userId)
	if(ipip<5){
		util.Message(w,2003)
		return
	}
	interest:=getInGroupPoolRequest.Interest
	requestId:=getInGroupPoolRequest.RequestId
	EnterInGroupPooL(userId,pincode,interest,createdAt,gender,sex)
	// util.Message(w,200)
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&GetInGroupPoolResponse{Code:200,Message:msg,RequestId:requestId,CreatedAt:createdAt}
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}