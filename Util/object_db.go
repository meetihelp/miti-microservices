package Util
import (
	// "log"
	// "fmt"
	database "miti-microservices/Database"
	// "time"
	// "github.com/jinzhu/gorm"
 // _ 	"github.com/jinzhu/gorm/dialects/postgres"
)

type Session struct{
	UserId string `gorm:"primary_key;type:varchar(100)"  validate:"required" json:"UserId"`
	SessionId string `gorm:"primary_key;type:varchar(100)"  validate:"required" json:"SessionId"`
	IP string `gorm:"type:varchar(100)" validate:"required" json:"IP"`
	CreatedAt string `gorm:"type:varchar(100)" json:"CreatedAt"`
	// User_agent string `gorm:"type:varchar(30)" validate:"required"`
	// Latitude string `gorm:"type:varchar(30)" validate:"required"`
	// Longitude string `gorm:"type:varchar(30)" validate:"required"`
	// OS string `gorm:"type:varchar(30)" validate:"required"`
}

type TemporarySession struct{
	UserId string `gorm:"primary_key;type:varchar(100)"  validate:"required" json:"UserId"`
	TemporarySessionId string `gorm:"primary_key;type:varchar(100)"  validate:"required" json:"SessionId"`
	IP string `gorm:"type:varchar(100)" validate:"required" json:"IP"`
	CreatedAt string `gorm:"type:varchar(100)" json:"CreatedAt"`
}

type Match struct{
	UserId1 string `gorm:"primary_key;type:varchar(100)"  validate:"required" json:"UserId1"`
	UserId2 string `gorm:"primary_key;type:varchar(100)"  validate:"required" json:"UserId2"`
}

// type OTPVerification struct{
// 	UserId string `gorm:"primary_key;type:varchar(100)"  validate:"required" json:"UserId"`
// 	SessionId string `gorm:"primary_key;type:varchar(100)"  validate:"required" json:"SessionId"`
// 	OTP string `gorm:"primary_key;varchar(100)" validate:"required"  json:"OTP"`
// 	CreatedAt string `gorm:"type:varchar" json:"CreatedAt"`
// }
func init(){
	db:=database.GetDB()
	db.AutoMigrate(&Session{})
	db.AutoMigrate(&TemporarySession{})
	db.AutoMigrate(&Match{})
	// db.AutoMigrate(&OTPVerification{})
}
