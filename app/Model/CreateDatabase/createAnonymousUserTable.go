package CreateDatabase

import (
	// "log"
	// "fmt"
	// "time"
	"github.com/jinzhu/gorm"
 _ 	"github.com/jinzhu/gorm/dialects/postgres"
)


type AnonymousUser struct{
	User_id string `gorm:"primary_key;varchar(100)"  json:"user_id"`
	Anonymous_id string `gorm:"primary_key;unique;varchar(100)"  json:"anonymous_user_id"`
	Chat_id string `gorm:"primary_key;varchar(100)"  json:"chat_id"`
	Status string `gorm:"type:varchar(6)" json:"status"`
	CreatedAt string `gorm:"type:varchar" json:"created_at"`
}
func createAnonymousUserTable(db *gorm.DB){	
	db.AutoMigrate(&AnonymousUser{})
}
