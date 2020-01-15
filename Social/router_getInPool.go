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

func GetInPool(w http.ResponseWriter, r *http.Request){
	header:=GetInPoolHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie
	userId,dErr:=util.GetUserIdFromSession(sessionId)
	fmt.Print("GetInPoolHeader")
	fmt.Println(header)
	if dErr=="Error"{
		fmt.Println("Session Does not exist GetInPool")
		util.Message(w,1003)
		return
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}

	getInPoolRequest:=GetInPoolRequest{}
	errQuestionData:=json.Unmarshal(requestBody,&getInPoolRequest)
	if errQuestionData!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}

	fmt.Print("GetInPoolBody:")
	fmt.Println(getInPoolRequest)

	requestId:=getInPoolRequest.RequestId
	profileData:=profile.GetProfileDB(userId)
	pincode:=profileData.Pincode
	createdAt:=util.GetTime()
	gender:=profileData.Gender
	sex:=profileData.Sex
	ipip:=profile.CheckIPIPStatus(userId)
	code:=200
	poolStatus:=PoolStatus{}
	if(ipip<5){
		code=2003
	}else{
		poolStatus=EnterInPooL(userId,pincode,createdAt,gender,sex,requestId)
		code=200
	}

	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(code)
	p:=&GetInPoolResponse{Code:code,Message:msg,IPIP:ipip,RequestId:requestId,PoolStatus:poolStatus}
	fmt.Print("GetInPoolResponse:")
	fmt.Println(*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	
	// util.Message(w,200)

}