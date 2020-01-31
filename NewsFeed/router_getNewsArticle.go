package NewsFeed

import(
	"net/http"
	// "log"
	"fmt"
	// "time"
	util "miti-microservices/Util"
	database "miti-microservices/Database"
	getNewsArticleCache "miti-microservices/Database/Cache/NewsFeedCache"
	"io/ioutil"
	"strconv"
	"encoding/json"
)

func GetNewsArticle(w http.ResponseWriter,r *http.Request){
	getNewsFeedArticleHeader:=GetNewsFeedArticleHeader{}
	util.GetHeader(r,&getNewsFeedArticleHeader)
	sessionId:=getNewsFeedArticleHeader.Cookie
	db:=database.DBConnection()
	userId,status:=util.GetUserIdFromSession2(db,sessionId)
	fmt.Println(userId)
	if status=="Error"{
		util.Message(w,1003)
		db.Close()
		return
	}

	//Read body data
	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		db.Close()
		return 
	}

	getNewsFeedArticleData :=GetNewsArticleDS{}
	errGetNewsFeedArticleData:=json.Unmarshal(requestBody,&getNewsFeedArticleData)
	if errGetNewsFeedArticleData!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		db.Close()
		return 
	}
	cache:=getNewsArticleCache.GetNewsArticleCache()
	string_news_id:=strconv.FormatInt(getNewsFeedArticleData.Id,10)
	articleData:=[]News{}
	x,found:=cache.Get(string_news_id)
	if(!found){
		fmt.Println("Cache miss for "+string_news_id)
		articleData=GetArticleAfterId(db,getNewsFeedArticleData.Id)
		cache.Set(string_news_id,articleData,0)
	}else{
		fmt.Println("Cache hit for "+string_news_id)
		articleData=x.([]News)
	}
	// articleData:=GetArticleAfterId(db,getNewsFeedArticleData.Id)
	// if(articleData.NewsFeedId==""){
	// 	util.Message(w,4000)
	// 	return
	// }
	SendNewsArticle(w,articleData)
	db.Close()
}