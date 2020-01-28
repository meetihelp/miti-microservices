package Image

// import(
// 	"net/http"
// 	"fmt"
// 	// CD "miti-microservices/Model/CreateDatabase"
// 	util "miti-microservices/Util"
// 	"io/ioutil"
// 	"encoding/json"
// )

// func GetUserImageList(w http.ResponseWriter, r *http.Request){
// 	getUserImageListHeader:=GetUserImageListHeader{}
// 	util.GetHeader(r,&getUserImageListHeader)
// 	sessionId:=getUserImageListHeader.Cookie
// 	_,getChatStatus:=util.GetUserIdFromSession(sessionId)
// 	// fmt.Println(userId)
// 	if getChatStatus=="Error"{
// 		util.Message(w,1003)
// 		return
// 	}

// 	//Read body data
// 	requestBody,err:=ioutil.ReadAll(r.Body)
// 	if err!=nil{
// 		fmt.Println("Could not read body")
// 		util.Message(w,1000)
// 		return 
// 	}

// 	getUserImageListData :=GetUserImageListDS{}
// 	errUserData:=json.Unmarshal(requestBody,&getUserImageListData)
// 	if errUserData!=nil{
// 		fmt.Println("Could not Unmarshall user data")
// 		util.Message(w,1001)
// 		return 
// 	}

// 	imageList:=GetUserImageListDB(getUserImageListData.UserId)
// 	SendImageList(w,imageList)
// }