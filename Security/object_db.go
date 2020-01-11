package Security

import (
	// "log"
	// "fmt"
	database "miti-microservices/Database"
// 	"github.com/jinzhu/gorm"
//  _ 	"github.com/jinzhu/gorm/dialects/postgres"
)

type PrimaryTrustChain struct{
	UserId string `gorm:"primary_key;type:varchar"  json:"UserId"`
	ChainName string `gorm:"primary_key;type:varchar" json:"ChainName"`
	Phone1 string `gorm:"type:varchar" json:"Phone1"`
	Phone2 string `gorm:"type:varchar" json:"Phone2"`
	Phone3 string `gorm:"type:varchar" json:"Phone3"`
	Phone4 string `gorm:"type:varchar" json:"Phone4"`
	Phone5 string `gorm:"type:varchar" json:"Phone5"`
	Phone6 string `gorm:"type:varchar" json:"Phone6"`
	RequestId string `gorm:"type:varchar" json:"RequestId"`
	UpdatedAt string `gorm:"type:varchar" json:"UpdatedAt"`
}


type SecondaryTrustChain struct{
	UserId string `gorm:"primary_key;type:varchar"  json:"UserId"`
	ChatId string `gorm:"primary_key;type:varchar" json:"ChatId"`
	CreatedAt string `gorm:"type:varchar" json:"CreatedAt"`
	RequestId string `gorm:"type:varchar" json:"RequestId"`
}

func init(){
	db:=database.GetDB()
	db.AutoMigrate(&PrimaryTrustChain{})
	db.AutoMigrate(&SecondaryTrustChain{})
}