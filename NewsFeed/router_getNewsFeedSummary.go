package NewsFeed

import(
	"net/http"
	// "log"
	"fmt"
	util "miti-microservices/Util"
	"io/ioutil"
	"encoding/json"
)

func GetNewsFeedSummary(w http.ResponseWriter,r *http.Request){
	getNewsFeedSummaryHeader:=GetNewsFeedSummaryHeader{}
	util.GetHeader(r,&getNewsFeedSummaryHeader)
	sessionId:=getNewsFeedSummaryHeader.Cookie
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

	getNewsFeedSummaryData :=GetNewsFeedSummaryDS{}
	errGetNewsFeedSummaryData:=json.Unmarshal(requestBody,&getNewsFeedSummaryData)
	if errGetNewsFeedSummaryData!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}

	summaryData:=GetSummary(getNewsFeedSummaryData.NewsFeedId)
	if(summaryData.NewsFeedId==""){
		util.Message(w,4000)
		return
	}
	SendSummary(w,summaryData)

}