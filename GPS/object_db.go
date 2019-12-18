package GPS

import(
	database "miti-microservices/Database"
)

type Location struct{
	Latitude string `json:"Latitude"`
	Longitude string `json:"Longitude"`
}

type UserLocation struct{
	UserId string `gorm:"primary_key;type:varchar(100)"  json:"UserId"`
	Longitude string `gorm:"type:varchar" json:"Longitude"`
	Latitude string `gorm:"type:varchar" json:"Latitude"`
	City string `gorm:"type:varchar" json:"City"`
	UpdatedAt string `gorm:"type:varchar" json:"UpdatedAt"`
}

type EventLocation struct{
	EventId string `gorm:"primary_key;type:varchar(100)"  json:"EventId"`
	EventType string `gorm:"type:varchar" json:"EventType"`
	Longitude string `gorm:"type:varchar" json:"Longitude"`
	Latitude string `gorm:"type:varchar" json:"Latitude"`
	City string `gorm:"type:varchar" json:"City"`
}

func init(){
	db:=database.GetDB()
	db.AutoMigrate(&UserLocation{})
	db.AutoMigrate(&EventLocation{})
}