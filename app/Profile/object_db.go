package Profile

import (
	// "log"
	// "fmt"
	database "app/Database"
// 	"github.com/jinzhu/gorm"
//  _ 	"github.com/jinzhu/gorm/dialects/postgres"
)

type Profile struct{
	UserId string `gorm:"primary_key;type:varchar(100)"  json:"UserId"`
	Name string `gorm:"type:varchar(40)" validate:"required" json:"Name"`
	DateOfBirth string `gorm:"type:varchar(100)" validate:"required" json:"DateOfBirth"`
	Job string `gorm:"type:varchar(30)" validate:"required" "json:"Job"`
	PicUrl string `gorm:"type:varchar(100)"  json:"PicUrl"`
	Gender string `gorm:"type:varchar(10)" validate:"required" json:"Gender"`
	Language string `gorm:"type:varchar(20)" validate:"required" json:"Language"`
	Country string `gorm:"type:varchar(30)" json:"Country"`
	Extraversion int `gorm:"type:int" json:"Extraversion"`
	Agreeableness int `gorm:"type:int" json:"Agreeableness"`
	Conscientiousness int `gorm:"type:int" json:"Conscientiousness"`
	EmotionalStability int `gorm:"type:int" json:"EmotionalStability"`
	Intellect int `gorm:"type:int" json:"Intellect"`
	Interest1 string `gorm:"type:varchar(100)" json:"Interest1"`
	Interest2 string `gorm:"type:varchar(100)" json:"Interest2"`
	Interest3 string `gorm:"type:varchar(100)" json:"Interest3"`
	Interest4 string `gorm:"type:varchar(100)" json:"Interest4"`
	Interest5 string `gorm:"type:varchar(100)" json:"Interest5"`
	MakingChoice string `gorm:"type:varchar(20)" json:"MakingChoice"`

}

type QuestionResponse struct{
	UserId string `gorm:"primary_key;type:varchar(100)"  json:"UserId"`
	QuestionId string `gorm:"primary_key";type:varchar(100)" json:"QuestionId"`
	Response string `gorm:"varchar(100)" json:"Response"`
}

type Question struct{
	Id int 
	Content string `gorm:"varchar(1000)" json:"Content"`
}


func init(){
	db:=database.GetDB()
	db.AutoMigrate(&Profile{})
	db.AutoMigrate(&QuestionResponse{})	
	db.AutoMigrate(&Question{})
}
