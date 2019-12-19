package NewsFeed

import(
	"net/http"
	"encoding/json"	
	util "miti-microservices/Util"
	"log"
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

func SendArticle(w http.ResponseWriter,articleData NewsFeedArticle){
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	msg:=util.GetMessageDecode(200)
	p:=&ArticleResponse{Code:200,Message:msg,NewsFeedId:articleData.NewsFeedId,Article:articleData.Article}
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}