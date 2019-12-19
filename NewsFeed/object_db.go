package NewsFeed

import(
	database "miti-microservices/Database"
)

type NewsFeed struct{
	NewsFeedId string `gorm:"primary_key;unique;varchar(100)" json:"NewsFeedId"`
	ReferenceText string `gorm:"type:varchar" json:"ReferenceText"`
	Sentiment string `gorm:"type:varchar" json:"Sentiment"`
	City string `gorm:"type:varchar" json:"City"`
	Event string `gorm:"type:varchar" json:"City"`
	Label string `gorm:"type:varchar" json:"Label"`
	ArticleSize string `gorm:"type:varchar" json:"ArticleSize"`
	Author string `gorm:"type:varchar" json:"Author"`
	ImageURL string `gorm:"type:varchar" json:"ImageURL"`
}

type NewsFeedSummary struct{
	NewsFeedId string `gorm:"primary_key;unique;varchar(100)" json:"NewsFeedId"`
	Summary string `gorm:"type:varchar" json:"Summary"`
}

type NewsFeedArticle struct{
	NewsFeedId string `gorm:"primary_key;unique;varchar(100)" json:"NewsFeedId"`
	Article string `gorm:"type:varchar" json:"Article"`
}

type NewsFeedReaction struct{
	NewsFeedId string `gorm:primary_key;unique;varchar(100)" json:"NewsFeedId"`
	UserId string `gorm:primary_key;unique;varchar(100)" json:"UserId"`
}

func init(){
	db:=database.GetDB()
	db.AutoMigrate(&NewsFeed{})
	db.AutoMigrate(&NewsFeedSummary{})
	db.AutoMigrate(&NewsFeedArticle{})
	db.AutoMigrate(&NewsFeedReaction{})
}