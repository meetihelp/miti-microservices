package GPS
import(
	"net/http"
	"fmt"
	// CD "miti-microservices/Model/CreateDatabase"
	util "miti-microservices/Util"
	"io/ioutil"
	"encoding/json"
)

func UpdateUserLocation(w http.ResponseWriter, r *http.Request){
	updateUserLocationHeader:=UpdateUserLocationHeader{}
	util.GetHeader(r,&updateUserLocationHeader)
	sessionId:=updateUserLocationHeader.Cookie
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


	updateUserLocationData :=Location{}
	errUserData:=json.Unmarshal(requestBody,&updateUserLocationData)
	if errUserData!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}
	fmt.Print("UpdateUserLocation Body:")
	fmt.Println(updateUserLocationData)

	UpdateUserLocationDB(userId,updateUserLocationData)
	util.Message(w,200)
}