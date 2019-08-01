package CreateDatabase

import (
	// "log"
	// "fmt"
	"github.com/jinzhu/gorm"
 _ 	"github.com/jinzhu/gorm/dialects/postgres"
)

type Session struct{
	User_id string `gorm:"primary_key;type:varchar(100)"  validate:"required"`
	Session_id string `gorm:"primary_key;type:varchar(100)"  validate:"required"`
	Browser_agent string `gorm:"type:varchar(30)" validate:"required"`
	Latitude string `gorm:"type:varchar(30)" validate:"required"`
	Longitude string `gorm:"type:varchar(30)" validate:"required"`
	OS string `gorm:"type:varchar(30)" validate:"required"`
}

func createSessionTable(db *gorm.DB){
	db.AutoMigrate(&Session{})
}
