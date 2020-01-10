package NewsFeed

import(
	"net/http"
	// "log"
	"fmt"
	util "miti-microservices/Util"
	"io/ioutil"
	"encoding/json"
)

func UpdateNewsFeedReaction(w http.ResponseWriter,r *http.Request){
	newsFeedReactionHeader:=NewsFeedReactionHeader{}
	util.GetHeader(r,&newsFeedReactionHeader)
	sessionId:=newsFeedReactionHeader.Cookie
	userId,status:=util.GetUserIdFromSession(sessionId)
	fmt.Println(userId)
	if status=="Error"{
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

	newsFeedReactionData :=NewsFeedReactionDS{}
	errNewsFeedReactionData:=json.Unmarshal(requestBody,&newsFeedReactionData)
	if errNewsFeedReactionData!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}

	UpdateNewsFeedReactionDB(userId,newsFeedReactionData)
	util.Message(w,200)
}

