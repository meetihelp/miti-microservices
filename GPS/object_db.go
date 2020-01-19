package GPS

import(
	database "miti-microservices/Database"
)

type Location struct{
	Latitude string `json:"Latitude"`
	Longitude string `json:"Longitude"`
}

type UserCurrentLocation struct{
	UserId string `gorm:"primary_key;type:varchar(100)"  json:"UserId"`
	Longitude string `gorm:"type:varchar" json:"Longitude"`
	Latitude string `gorm:"type:varchar" json:"Latitude"`
	City string `gorm:"type:varchar" json:"City"`
	Pincode string `gorm:"type:varchar" json:"Pincode"`
	UpdatedAt string `gorm:"type:varchar" json:"UpdatedAt"`
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

type LocationMean struct{
	City string `gorm:"type:varchar" json:"City"`
	Latitude string `gorm:"type:varchar" json:"Latitude"`
	Longitude string `gorm:"type:varchar" json:"Longitude"`
}

type CityPincode struct{
	OfficeName string `gorm:"type:varchar" json:"OfficeName"`
	Pincode string `gorm:"type:varchar" json:"Pincode"`
	OfficeType string `gorm:"type:varchar" json:"OfficeType"`
	DeliveryStatus string `gorm:"type:varchar" json:"DeliveryStatus"`
	DivisionName string `gorm:"type:varchar" json:"DivisionName"`
	RegionName string `gorm:"type:varchar" json:"RegionName"`
	CircleName string `gorm:"type:varchar" json:"CircleName"`
	Taluk string `gorm:"type:varchar" json:"Taluk"`
	DistrictName string `gorm:"type:varchar" json:"DistrictName"`
	StateName string `gorm:"type:varchar" json:"StateName"`
	Telephone string `gorm:"type:varchar" json:"Telephone"`
	RelatedSuboffice string `gorm:"type:varchar" json:"RelatedSuboffice"`
	RelatedHeadoffice string `gorm:"type:varchar" json:"RelatedHeadoffice"`
	Latitude string `gorm:"type:varchar" json:"Latitude"`
	Longitude string `gorm:"type:varchar" json:"Longitude"`
}

func init(){
	db:=database.GetDB()
	// db.AutoMigrate(&UserLocation{})
	// db.AutoMigrate(&EventLocation{})
	// db.AutoMigrate(&LocationMean{})
	// db.AutoMigrate(&CityPincode{})
	db.AutoMigrate(&UserCurrentLocation{})

}