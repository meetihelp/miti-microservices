package Social

// import(
// 	"fmt"
// 	"net/http"
// 	"log"
// 	"io/ioutil"
// 	// "strings"
// 	"encoding/json"
// 	gps "miti-microservices/GPS"
// 	profile "miti-microservices/Profile"
//    util "miti-microservices/Util"
// )

// func GetInGroupPool(w http.ResponseWriter, r *http.Request){
// 	header:=GetInGroupPoolHeader{}
// 	util.GetHeader(r,&header)
// 	sessionId:=header.Cookie
// 	userId,dErr:=util.GetUserIdFromSession(sessionId)
// 	fmt.Print("GetInGroupPoolHeader:")
// 	fmt.Println(header)
// 	if dErr=="Error"{
// 		fmt.Println("Session Does not exist")
// 		util.Message(w,1003)
// 		return
// 	}

// 	requestBody,err:=ioutil.ReadAll(r.Body)
// 	if err!=nil{
// 		fmt.Println("Could not read body")
// 		util.Message(w,1000)
// 		return 
// 	}

// 	getInGroupPoolRequest:=GetInGroupPoolRequest{}
// 	errQuestionData:=json.Unmarshal(requestBody,&getInGroupPoolRequest)
// 	if errQuestionData!=nil{
// 		fmt.Println("Could not Unmarshall profile data")
// 		util.Message(w,1001)
// 		return
// 	}

// 	fmt.Print("GetInGroupPoolRequest:")
// 	fmt.Println(getInGroupPoolRequest)
// 	interest:=getInGroupPoolRequest.Interest
// 	requestId:=getInGroupPoolRequest.RequestId
// 	interestStatus:=util.CheckInterestAvailablity(interest)
// 	if(interestStatus=="Error"){
// 		util.Message(w,1002)
// 		return
// 	}
// 	pincode:=gps.GetUserCurrentPincode(userId)
// 	profileData:=profile.GetProfileDB(userId)
// 	createdAt:=util.GetTime()
// 	gender:=profileData.Gender
// 	sex:=profileData.Sex
// 	// ipip:=profile.CheckIPIPStatus(userId)
// 	groupPoolStatus:=GroupPoolStatusHelper{}
// 	code:=200
// 	chatId,groupAvailabilty:=GetGroupAvailabilty(userId,pincode,interest,requestId)
// 	if(groupAvailabilty=="already"){
// 		groupPoolStatus=GetGroupPoolStatus(userId,pincode,interest)
// 	}else if(groupAvailabilty=="None"){
// 		groupPoolStatus=EnterInGroupPooL(userId,pincode,interest,createdAt,gender,sex)
// 		code=2004
// 	}else if(groupAvailabilty=="permanent"){
// 		groupPoolStatus=InsertInGroup(chatId,pincode,userId,"permanent",interest,requestId)
// 	}else{
// 		groupPoolStatus=InsertInGroup(chatId,pincode,userId,"temporary",interest,requestId)
// 	}
// 	// code:=200
// 	// groupPoolStatus:=GroupPoolStatusHelper{}
// 	// if(ipip<5){
// 	// 	code=2003
// 	// }else{
// 	// 	// requestId:=getInGroupPoolRequest.RequestId
// 	// 	groupPoolStatus=EnterInGroupPooL(userId,pincode,interest,createdAt,gender,sex)
// 	// }
// 	// util.Message(w,200)
// 	w.Header().Set("Content-Type", "application/json")
// 	msg:=util.GetMessageDecode(code)
// 	p:=&GetInGroupPoolResponse{Code:code,Message:msg,Interest:interest,
// 			CreatedAt:createdAt,Status:groupPoolStatus,RequestId:requestId}
// 	fmt.Print("GetInGroupPoolResponse:")
// 	fmt.Println(*p)
// 	enc := json.NewEncoder(w)
// 	err= enc.Encode(p)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }