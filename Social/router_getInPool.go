package Social

import(
	"fmt"
	"net/http"
	// "log"
	// "io/ioutil"
	// "strings"
	// "encoding/json"
	profile "miti-microservices/Profile"
   util "miti-microservices/Util"
)

func GetInPool(w http.ResponseWriter, r *http.Request){
	header:=GetInPoolHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie
	userId,dErr:=util.GetUserIdFromSession(sessionId)
	if dErr=="Error"{
		fmt.Println("Session Does not exist")
		util.Message(w,1003)
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
	EnterInPooL(userId,pincode,createdAt,gender,sex)
	util.Message(w,200)
}