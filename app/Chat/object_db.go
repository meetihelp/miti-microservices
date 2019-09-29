package Chat

import (
	// "log"
	// "fmt"
	database "app/Database"
	"time"
	// "github.com/jinzhu/gorm"
 // _ 	"github.com/jinzhu/gorm/dialects/postgres"
)

type ChatDetail struct{
	Temp_User_id string `gorm:"varchar(100)"  json:"TempUserId"`
	Actual_User_id string `gorm:"primary_key;varchar(100)"  json:"ActualUserId"`
	Chat_id string `gorm:"primary_key;varchar(100)"  json:"ChatId"`
	Chat_type string `gorm:"varchar(100)"  json:"ChatType"`
	CreatedAt string `gorm:"varchar(100)" json:"CreatedAt"`
	LastUpdate string `gorm:"varchar(100)" json:"LastUpdate"`
	User_index int `gorm:"type:int" json:"Index"`
}
type Chat struct{
	User_id string `gorm:"varchar(100)"  json:"UserId"`
	Chat_id string `gorm:"varchar(100)"  json:"ChatId"`
	Message_id string `gorm:"primary_key;unique;varchar(100)"  json:"MessageId"`
	Message_type string `gorm:"varchar(100)"  json:"MessageType"`
	Message_content string `gorm:"varchar(100)"  json:"MessageContent"`
	CreatedAt string `gorm:"type:varchar" json:"CreatedAt"`
	Index int `gorm:"type:int" json:"Index"`
}
type ReadBy struct{
	User_id string `gorm:"primary_key;varchar(100)"  json:"user_id"`
	Message_id string `gorm:"primary_key;varchar(100)"  json:"message_id"`
	Status string `gorm:"type:varchar(10)" json:"status"`
	ReadAt time.Time `gorm:"type:time"`
}

func init(){	
	db:=database.GetDB()
	db.AutoMigrate(&ChatDetail{})
	db.AutoMigrate(&Chat{})
	db.AutoMigrate(&ReadBy{})
}