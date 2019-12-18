package GPS
import(
	"net/http"
	"fmt"
	"strconv"
	// CD "app/Model/CreateDatabase"
	util "app/Util"
	"io/ioutil"
	"encoding/json"
)

func GetUserListByLocation(w http.ResponseWriter, r *http.Request){
	getUserListByLocationHeader:=GetUserListByLocationHeader{}
	util.GetHeader(r,&getUserListByLocationHeader)
	sessionId:=getUserListByLocationHeader.Cookie
	userId,getChatStatus:=util.GetUserIdFromSession(sessionId)
	fmt.Println(userId)
	if getChatStatus=="Error"{
		util.Message(w,1003)
		return
	}

	//Read body data
	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}

	getUserListByLocationData :=GetUserListByLocationDS{}
	errUserData:=json.Unmarshal(requestBody,&getUserListByLocationData)
	if errUserData!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}
	location:=Location{}
	location.Latitude=getUserListByLocationData.Latitude
	location.Longitude=getUserListByLocationData.Longitude
	distance,_:=strconv.ParseFloat(getUserListByLocationData.Distance,64)
	userList:=GetUserListByLocationDB(location,distance)
	// fmt.Println(userList)
	SendUserList(w,userList)
}