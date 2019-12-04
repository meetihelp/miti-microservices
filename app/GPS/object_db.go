package GPS

import(
	database "app/Database"
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
}

func init(){
	db:=database.GetDB()
	db.AutoMigrate(&UserLocation{})
}