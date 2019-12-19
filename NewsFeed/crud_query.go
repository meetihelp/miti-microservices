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

func UpdateNewsFeedReactionDB(userId string,newsFeedReactionData NewsFeedReactionDS){
	db:=database.GetDB()
	newsFeedReaction:=NewsFeedReaction{}
	db.Where("user_id=? and news_feed_id=?",userId,newsFeedReactionData.NewsFeedId).Find(&newsFeedReaction)
	if(newsFeedReaction.Reaction==""){
		newsFeedReaction.UserId=userId
		newsFeedReaction.NewsFeedId=newsFeedReactionData.NewsFeedId
		newsFeedReaction.Reaction=newsFeedReactionData.Reaction
		db.Create(&newsFeedReaction)
		return
	}else{
		db.Model(&newsFeedReaction).Where("user_id=? and news_feed_id=?",userId,newsFeedReactionData.NewsFeedId).Update("reaction",newsFeedReactionData.Reaction)
		return
	}

}