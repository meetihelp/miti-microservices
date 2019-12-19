package NewsFeed

import(
	"net/http"
	// "log"
	"fmt"
	util "miti-microservices/Util"
	"io/ioutil"
	"encoding/json"
)

func GetNewsFeedArticle(w http.ResponseWriter,r *http.Request){
	getNewsFeedArticleHeader:=GetNewsFeedArticleHeader{}
	util.GetHeader(r,&getNewsFeedArticleHeader)
	sessionId:=getNewsFeedArticleHeader.Cookie
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

	getNewsFeedArticleData :=GetNewsFeedArticleDS{}
	errGetNewsFeedArticleData:=json.Unmarshal(requestBody,&getNewsFeedArticleData)
	if errGetNewsFeedArticleData!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}

	articleData:=GetArticle(getNewsFeedArticleData.NewsFeedId)
	if(articleData.NewsFeedId==""){
		util.Message(w,4000)
		return
	}
	SendArticle(w,articleData)

}