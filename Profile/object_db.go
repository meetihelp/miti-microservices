package Profile

import (
	// "log"
	// "fmt"
	database "miti-microservices/Database"
// 	"github.com/jinzhu/gorm"
//  _ 	"github.com/jinzhu/gorm/dialects/postgres"
)

type Profile struct{
	UserId string `gorm:"primary_key;type:varchar(100)"  json:"UserId"`
	Name string `gorm:"type:varchar(40)" validate:"required" json:"Name"`
	DateOfBirth string `gorm:"type:varchar(100)" validate:"required" json:"DateOfBirth"`
	Job string `gorm:"type:varchar(30)" validate:"required" "json:"Job"`
	ProfilePicID string `gorm:"type:varchar" json:"ProfilePicId"`
	// ProfilePicURL string `gorm:"type:varchar(100)"  json:"ProfilePicURl"`
	Gender string `gorm:"type:varchar(10)" validate:"required" json:"Gender"`
	Sex string `gorm:"type:varchar(10)"  json:"Sex"`
	RelationshipStatus string `gorm:"type:varchar(10)"  json:"RelationshipStatus"`
	ParentsAddress string `gorm:"type:varchar(10)"  json:"ParentsAddress"`
	Language string `gorm:"type:varchar(20)" validate:"required" json:"Language"`
	Country string `gorm:"type:varchar(30)" json:"Country"`
	Extraversion int `gorm:"type:int" json:"Extraversion"`
	Agreeableness int `gorm:"type:int" json:"Agreeableness"`
	Conscientiousness int `gorm:"type:int" json:"Conscientiousness"`
	EmotionalStability int `gorm:"type:int" json:"EmotionalStability"`
	Intellect int `gorm:"type:int" json:"Intellect"`
	MakingChoice string `gorm:"type:varchar(20)" json:"MakingChoice"`
	Pincode string `gorm:"type:varchar" json:"Pincode"`
	// Latitude string `gorm:"type:varchar(20)" json:"Latitude"`
	// Longitude string `gorm:"type:varchar(20)" json:"Longitude"`
}

type Match struct{
	UserId1 string `gorm:"primary_key;type:varchar" json:"UserId1"`
	UserId2 string `gorm:"primary_key;type:varchar" json:"UserId2"`
	Like1 string `gorm:"primary_key;type:varchar" json:"Like1"`
	Like2 string `gorm:"primary_key;type:varchar" json:"Like2"`
}

type Interest struct{
	UserId string `gorm:"primary_key;type:varchar(100)"  json:"UserId"`
	InterestIndoorPassive1 string `gorm:"type:varchar(100)" json:"InterestIndoorPassive1"`
	InterestIndoorPassive2 string `gorm:"type:varchar(100)" json:"InterestIndoorPassive2"`
	InterestIndoorActive1 string `gorm:"type:varchar(100)" json:"InterestIndoorActive1"`
	InterestIndoorActive2 string `gorm:"type:varchar(100)" json:"InterestIndoorActive2"`	
	InterestOutdoorPassive1 string `gorm:"type:varchar(100)" json:"InterestOutdoorPassive1"`
	InterestOutdoorPassive2 string `gorm:"type:varchar(100)" json:"InterestOutdoorPassive2"`
	InterestOutdoorActive1 string `gorm:"type:varchar(100)" json:"InterestOutdoorActive1"`
	InterestOutdoorActive2 string `gorm:"type:varchar(100)" json:"InterestOutdoorActive2"`
	InterestOthers1 string `gorm:"type:varchar(100)" json:"InterestOthers1"`
	InterestOthers2 string `gorm:"type:varchar(100)" json:"InterestOthers2"`
	InterestIdeology1 string `gorm:"type:varchar(100)" json:"InterestIdeology1"`
	InterestIdeology2 string `gorm:"type:varchar(100)" json:"InterestIdeology2"`
}


type QuestionResponse struct{
	UserId string `gorm:"primary_key;type:varchar(100)"  json:"UserId"`
	QuestionId1 int `gorm:"type:int" json:"QuestionId1"`
	QuestionId2 int `gorm:"type:int" json:"QuestionId2"`
	QuestionId3 int `gorm:"type:int" json:"QuestionId3"`
	QuestionId4 int `gorm:"type:int" json:"QuestionId4"`
	QuestionId5 int `gorm:"type:int" json:"QuestionId5"`
	QuestionId6 int `gorm:"type:int" json:"QuestionId6"`
	QuestionId7 int `gorm:"type:int" json:"QuestionId7"`
	QuestionId8 int `gorm:"type:int" json:"QuestionId8"`
	QuestionId9 int `gorm:"type:int" json:"QuestionId9"`
	QuestionId10 int `gorm:"type:int" json:"QuestionId10"`
	QuestionId11 int `gorm:"type:int" json:"QuestionId11"`
	QuestionId12 int `gorm:"type:int" json:"QuestionId12"`
	QuestionId13 int `gorm:"type:int" json:"QuestionId13"`
	QuestionId14 int `gorm:"type:int" json:"QuestionId14"`
	QuestionId15 int `gorm:"type:int" json:"QuestionId15"`
	QuestionId16 int `gorm:"type:int" json:"QuestionId16"`
	QuestionId17 int `gorm:"type:int" json:"QuestionId17"`
	QuestionId18 int `gorm:"type:int" json:"QuestionId18"`
	QuestionId19 int `gorm:"type:int" json:"QuestionId19"`
	QuestionId20 int `gorm:"type:int" json:"QuestionId20"`
	QuestionId21 int `gorm:"type:int" json:"QuestionId21"`
	QuestionId22 int `gorm:"type:int" json:"QuestionId22"`
	QuestionId23 int `gorm:"type:int" json:"QuestionId23"`
	QuestionId24 int `gorm:"type:int" json:"QuestionId24"`
	QuestionId25 int `gorm:"type:int" json:"QuestionId25"`
	QuestionId26 int `gorm:"type:int" json:"QuestionId26"`
	QuestionId27 int `gorm:"type:int" json:"QuestionId27"`
	QuestionId28 int `gorm:"type:int" json:"QuestionId28"`
	QuestionId29 int `gorm:"type:int" json:"QuestionId29"`
	QuestionId30 int `gorm:"type:int" json:"QuestionId30"`
	QuestionId31 int `gorm:"type:int" json:"QuestionId31"`
	QuestionId32 int `gorm:"type:int" json:"QuestionId32"`
	QuestionId33 int `gorm:"type:int" json:"QuestionId33"`
	QuestionId34 int `gorm:"type:int" json:"QuestionId34"`
	QuestionId35 int `gorm:"type:int" json:"QuestionId35"`
	QuestionId36 int `gorm:"type:int" json:"QuestionId36"`
	QuestionId37 int `gorm:"type:int" json:"QuestionId37"`
	QuestionId38 int `gorm:"type:int" json:"QuestionId38"`
	QuestionId39 int `gorm:"type:int" json:"QuestionId39"`
	QuestionId40 int `gorm:"type:int" json:"QuestionId40"`
	QuestionId41 int `gorm:"type:int" json:"QuestionId41"`
	QuestionId42 int `gorm:"type:int" json:"QuestionId42"`
	QuestionId43 int `gorm:"type:int" json:"QuestionId43"`
	QuestionId44 int `gorm:"type:int" json:"QuestionId44"`
	QuestionId45 int `gorm:"type:int" json:"QuestionId45"`
	QuestionId46 int `gorm:"type:int" json:"QuestionId46"`
	QuestionId47 int `gorm:"type:int" json:"QuestionId47"`
	QuestionId48 int `gorm:"type:int" json:"QuestionId48"`
	QuestionId49 int `gorm:"type:int" json:"QuestionId49"`
	QuestionId50 int `gorm:"type:int" json:"QuestionId50"`
	
}

type Question struct{
	Id int `gorm:"int" json:"Id"`
	Content string `gorm:"varchar(1000)" json:"Content"`
	Type int `gorm:"type:int" json:"Type"`
	Factor int `gorm:"int" json:"Factor"`
}

type Status struct{
	UserId string `gorm:"primary_key;type:varchar"  json:"UserId"`
	StatusId string `gorm:"primary_key";type:varchar json:"StatusId"`
	CreatedAt string `gorm:"type:varchar" json:"CreatedAt"`
	RequestId string `gorm:"type:varchar" json:"RequestId"`
	StatusContent string `gorm:"type:varchar" json:"StatusContent"`
	ActiveStatus string `gorm:"type:varchar" json:"ActiveStatus"`
	//AccessType : Public or Friends or Family or Acquintence
	AccessType string `gorm:"type:varchar" json:"AccessType"`
}

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

type CheckInterest struct{
	UserId1 string `gorm:"primary_key;type:varchar"  json:"UserId1"`
	UserId2 string `gorm:"primary_key;type:varchar"  json:"UserId2"`
	Interest string `gorm:"primary_key;type:varchar" json:"Interest"`
	CreatedAt string `gorm:"type:varchar" json:"CreatedAt"`
}


func init(){
	db:=database.GetDB()
	db.AutoMigrate(&Profile{})
	db.AutoMigrate(&Match{})
	db.AutoMigrate(&QuestionResponse{})	
	db.AutoMigrate(&Question{})
	db.AutoMigrate(&Interest{})
	db.AutoMigrate(&Status{})
	db.AutoMigrate(&PrimaryTrustChain{})
	db.AutoMigrate(&SecondaryTrustChain{})
	db.AutoMigrate(&CheckInterest{})
}
