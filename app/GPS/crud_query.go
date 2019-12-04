package GPS

import(
	// util "app/Util"
	database "app/Database"
	// "time"
	// "fmt"
)

func GetUserListByLocationDB(location Location,distance float64) ([]string){
	db:=database.GetDB()
	profile:=[]UserLocation{}
	city:=GetCity(location)
	db.Where("city=?",city).Find(&profile)
	var userList []string
	personLocation:=Location{}
	for _,p:=range profile{
		personLocation.Latitude=p.Latitude
		personLocation.Longitude=p.Longitude
		d:=CalculateDistance(location,personLocation)
		if d<distance{
			userList=append(userList,p.UserId)	
		}
	}
	return userList
}

func UpdateUserLocationDB(userId string,location Location){
	db:=database.GetDB()
	city:=GetCity(location)
	userLocation:=UserLocation{}
	db.Where("user_id=?",userId).Find(&userLocation)
	if userLocation.UserId==""{
		userLocation.UserId=userId
		userLocation.Latitude=location.Latitude
		userLocation.Longitude=location.Longitude
		userLocation.City=city
		db.Create(&userLocation)
		return
	}
	db.Model(&userLocation).Where("user_id=?",userId).Updates(UserLocation{Latitude:location.Latitude,
		Longitude:location.Longitude,City:city})
}
