package CreateDatabase

import (
	// "log"
	// "fmt"
	"time"
	"github.com/jinzhu/gorm"
 _ 	"github.com/jinzhu/gorm/dialects/postgres"
)


type ReadBy struct{
	User_id string `gorm:"primary_key;varchar(100)"  json:"user_id"`
	Message_id string `gorm:"primary_key;varchar(100)"  json:"message_id"`
	Status string `gorm:"type:varchar(10)" json:"status"`
	ReadAt time.Time `gorm:"type:time"`
}

func createReadByTable(db *gorm.DB){
	db.AutoMigrate(&ReadBy{})
}