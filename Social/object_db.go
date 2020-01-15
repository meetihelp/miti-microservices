package Social

import(
	database "miti-microservices/Database"
)
type Pool struct{
	UserId string `gorm:"primary_key;type:varchar" json:"UserId"`
	Pincode string `gorm:"type:varchar" json:"Pincode"`
	AreaCode string `gorm:"type:varchar" json:"AreaCode"`
	CreatedAt string `gorm:"type:varchar" json:"CreatedAt"`
	Gender string `gorm:"type:varchar" json:"Gender"`
	Sex string `gorm:"type:varchar" json:"Sex"`
}

type PoolWaiting struct{
	UserId string `gorm:"primary_key;type:varchar" json:"UserId"`
	Pincode string `gorm:"type:varchar" json:"Pincode"`
	AreaCode string `gorm:"type:varchar" json:"AreaCode"`
	CreatedAt string `gorm:"type:varchar" json:"CreatedAt"`
	Gender string `gorm:"type:varchar" json:"Gender"`
	Sex string `gorm:"type:varchar" json:"Sex"` 
	RequestId string `gorm:"varchar" json:"RequestId"`
}


type UserPool struct{
	UserId string `gorm:"primary_key;type:varchar" json:"UserId"`
	PoolSuccessCount int64 `gorm:"type:int" json:"PoolSuccessCount"`
	UserLikeCount int64 `gorm:"type:int" json:"UserLikeCount"`
	OtherLikeCount int64 `gorm:"type:int" json:"OtherLikeCount"`
	CurrentStatus int `gorm:"type:int" json:"CurrentStatus"`
	//CurrentStatus:->Never In Pool=0,First Time in Pool=1,Pooled but did not like=2,
	//Pooled but was not liked=3,Pooled and matched=4
	MatchCount int64 `gorm:"type:int" json:"MatchCount"`
}

type PoolStatus struct{
	UserId string `gorm:"primary_key;type:varchar" json:"UserId"`
	ChatId string `gorm:"type:varchar" json:"ChatId"`
	MatchUserId string `gorm:type:varchar" json:"MatchUserId"`
	Status string `gorm:"type:varchar" json:"Status"`
	CreatedAt string `gorm:"type:varchar" json:"CreatedAt"`
	Like1Time string `gorm:"type:varchar" json:"Like1Time"`
	Like2TIme string `gorm:"type:varchar" json:"Like2TIme"`
	MatchTime string `gorm:"type:varchar" json:"MatchTime"`
}


type PoolLog struct{
	UserId1 string `gorm:"primary_key;type:varchar" json:"UserId1"`
	UserId2 string `gorm:"primary_key;type:varchar" json:"UserId2"`
	CreatedAt string `gorm:"type:varchar" json:"CreatedAt"`
	Like1 string `gorm:"type:varchar" json:"Like1"`
	Like2 string `gorm:"type:varchar" json:"Like2"`
	Like1Time string `gorm:"type:varchar" json:"Like1Time"`
	Like2TIme string `gorm:"type:varchar" json:"Like2TIme"`
	MatchTime string `gorm:"type:varchar" json:"MatchTime"`
}


type GroupPool struct{
	UserId string `gorm:"primary_key;type:varchar" json:"UserId"`
	Pincode string `gorm:"type:varchar" json:"Pincode"`
	Interest string `gorm:"type:varchar" json:"Interest"`
	AreaCode string `gorm:"type:varchar" json:"AreaCode"`
	CreatedAt string `gorm:"type:varchar" json:"CreatedAt"`
	Gender string `gorm:"type:varchar" json:"Gender"`
	Sex string `gorm:"type:varchar" json:"Sex"`
}

type GroupPoolWaiting struct{
	UserId string `gorm:"primary_key;type:varchar" json:"UserId"`
	Pincode string `gorm:"type:varchar" json:"Pincode"`
	Interest string `gorm:"type:varchar" json:"Interest"`
	AreaCode string `gorm:"type:varchar" json:"AreaCode"`
	CreatedAt string `gorm:"type:varchar" json:"CreatedAt"`
	Gender string `gorm:"type:varchar" json:"Gender"`
	Sex string `gorm:"type:varchar" json:"Sex"` 
}


type GroupUserPool struct{
	UserId string `gorm:"primary_key;type:varchar" json:"UserId"`
	PoolSuccessCount int64 `gorm:"type:int" json:"PoolSuccessCount"`
	UserLikeCount int64 `gorm:"type:int" json:"UserLikeCount"`
	OtherLikeCount int64 `gorm:"type:int" json:"OtherLikeCount"`
	CurrentStatus int `gorm:"type:int" json:"CurrentStatus"`
	//CurrentStatus:->Never In Pool=0,First Time in Pool=1,Pooled but did not like=2,
	//Pooled but was not liked=3,Pooled and matched=4
	MatchCount int64 `gorm:"type:int" json:"MatchCount"`
}

type GroupPoolStatus struct{
	UserId string `gorm:"primary_key;type:varchar" json:"UserId"`
	MatchUserId string `gorm:type:varchar" json:"MatchUserId"`
	Status string `gorm:"type:varchar" json:"Status"`
	CreatedAt string `gorm:"type:varchar" json:"CreatedAt"`
	Like1Time string `gorm:"type:varchar" json:"Like1Time"`
	Like2TIme string `gorm:"type:varchar" json:"Like2TIme"`
	MatchTime string `gorm:"type:varchar" json:"MatchTime"`
}
func init(){
	db:=database.GetDB()
	db.AutoMigrate(&Pool{})
	db.AutoMigrate(&PoolWaiting{})
	db.AutoMigrate(&UserPool{})
	db.AutoMigrate(&PoolStatus{})
	db.AutoMigrate(&PoolLog{})
}