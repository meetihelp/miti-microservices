package GPS
import(
	"net/http"
	"fmt"
	// CD "miti-microservices/Model/CreateDatabase"
	util "miti-microservices/Util"
	database "miti-microservices/Database"
	"io/ioutil"
	"encoding/json"
)

func UpdateUserLocation(w http.ResponseWriter, r *http.Request){
	updateUserLocationHeader:=UpdateUserLocationHeader{}
	util.GetHeader(r,&updateUserLocationHeader)
	sessionId:=updateUserLocationHeader.Cookie

	db:=database.DBConnection()
	list:=[]bool{false,false,false,false,false,false}
	errorList:=util.GetErrorList(list)


	userId,getChatStatus,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	fmt.Println(userId)
	if getChatStatus=="Error"{
		util.Message(w,1003)
		return
	}

	//Read body data
	requestBody,err:=ioutil.ReadAll(r.Body)
	if (err!=nil && !util.ErrorListStatus(errorList)){
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

	dbError=UpdateUserCurrentLocation(db,userId,updateUserLocationData.Latitude,updateUserLocationData.Longitude)
	errorList.DatabaseError=dbError
	util.Message(w,200)
}