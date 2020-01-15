package Social

// import(
// 	"fmt"
// 	"net/http"
// 	// "log"
// 	// "io/ioutil"
// 	// "strings"
// 	// "encoding/json"
// 	profile "miti-microservices/Profile"
//    util "miti-microservices/Util"
// )

// func CancelPoolRouter(w http.ResponseWriter, r *http.Request){
// 	header:=GetInPoolHeader{}
// 	util.GetHeader(r,&header)


// 	sessionId:=header.Cookie
// 	userId,dErr:=util.GetUserIdFromSession(sessionId)
// 	if dErr=="Error"{
// 		fmt.Println("Session Does not exist")
// 		util.Message(w,1003)
// 		return
// 	}

// 	DeleteWaitPool(userId)
// 	profileData:=profile.GetProfileDB(userId)
// 	areaCode:=profileData.Pincode
// 	gender:=profileData.Gender
// 	DeletePool(userId,areaCode,gender)
// 	// util.Message(w,200)

// 	code:=200
// 	w.Header().Set("Content-Type", "application/json")
// 	msg:=util.GetMessageDecode(code)
// 	p:=&CancelGroupPoolResponse{Code:code,Message:msg,RequestId:requestId,Interest:interest}
// 	enc := json.NewEncoder(w)
// 	err:= enc.Encode(p)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }