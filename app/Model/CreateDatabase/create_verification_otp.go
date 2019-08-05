package CreateDatabase

import (
	// "log"
	// "fmt"
	"github.com/jinzhu/gorm"
 _ 	"github.com/jinzhu/gorm/dialects/postgres"
)

type OTP_verification struct{
	User_id string `gorm:"primary_key;type:varchar(100)"  validate:"required"`
	Verification_otp int `gorm:"primary_key;varchar(100)" validate:"required"`
}

func createVerification_OTPTable(db *gorm.DB){
	db.AutoMigrate(&OTP_verification{})
}