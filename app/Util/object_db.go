package Util
import (
	// "log"
	// "fmt"
	database "app/Database"
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

type UserVerificationSession struct{
	UserId string `gorm:"primary_key;type:varchar(100)"  validate:"required" json:"UserId"`
	UserVerificationSessionId string `gorm:"primary_key;type:varchar(100)"  validate:"required" json:"SessionId"`
	IP string `gorm:"type:varchar(100)" validate:"required" json:"IP"`
	CreatedAt string `gorm:"type:varchar(100)" json:"CreatedAt"`
}

type Match struct{
	UserId1 string `gorm:"primary_key;type:varchar(100)"  validate:"required" json:"UserId1"`
	UserId2 string `gorm:"primary_key;type:varchar(100)"  validate:"required" json:"UserId2"`
}

func init(){
	db:=database.GetDB()
	db.AutoMigrate(&Session{})
	db.AutoMigrate(&UserVerificationSession{})
	db.AutoMigrate(&Match{})
}
