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

// func GetArticle(newsFeedId string) NewsFeedArticle{
// 	db:=database.GetDB()
// 	articleData:=NewsFeedArticle{}
// 	db.Where("news_feed_id=?",newsFeedId).Find(&articleData)
// 	return articleData
// }

func GetArticle(id int64) string{
	db:=database.GetDB()
	articleData:=News{}
	db.Where("id=?",id).Find(&articleData)
	return articleData.Spinned
}

func GetArticleAfterId(id int64) ([]News){
	db:=database.GetDB()
	news:=[]News{}
	db.Where("id>?",id).Limit(NUMBEROFARTICLE).Find(&news)
	return news
}
func UpdateNewsFeedReactionDB(userId string,newsFeedReactionData NewsFeedReactionDS){
	db:=database.GetDB()
	newsFeedReaction:=NewsFeedReaction{}
	db.Where("user_id=? and id=?",userId,newsFeedReactionData.Id).Find(&newsFeedReaction)
	if(newsFeedReaction.Reaction==""){
		newsFeedReaction.UserId=userId
		newsFeedReaction.Id=newsFeedReactionData.Id
		newsFeedReaction.Reaction=newsFeedReactionData.Reaction
		db.Create(&newsFeedReaction)
		return
	}else{
		db.Model(&newsFeedReaction).Where("user_id=? and id=?",userId,newsFeedReactionData.Id).Update("reaction",newsFeedReactionData.Reaction)
		return
	}

}