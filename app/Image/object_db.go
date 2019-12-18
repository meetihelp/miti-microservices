package Image

import(
	database "app/Database"
)

type UserImage struct{
	UserId string `gorm:"primary_key;unique;varchar(100)"  json:"UserId"`
	ImageId string `gorm:"primary_key;type:varchar" json:"ImageId"`
}

type EventImage struct{
	EventId string `gorm:"primary_key;unique;varchar(100)"  json:"EventId"`
	ImageId string `gorm:"primary_key;type:varchar" json:"ImageId"`
}

func init(){
	db:=database.GetDB()
	db.AutoMigrate(&UserImage{})
	db.AutoMigrate(&EventImage{})
}