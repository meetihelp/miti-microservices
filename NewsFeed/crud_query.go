package NewsFeed

import(
	"fmt"
	database "miti-microservices/Database"
	util "miti-microservices/Util"
	"github.com/jinzhu/gorm"
)

func GetLabelId(db *gorm.DB,label string,userId string) (int64,bool){
	userFeedStatus:=UserFeedStatus{}
	err:=db.Where("label=? AND user_id=?",label,userId).Find(&userFeedStatus).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return userFeedStatus.Id,true
	}
	if(userFeedStatus.UserId==""){
		userFeedStatus.UserId=userId
		userFeedStatus.Label=label
		userFeedStatus.Id=0
		createdAt:=util.GetTime()
		date:=util.GetDateFromTime(createdAt)
		userFeedStatus.UpdatedAt=date
		err:=db.Create(&userFeedStatus).Error
		if(err!=nil){
			return userFeedStatus.Id,true
		}
	}

	return userFeedStatus.Id,false

}

func AreAllArticleDone(db *gorm.DB,userId string)(string,int,bool){
	count:=0
	userFeedStatus:=[]UserFeedStatus{}
	today:=util.GetDateFromTime(util.GetTime())
	err:=db.Where("user_id=? AND updated_at=?",userId,today).Find(&userFeedStatus).Count(&count).Error
	fmt.Print(today+"->Count:")
	fmt.Println(count)
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return "No",0,true
	}
	if(count<40){
		return "No",count,false
	}
	return "Yes",count,false
}

func GetGuiltyPleasure(db *gorm.DB,label string)([]GuiltyPleasure,bool){
	guiltyPleasure:=[]GuiltyPleasure{}
	userFeedStatus:=UserFeedStatus{}
	err:=db.Order("id desc").Limit(1).Where("label=?",label).Find(&userFeedStatus).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return guiltyPleasure,true
	}
	id:=userFeedStatus.Id

	err=db.Table("guilty_pleasure").Order("id").Limit(20).Where("label=? AND id>?",label,id).Find(&guiltyPleasure).Error
	if(err!=nil){
		return guiltyPleasure,true
	}
	return guiltyPleasure,false
}

func GetNews(db *gorm.DB,label string)([]News,bool){
	news:=[]News{}
	userFeedStatus:=UserFeedStatus{}
	err:=db.Order("id desc").Limit(1).Where("label=?",label).Find(&userFeedStatus).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return news,true
	}
	id:=userFeedStatus.Id

	err=db.Table("news").Order("id").Limit(20).Where("label=? AND id>?",label,id).Find(&news).Error
	if(err!=nil){
		return news,true
	}
	return news,false
}

func UpdateUserNewsFeedStatus(db *gorm.DB,userId string,label string,newsId []int64) bool{
	updatedAt:=util.GetDateFromTime(util.GetTime())
	for _,id:=range newsId{
		userFeedStatus:=UserFeedStatus{}
		userFeedStatus.UserId=userId
		userFeedStatus.Label=label
		userFeedStatus.UpdatedAt=updatedAt
		userFeedStatus.Id=id
		err:=db.Create(&userFeedStatus).Error
		if(err!=nil){
			return true
		}
	}
	return false
}

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

func GetArticleAfterId(db *gorm.DB,id int64) ([]News){
	// db:=database.GetDB()
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