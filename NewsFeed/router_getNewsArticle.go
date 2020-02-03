package NewsFeed

import(
	"net/http"
	"log"
	"fmt"
	gocache "github.com/patrickmn/go-cache"
	"github.com/jinzhu/gorm"
	util "miti-microservices/Util"
	database "miti-microservices/Database"
	getNewsArticleCache "miti-microservices/Database/Cache/NewsFeedCache"
	profile "miti-microservices/Profile"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

func GetNewsArticle(w http.ResponseWriter,r *http.Request){
	ipAddress:=util.GetIPAddress(r)
	getNewsFeedArticleHeader:=GetNewsFeedArticleHeader{}

	content:=GetNewsFeedArticleResponse{}
	statusCode:=0

	responseHeader:=GetNewsFeedArticleResponseHeader{}
	var data map[string]string

	db:=database.DBConnection()
	//Session,TemporarySession,Body,Unmarshal,Sanatize,Database
	list:=[]bool{false,false,false,false,false,false}
	errorList:=util.GetErrorList(list)

	util.GetHeader(r,&getNewsFeedArticleHeader)
	sessionId:=getNewsFeedArticleHeader.Cookie
	userId,status,dbError:=util.GetUserIdFromSession3(db,sessionId)
	errorList.DatabaseError=dbError
	util.APIHitLog("GetNewsArticle",ipAddress,sessionId)
	if status=="Error"{
		// fmt.Println("GetNewsArticle line 39")
		errorList.SessionError=true
	}

	//Read body data
	requestBody,err:=ioutil.ReadAll(r.Body)
	if (err!=nil && !util.ErrorListStatus(errorList)){
		// fmt.Println("GetNewsArticle line 46")
		errorList.BodyReadError=true
	}

	getNewsFeedArticleData :=GetNewsArticleDS{}
	if(!util.ErrorListStatus(errorList)){
		// fmt.Println("GetNewsArticle line 52")
		errGetNewsFeedArticleData:=json.Unmarshal(requestBody,&getNewsFeedArticleData)
		if errGetNewsFeedArticleData!=nil{
			errorList.UnmarshallingError=true
		}	
	}
	
	if(!util.ErrorListStatus(errorList)){
		// fmt.Println("GetNewsArticle line 60")
		sanatizationStatus :=Sanatize(getNewsFeedArticleData)
		if(sanatizationStatus=="Error"){
			errorList.SanatizationError=true
		}
	}

	label:=getNewsFeedArticleData.Label
	if(label==""){
		// fmt.Println("GetNewsArticle line 69")
		label="FoodPorn"
	}
	
	

	var isDone string
	numOfArticle:=0
	if(!util.ErrorListStatus(errorList)){
		// fmt.Println("GetNewsArticle line 82")
		isDone,numOfArticle,dbError=AreAllArticleDone(db,userId)
		errorList.DatabaseError=dbError
		if(isDone=="Yes"){
			// fmt.Println("GetNewsArticle line 87")
			statusCode=5000
		}
	}

	var interest []string
	if(!util.ErrorListStatus(errorList) && isDone=="No"){
		// fmt.Println("GetNewsArticle line 94")
		interest,dbError=profile.GetUserInterest(db,userId)
		errorList.DatabaseError=dbError
	}
	


	news:=[]News{}
	cache:=getNewsArticleCache.GetNewsArticleCache()
	
	numOfLabelArticle:=0
	if(numOfArticle>2){
		// fmt.Println("GetNewsArticle line 106")
		numOfLabelArticle=2
	}else{
		// fmt.Println("GetNewsArticle line 109")
		numOfLabelArticle=numOfArticle
	}

	newsId:=make([]int64,0)
	var nextLabel string
	if(isDone=="No"){
		nextLabel=label
		for true{
			if(util.ErrorListStatus(errorList) || len(news)>0){
				break
			}
			nextLabel=GetNextLabel(nextLabel,interest)
			fmt.Println("Next Label:"+nextLabel)
			var id int64
			if(!util.ErrorListStatus(errorList)){
				// fmt.Println("GetNewsArticle line 74")
				id,dbError=GetLabelId(db,nextLabel,userId)
				errorList.DatabaseError=dbError	
				fmt.Println(id)
			}
			if(!util.ErrorListStatus(errorList)){
				newsId=make([]int64,0)
				news,newsId,dbError=getNews(db,cache,nextLabel,id,numOfLabelArticle)
				errorList.DatabaseError=dbError	
				fmt.Println(newsId)
			}
			
			// fmt.Println("GetNewsArticle line 117")
		}		
	}

	if(!util.ErrorListStatus(errorList) && isDone=="No"){
		// fmt.Println("GetNewsArticle line 122")
		// fmt.Println(newsId)
		fmt.Println("Updating User Feed Status For "+nextLabel)
		fmt.Println(newsId)
		dbError=UpdateUserNewsFeedStatus(db,userId,nextLabel,newsId)
		errorList.DatabaseError=dbError
	}

	if(!util.ErrorListStatus(errorList) && isDone=="No" && statusCode==0){
		// fmt.Println("GetNewsArticle line 128")
		statusCode=200
	}
	
	code:=util.GetCode(errorList)
	if(code==200){
		// fmt.Println("GetNewsArticle line 134")
		content.Code=statusCode
	}else{
		// fmt.Println("GetNewsArticle line 137")
		content.Code=code
	}
	content.Message=util.GetMessageDecode(content.Code)
	newsResponse:=make([]NewsResponse,0)
	for _,g:=range news{
		n:=NewsResponse{}
		n.Id=g.Id
		n.Summary=g.Summary
		n.Sentiment=g.Sentiment
		n.Location=g.Location
		n.Event=g.Event
		n.Label=g.Label
		n.Title=g.Title
		n.ImageURL=g.ImageURL
		n.Flag=g.Flag
		n.ArticleURL=g.ReferenceArticleURL
		newsResponse=append(newsResponse,n)
	}
	content.NewsData=newsResponse

	responseHeader.ContentType="application/json"
    headerBytes:=new(bytes.Buffer)
    json.NewEncoder(headerBytes).Encode(responseHeader)
    responseHeaderBytes:=headerBytes.Bytes()
    if err := json.Unmarshal(responseHeaderBytes, &data); err != nil {
        panic(err)
    }
    w=util.GetResponseFormatHeader(w,data)
	p:=&content
	util.ResponseLog("GetNewsArticle",ipAddress,sessionId,content.Code,*p)
	enc := json.NewEncoder(w)
	err= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}


func getGuiltyPleasure(db *gorm.DB,cache *gocache.Cache,label string,id int64,numOfArticle int)([]GuiltyPleasure,[]int64,bool){
	x,found:=cache.Get(label)
	guiltyPleasure:= []GuiltyPleasure{}
	var dbError bool
	if(!found){
		// fmt.Println("Cache miss for "+label)
		guiltyPleasure,dbError=GetGuiltyPleasure(db,label)
		cache.Set(label,guiltyPleasure,0)
	}else{
		// fmt.Println("Cache hit for "+label)
		guiltyPleasure=x.([]GuiltyPleasure)
		dbError=false
	}

	guiltyPleasureResponse:=make([]GuiltyPleasure,0)
	count:=0
	newsId:=make([]int64,0)
	for _,g:=range guiltyPleasure{
		if(count>numOfArticle){
			break
		}
		if(g.Id>id){
			guiltyPleasureResponse=append(guiltyPleasureResponse,g)
			newsId=append(newsId,g.Id)
			count++
		}
	}

	return guiltyPleasureResponse,newsId,dbError
}

func getNews(db *gorm.DB,cache *gocache.Cache,label string,id int64,numOfArticle int)([]News,[]int64,bool){
	x,found:=cache.Get(label)
	news:= []News{}
	var dbError bool
	if(!found || len(news)==0){
		// fmt.Println("Cache miss for "+label)
		news,dbError=GetNews(db,label,id)
		cache.Set(label,news,0)
	}else{
		// fmt.Println("Cache hit for "+label)
		news=x.([]News)
		dbError=false
	}

	newsResponse:=make([]News,0)
	count:=0
	newsId:=make([]int64,0)
	for _,g:=range news{
		if(count>numOfArticle){
			break
		}
		if(g.Id>id){
			newsResponse=append(newsResponse,g)
			newsId=append(newsId,g.Id)
			count++
		}
	}

	return newsResponse,newsId,dbError
}

func GetNextLabel(label string,interest []string) (string){
	for index,interestLabel :=range interest{
		if(label==interestLabel){
			if(index!=len(interest)-1){
				return interest[index+1]
			}else{
				return "FoodPorn"
			}
		}
	}

	if(label==""){
		return "FoodPorn"
	}
	if(label=="FoodPorn"){
		return "memes"
	}

	if(label=="memes"){
		return "travel"
	}

	if(label=="travel"){
		if(len(interest)>0){
			return interest[0]
		}else{
			return "FoodPorn"
		}
		
	}

	return "FoodPorn"

}

func GetTable(label string) string{
	if(label==""){
		return "Error"
	}

	if(label=="FoodPorn"){
		return "GuiltyPleasure"
	}

	if(label=="memes"){
		return "GuiltyPleasure"
	}

	return "News"
}