package CreateDatabase

import (
	// "log"
	// "fmt"
	// "time"
	"github.com/jinzhu/gorm"
 _ 	"github.com/jinzhu/gorm/dialects/postgres"
)


type Chat struct{
	User_id string `gorm:"varchar(100)"  json:"user_id"`
	Chat_id string `gorm:"varchar(100)"  json:"chat_id"`
	Message_id string `gorm:"primary_key;unique;varchar(100)"  json:"message_id"`
	Message_type string `gorm:"varchar(100)"  json:"message_type"`
	Message_content string `gorm:"varchar(100)"  json:"message_content"`
	CreatedAt string `gorm:"type:varchar" json:"createdat"`
	Index int `gorm:"type:int" json:"index"`
}
func createChatTable(db *gorm.DB){	
	db.AutoMigrate(&Chat{})
}