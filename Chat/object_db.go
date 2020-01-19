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
	RequestId string `gorm:"type:varchar" json:"RequestId"`
}
type MessageRequest struct{
	SenderUserId string `gorm:"primary_key;type:varchar" json:"SenderUserId"`
	SenderPhone string `gorm:"type:varchar" json:"SenderPhone"`
	SenderName string `gorm:"type:varchar" json:"SenderName"`
	Phone string `gorm:"primary_key;type:varchar" json:"Phone"`
	RequestId string `gorm:"varchar" json:"RequestId"`
	Status string `gorm:"varchar" json:"AcceptStatus"`
	MessageId string `gorm:"type:varchar"  json:"MessageId"`
	MessageType string `gorm:"varchar"  json:"MessageType"`
	MessageContent string `gorm:"varchar"  json:"MessageContent"`
	CreatedAt string `gorm:"varchar" json:"CreatedAt"`
	ActionRequestId string `gorm:"varchar" json:"ActionRequestId"`
	UpdatedAt string `gorm:"varchar" json:"UpdatedAt"`

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
	// db.AutoMigrate(&MessageRequest{})
	// db.AutoMigrate(&ChatDetail{})
	// db.AutoMigrate(&Chat{})
	// db.AutoMigrate(&ReadBy{})
}