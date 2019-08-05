package CreateDatabase

import (
	// "log"
	// "fmt"
	"github.com/jinzhu/gorm"
 _ 	"github.com/jinzhu/gorm/dialects/postgres"
)

type Email_verification struct{
	User_id string `gorm:"primary_key;type:varchar(100)"  validate:"required"`
	Verification_token string `gorm:"primary_key;varchar(100)" validate:"required"`
}

func createVerification_EmailTable(db *gorm.DB){
	db.AutoMigrate(&Email_verification{})
}