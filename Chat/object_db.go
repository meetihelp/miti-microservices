	package Chat

import (
	// "log"
	// "fmt"
	database "miti-microservices/Database"
	// "time"
	// "github.com/jinzhu/gorm"
 // _ 	"github.com/jinzhu/gorm/dialects/postgres"
)

type ChatDetail struct{
	TempUserId string `gorm:"varchar(100)"  json:"TempUserId"`
	ActualUserId string `gorm:"primary_key;varchar(100)"  json:"ActualUserId"`
	ChatId string `gorm:"primary_key;varchar(100)"  json:"ChatId"`
	ChatType string `gorm:"varchar(100)"  json:"ChatType"`
	CreatedAt string `gorm:"varchar(100)" json:"CreatedAt"`
	LastUpdate string `gorm:"varchar(100)" json:"LastUpdate"`
	UserIndex int `gorm:"type:int" json:"Index"`
}
type Chat struct{
	UserId string `gorm:"varchar(100)"  json:"UserId"`
	ChatId string `gorm:"varchar(100)"  json:"ChatId"`
	MessageId string `gorm:"primary_key;unique;varchar(100)"  json:"MessageId"`
	MessageType string `gorm:"varchar(100)"  json:"MessageType"`
	MessageContent string `gorm:"varchar(100)"  json:"MessageContent"`
	CreatedAt string `gorm:"type:varchar" json:"CreatedAt"`
	Index int `gorm:"type:int" json:"Index"`
}
type ReadBy struct{
	UserId string `gorm:"primary_key;varchar(100)"  json:"UserId"`
	MessageId string `gorm:"primary_key;varchar(100)"  json:"MessageId"`
	Status string `gorm:"type:varchar(10)" json:"Status"`
	ReadAt string `gorm:"type:varchar(100)" json"ReadAt"`
}

func init(){	
	db:=database.GetDB()
	db.AutoMigrate(&ChatDetail{})
	db.AutoMigrate(&Chat{})
	db.AutoMigrate(&ReadBy{})
}