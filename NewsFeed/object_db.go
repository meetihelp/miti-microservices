package NewsFeed

import(
	database "miti-microservices/Database"
)

type News struct{
	Id int64 `gorm:"primary_key;unique;varchar(100)" json:"Id"`
	ReferenceText string `gorm:"type:varchar" json:"ReferenceText"`
	Summary string `gorm:"type:varchar" json:"Summary"`
	Spinned string `gorm:"type:varchar" json:"Spinned"`
	Sentiment string `gorm:"type:varchar" json:"Sentiment"`
	Location string `gorm:"type:varchar" json:"Location"`
	Event string `gorm:"type:varchar" json:"Event"`
	Label string `gorm:"type:varchar" json:"Label"`
	Title string `gorm:"type:varchar" json:"Title"`
	ArticleSize string `gorm:"type:varchar" json:"ArticleSize"`
	Author string `gorm:"type:varchar" json:"Author"`
	ImageURL string `gorm:"type:varchar" json:"ImageURL"`
	ReferenceArticleURL string `gorm:"type:varchar" json:"ReferenceArticleURL"`
	Mitidatetime string `gorm:"type:varchar" json:"Mitidatatime"`
	Flag string `gorm:"type:varchar" json:"Flag"`

}

type GuiltyPleasure struct{
	Id int64 `gorm:"primary_key;unique;varchar(100)" json:"Id"`
	ReferenceText string `gorm:"type:varchar" json:"ReferenceText"`
	Summary string `gorm:"type:varchar" json:"Summary"`
	Spinned string `gorm:"type:varchar" json:"Spinned"`
	Sentiment string `gorm:"type:varchar" json:"Sentiment"`
	Location string `gorm:"type:varchar" json:"Location"`
	Event string `gorm:"type:varchar" json:"Event"`
	Label string `gorm:"type:varchar" json:"Label"`
	Title string `gorm:"type:varchar" json:"Title"`
	ArticleSize string `gorm:"type:varchar" json:"ArticleSize"`
	Author string `gorm:"type:varchar" json:"Author"`
	ImageURL string `gorm:"type:varchar" json:"ImageURL"`
	ReferenceArticleURL string `gorm:"type:varchar" json:"ReferenceArticleURL"`
	Mitidatetime string `gorm:"type:varchar" json:"Mitidatatime"`
	Flag string `gorm:"type:varchar" json:"Flag"`	
}


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
	Id int64 `gorm:"primary_key;varchar(100)" json:"Id"`
	UserId string `gorm:primary_key;varchar(100)" json:"UserId"`
	Reaction string `gorm:"type:varchar" json:"Reaction"`
}

type UserFeedStatus struct{
	UserId string `gorm:"primary_key;varchar" json:"UserId"`
	Label string `gorm:"primary_key;varchar" json:"Label"`
	Id int64 `gorm:"primary_key;varchar" json:"Id"`
	UpdatedAt string `gorm:"varchar" json:"UpdatedAt"`
}

func init(){
	db:=database.GetDB()
	db.AutoMigrate(&UserFeedStatus{})
	// db.AutoMigrate(&GuiltyPleasure{})
	// db.AutoMigrate(&News{})
	// // db.AutoMigrate(&NewsFeed{})
	// // db.AutoMigrate(&NewsFeedSummary{})
	// // db.AutoMigrate(&NewsFeedArticle{})
	// db.AutoMigrate(&NewsFeedReaction{})
}