package NewsFeed

import(
	database "miti-microservices/Database"
)
func GetSummary(newsFeedId string) NewsFeedSummary{
	db:=database.GetDB()
	summaryData:=NewsFeedSummary{}
	db.Where("news_feed_id=?",newsFeedId).Find(&summaryData)
	return summaryData
}

func GetArticle(newsFeedId string) NewsFeedArticle{
	db:=database.GetDB()
	articleData:=NewsFeedArticle{}
	db.Where("news_feed_id=?",newsFeedId).Find(&articleData)
	return articleData
}