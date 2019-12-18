package Event

import(
	database "app/Database"
)

type Event struct{
	EventId string `gorm:"primary_key;type:varchar" json:"EventId"`
	EventName string `gorm:"type:varchar" json:"EventName"`
	EventPicURL string `gorm:"type:varchar" json:"EventPicURL"`
	EventType string `gorm:"type:varchar" json:"EventType"`
	Time string `gorm:"type:varchar" json:"Time"`
	OrganiserId string `gorm:"type:varchar" json:"OrganiserId"`
	Latitude string `gorm:"type:varchar" json:"Latitude"`
	Longitude string `gorm:"type:varchar" json:"Longitude"`
	CreatedAt string `gorm:"type:varchar" json:"CreatedAt'`
}

func init(){
	db:=database.GetDB()
	db.AutoMigrate(&Event{})
}