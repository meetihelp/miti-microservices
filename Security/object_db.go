package Security

import (
	// "log"
	// "fmt"
	// database "miti-microservices/Database"
// 	"github.com/jinzhu/gorm"
//  _ 	"github.com/jinzhu/gorm/dialects/postgres"
)

type PrimaryTrustChain struct{
	UserId string `gorm:"primary_key;type:varchar"  json:"UserId"`
	ChainName string `gorm:"type:varchar" json:"ChainName"`
	ChainId string `gorm:"primary_key;type:varchar" json:"ChatId"`
	Phone1 string `gorm:"type:varchar" json:"Phone1"`
	Name1 string `gorm:"type:varchar" json:"Name1"`
	Phone2 string `gorm:"type:varchar" json:"Phone2"`
	Name2 string `gorm:"type:varchar" json:"Name2"`
	Phone3 string `gorm:"type:varchar" json:"Phone3"`
	Name3 string `gorm:"type:varchar" json:"Name3"`
	Phone4 string `gorm:"type:varchar" json:"Phone4"`
	Name4 string `gorm:"type:varchar" json:"Name4"`
	Phone5 string `gorm:"type:varchar" json:"Phone5"`
	Name5 string `gorm:"type:varchar" json:"Name5"`
	Phone6 string `gorm:"type:varchar" json:"Phone6"`
	Name6 string `gorm:"type:varchar" json:"Name6"`
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
	// db:=database.GetDB()
	// db.AutoMigrate(&PrimaryTrustChain{})
	// db.AutoMigrate(&SecondaryTrustChain{})
}