package GPS

import(
	util "miti-microservices/Util"
	database "miti-microservices/Database"
	"math"
	"github.com/jinzhu/gorm"
	// "time"
	// "fmt"
)

func GetUserListByLocationDB(location Location,distance float64) ([]string){
	db:=database.GetDB()
	profile:=[]UserLocation{}
	city,_:=GetCity(db,location)
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
	city,_:=GetCity(db,location)
	pincode,_:=GetPincode(db,location,city)
	db.Table("profiles").Where("user_id=?",userId).Update("pincode",pincode)
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
	updatedAt:=util.GetTime()
	db.Model(&userLocation).Where("user_id=?",userId).Updates(UserLocation{Latitude:location.Latitude,
		Longitude:location.Longitude,City:city,UpdatedAt:updatedAt})
}

// func GetEventListByLocationDB(eventType string,location Location,distance float64) ([]string){
// 	db:=database.GetDB()
// 	event:=[]EventLocation{}
// 	city:=GetCity(location)
// 	db.Where("city=? AND event_type=?",city,eventType).Find(&event)
// 	var eventList []string
// 	eventLocation:=Location{}
// 	for _,e:=range event{
// 		eventLocation.Latitude=e.Latitude
// 		eventLocation.Longitude=e.Longitude
// 		d:=CalculateDistance(location,eventLocation)
// 		if d<distance{
// 			eventList=append(eventList,e.EventId)	
// 		}
// 	}
// 	return eventList
// }

// func InsertEventLocation(eventId string,eventType string,latitude string,longitude string) string{
// 	db:=database.GetDB()
// 	eventLocation:=EventLocation{}
// 	eventLocation.EventId=eventId
// 	eventLocation.EventType=eventType
// 	eventLocation.Latitude=latitude
// 	eventLocation.Longitude=longitude
// 	location:=Location{}
// 	location.Latitude=latitude
// 	location.Longitude=longitude
// 	eventLocation.City=GetCity(location)
// 	db.Create(&eventLocation)
// 	return "Ok"
// }

func GetCity(db *gorm.DB,location Location) (string,bool){
	locationCity:=Location{}
	min_distance:=math.MaxFloat64
	min_city:="Could Not Find"

	cityList:=[]LocationMean{}
	err:=db.Find(&cityList).Error
	if(err!=nil){
		return "",true
	}

	for _,locationMean:=range cityList{
		locationCity.Latitude=locationMean.Latitude
		locationCity.Longitude=locationMean.Longitude
		if(locationCity.Latitude=="NA" || locationCity.Longitude=="NA"){
			// dis:=math.MaxFloat64
		}else{
			dis:=CalculateDistance(location,locationCity)
			if(dis<min_distance){
				min_distance=dis
				min_city=locationMean.City
			}	
		}
		
	}
	return min_city,false
}

func GetPincode(db *gorm.DB,location Location,city string) (string,bool){
	cityPincode:=[]CityPincode{}
	err:=db.Table("city_pincodes").Where("region_name=? OR district_name=?",city,city).Find(&cityPincode).Error
	if(err!=nil){
		return "",true
	}
	locationRegion:=Location{}
	min_distance:=math.MaxFloat64
	pincode:="Error"
	for _,c:=range cityPincode{
		locationRegion.Latitude=c.Latitude
		locationRegion.Longitude=c.Longitude
		if(locationRegion.Latitude=="NA" || locationRegion.Longitude=="NA"){
			// dis:=math.MaxFloat64
		}else{
			dis:=CalculateDistance(location,locationRegion)	
			if(dis<min_distance){
				min_distance=dis
				pincode=c.Pincode
			}
		}
		
	}
	return pincode,false
}

func UpdateUserCurrentLocation(db *gorm.DB,userId string,latitude string,longitude string) bool{
	userCurrentLocationTemp:=UserCurrentLocation{}
	userCurrentLocation:=UserCurrentLocation{}
	location:=Location{}
	location.Latitude=latitude
	location.Longitude=longitude

	city,dbError:=GetCity(db,location)
	if(dbError){
		return true
	}
	pincode,dbError:=GetPincode(db,location,city)
	if(dbError){
		return true
	}

	userCurrentLocation.UserId=userId
	userCurrentLocation.Latitude=latitude
	userCurrentLocation.Longitude=longitude
	userCurrentLocation.UpdatedAt=util.GetTime()
	userCurrentLocation.City=city
	userCurrentLocation.Pincode=pincode

	err:=db.Where("user_id=?",userId).Find(&userCurrentLocationTemp).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return true
	}
	if(userCurrentLocationTemp.UserId==""){
		err:=db.Create(&userCurrentLocation).Error
		if(err!=nil){
			return true
		}
	}else{
		err:=db.Save(&userCurrentLocation).Error
		if(err!=nil){
			return true
		}
	}

	return false
}

func GetUserCurrentPincode(db *gorm.DB,userId string) (string,bool){
	userCurrentLocation:=UserCurrentLocation{}
	err:=db.Where("user_id=?",userId).Find(&userCurrentLocation).Error
	if(err!=nil){
		return "",true
	}
	return userCurrentLocation.Pincode,false
}