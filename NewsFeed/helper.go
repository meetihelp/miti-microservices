package NewsFeed

import(
	"net/http"
	"encoding/json"	
	util "miti-microservices/Util"
	"log"
)

const(
	NUMBEROFARTICLE=2
	NUMBEROFARTICLE2=1
)

func SendSummary(w http.ResponseWriter,summaryData NewsFeedSummary){
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	msg:=util.GetMessageDecode(200)
	p:=&SummaryResponse{Code:200,Message:msg,NewsFeedId:summaryData.NewsFeedId,Summary:summaryData.Summary}
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}

// func SendArticle(w http.ResponseWriter,articleData NewsFeedArticle){
// 	w.Header().Set("Content-Type", "application/json")
// 	enc := json.NewEncoder(w)
// 	msg:=util.GetMessageDecode(200)
// 	p:=&ArticleResponse{Code:200,Message:msg,NewsFeedId:articleData.NewsFeedId,Article:articleData.Article}
// 	err:= enc.Encode(p)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func SendArticle(w http.ResponseWriter,articleData NewsFeedArticleResponse){
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	msg:=util.GetMessageDecode(200)
	p:=&ArticleResponse{Code:200,Message:msg,Response:articleData}
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}

func SendNewsArticle(w http.ResponseWriter,data []News){
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	msg:=util.GetMessageDecode(200)
	news:=[]NewsResponse{}
	for _,d:=range data{
		n:=NewsResponse{}
		n.Id=d.Id
		n.Summary=d.Summary
		n.Sentiment=d.Sentiment
		n.Location=d.Location
		n.Event=d.Event
		n.Label=d.Label
		n.Title=d.Title
		n.ImageURL=d.ImageURL
		n.Flag=d.Flag
		news=append(news,n)
	}
	p:=&NewsArticleResponse{Code:200,Message:msg,NewsData:news}
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}