package CreateDatabase

import (
	// "log"
	// "fmt"
	// "time"
	"github.com/jinzhu/gorm"
 _ 	"github.com/jinzhu/gorm/dialects/postgres"
)

type ChatDetail struct{
	Temp_User_id string `gorm:"varchar(100)"  json:"temp_user_id"`
	Actual_User_id string `gorm:"varchar(100)"  json:"actaul_user_id"`
	Chat_id string `gorm:"varchar(100)"  json:"chat_id"`
	Chat_type string `gorm:"varchar(100)"  json:"chat_type"`
	CreatedAt string `gorm:"varchar(100)" json:"createdat"`
	LastUpdate string `gorm:"varchar(100)" json:"lastupdate"`
	User_index int `gorm:"type:int" json:"user_index"`
}

func createChatDetailTable(db *gorm.DB){	
	db.AutoMigrate(&ChatDetail{})
}