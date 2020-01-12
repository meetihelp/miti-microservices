package Image

import(
	// database "miti-microservices/Database"
)

type UserImage struct{
	UserId string `gorm:"primary_key;varchar(100)"  json:"UserId"`
	ImageId string `gorm:"primary_key;type:varchar" json:"ImageId"`
	AccessType string `gorm:"type:varchar" json:"AccessType"`
	ActualFileName string `gorm:"type:varchar" json:"ActualFileName"`
	Size int `gorm:"type:int" json:"Size"`
	Format string `gorm:"type:varchar" json:"Format"`
	Bucket string `gorm:"type:varchar" json:"Bucket"`
	CreatedAt string `gorm:"type:varchar" json:"CreatedAt"`
	Latitude string `gorm:"type:varchar" json:"Latitude"`
	Longitude string `gorm:"type:varchar" json:"Longitude"`
	Dimension string `gorm:"type:varchar" json:"Dimension"`
	RequestId string `gorm:"type:varchar" json:"RequestId"`
	GeneratedName string `gorm:"type:varchar" json:"GeneratedName"`
}

type EventImage struct{
	EventId string `gorm:"primary_key;unique;varchar(100)"  json:"EventId"`
	ImageId string `gorm:"primary_key;type:varchar" json:"ImageId"`
	ImageType string `gorm:"type:varchar" json:"ImageType"`
	ActualFileName string `gorm:"type:varchar" json:"ActualFileName"`
}

func init(){
	// db:=database.GetDB()
	// db.AutoMigrate(&UserImage{})
	// db.AutoMigrate(&EventImage{})
}