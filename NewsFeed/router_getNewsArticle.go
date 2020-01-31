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
		errorList.SessionError=true
	}

	//Read body data
	requestBody,err:=ioutil.ReadAll(r.Body)
	if (err!=nil && !util.ErrorListStatus(errorList)){
		errorList.BodyReadError=true
	}

	getNewsFeedArticleData :=GetNewsArticleDS{}
	if(!util.ErrorListStatus(errorList)){
		errGetNewsFeedArticleData:=json.Unmarshal(requestBody,&getNewsFeedArticleData)
		if errGetNewsFeedArticleData!=nil{
			errorList.UnmarshallingError=true
		}	
	}
	
	if(!util.ErrorListStatus(errorList)){
		sanatizationStatus :=Sanatize(getNewsFeedArticleData)
		if(sanatizationStatus=="Error"){
			errorList.SanatizationError=true
		}
	}

	label:=getNewsFeedArticleData.Label
	id,dbError:=GetLabelId(db,label,userId)
	errorList.DatabaseError=dbError

	nextLabel:=GetNextLabel(label)
	var isDone string
	numOfArticle:=0
	if(!util.ErrorListStatus(errorList)){
		isDone,dbError=AreAllArticleDone(db,userId)
		errorList.DatabaseError=dbError
		if(isDone=="Yes"){
			statusCode=5000
		}
	}

	guiltyPleasure:=[]GuiltyPleasure{}
	cache:=getNewsArticleCache.GetNewsArticleCache()
	
	numOfLabelArticle:=0
	if(numOfArticle>2){
		numOfLabelArticle=2
	}else{
		numOfLabelArticle=numOfArticle
	}

	if(!util.ErrorListStatus(errorList) && isDone=="No"){
		guiltyPleasure,dbError=getGuiltyPleasure(db,cache,nextLabel,id,numOfLabelArticle)
		errorList.DatabaseError=dbError
	}


	if(!util.ErrorListStatus(errorList) && isDone=="No"){
		statusCode=200
	}
	
	code:=util.GetCode(errorList)
	if(code==200){
		content.Code=statusCode
	}else{
		content.Code=code
	}
	content.Message=util.GetMessageDecode(content.Code)
	newsResponse:=make([]NewsResponse,0)
	for _,g:=range guiltyPleasure{
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


func getGuiltyPleasure(db *gorm.DB,cache *gocache.Cache,label string,id int64,numOfArticle int)([]GuiltyPleasure,bool){
	x,found:=cache.Get(label)
	guiltyPleasure:= []GuiltyPleasure{}
	var dbError bool
	if(!found){
		fmt.Println("Cache miss for "+label)
		guiltyPleasure,dbError=GetGuiltyPleasure(db,label)
	}else{
		fmt.Println("Cache hit for foodPorn")
		guiltyPleasure=x.([]GuiltyPleasure)
		dbError=false
	}

	guiltyPleasureResponse:=make([]GuiltyPleasure,0)
	count:=0
	for _,g:=range guiltyPleasure{
		if(count>numOfArticle){
			break
		}
		if(g.Id>id){
			guiltyPleasureResponse=append(guiltyPleasureResponse,g)
			count++
		}
	}

	return guiltyPleasureResponse,dbError
}

func GetNextLabel(label string) string{
	if(label==""){
		return "FoodPorn"
	}
	if(label=="FoodPorn"){
		return "memes"
	}

	if(label=="memes"){
		return "FoodPorn"
	}

	return "FoodPorn"

}