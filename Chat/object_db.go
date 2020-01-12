	package Chat

import (
	// "log"
	// "fmt"
	// database "miti-microservices/Database"
	// "time"
	// "github.com/jinzhu/gorm"
 // _ 	"github.com/jinzhu/gorm/dialects/postgres"
)

type ChatDetail struct{
	TempUserId string `gorm:"varchar(100)"  json:"TempUserId"`
	ActualUserId string `gorm:"primary_key;varchar(100)"  json:"ActualUserId"`
	ChatId string `gorm:"primary_key;varchar(100)"  json:"ChatId"`
	ChatType string `gorm:"varchar(100)"  json:"ChatType"`
	CreatedAt string `gorm:"varchar(100)" json:"CreatedAt"`
	LastUpdate string `gorm:"varchar(100)" json:"LastUpdate"`
	UserIndex int `gorm:"type:int" json:"Index"`
	Name string `gorm:"type:varchar" json:"Name"`
}
type MessageRequest struct{
	UserId string `gorm:"primary_key;type:varchar" json:"UserId"`
	UserPhone string `gorm:"type:varchar" json:"UserPhone"`
	Phone string `gorm:"primary_key;type:varchar" json:"Phone"`
	RequestId string `gorm:"varchar" json:"RequestId"`
	AcceptStatus string `gorm:"varchar" json:"AcceptStatus"`
	MessageType string `gorm:"varchar"  json:"MessageType"`
	MessageContent string `gorm:"varchar"  json:"MessageContent"`
	CreatedAt string `gorm:"varchar" json:"CreatedAt"`

}
type Chat struct{
	UserId string `gorm:"varchar(100)"  json:"UserId"`
	ChatId string `gorm:"primary_key;varchar(100)"  json:"ChatId"`
	MessageId string `gorm:"primary_key;unique;varchar(100)"  json:"MessageId"`
	MessageType string `gorm:"varchar(100)"  json:"MessageType"`
	MessageContent string `gorm:"varchar(100)"  json:"MessageContent"`
	CreatedAt string `gorm:"type:varchar" json:"CreatedAt"`
	RequestId string `gorm:"type:varchar" json:"RequestId"`
}
type ReadBy struct{
	UserId string `gorm:"primary_key;varchar(100)"  json:"UserId"`
	MessageId string `gorm:"primary_key;varchar(100)"  json:"MessageId"`
	Status string `gorm:"type:varchar(10)" json:"Status"`
	ReadAt string `gorm:"type:varchar(100)" json"ReadAt"`
}

func init(){	
	// db:=database.GetDB()
	// db.AutoMigrate(&ChatDetail{})
	// db.AutoMigrate(&Chat{})
	// db.AutoMigrate(&ReadBy{})
}