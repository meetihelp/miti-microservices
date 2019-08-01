package CreateDatabase

import (
	// "log"
	// "fmt"
	"github.com/jinzhu/gorm"
 _ 	"github.com/jinzhu/gorm/dialects/postgres"
)

type Profile struct{
	User_id string `gorm:"primary_key;type:varchar(100)"  json:"user_id"`
	Name string `gorm:"type:varchar(40)" validate:"required" json:"name"`
	Date_of_Birth string `gorm:"type:varchar(100)" validate:"required" json:"dob"`
	Job string `gorm:"type:varchar(30)" validate:"required" "json:"job"`
	Pic_url string `gorm:"type:varchar(100)"  json:"pic_url"`
	Gender string `gorm:"type:varchar(10)" validate:"required" json:"gender"`
	Language string `gorm:"type:varchar(20)" validate:"required" json:"language"`
}

func createProfileTable(db *gorm.DB){
	db.AutoMigrate(&Profile{})
}
