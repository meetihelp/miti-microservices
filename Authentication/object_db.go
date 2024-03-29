package Authentication
import (
	// database "miti-microservices/Database"
)


type User struct{
	UserId string `gorm:"primary_key;unique;varchar(100)"  json:"UserId"`
	Phone string `gorm:"type:varchar(13)" validate:"omitempty" json:"Phone"`
	Email string  `gorm:"type:varchar(30)" validate:"omitempty,email" json:"Email"`
	// Password string `gorm:"type:varchar(100)" validate:"required" json:"Password"`
	Status string `gorm:"type:varchar(3)" json:"Status"`  //Verified/Unverified/Deleted
	ProfileCreationStatus string `gorm:"type:varchar(3)" json:"ProfileCreationStatus"`
	PreferenceCreationStatus int `gorm:"type:int" json:"PreferenceCreationStatus"`
	IPIPStatus int `gorm:"type:int" json:"IPIPStatus"`
	CreatedAt string `gorm:"type:Varchar(100)" json:"CreatedAt"`
}
type OTPVerification struct{
	UserId string `gorm:"primary_key;type:varchar(100)"  validate:"required" json:"UserId"`
	SessionId string `gorm:"primary_key;type:varchar(100)"  validate:"required" json:"SessionId"`
	OTP string `gorm:"varchar(6)" validate:"required"  json:"OTP"`
	CreatedAt string `gorm:"type:varchar" json:"CreatedAt"`
	ResendCount int `gorm:"type:int" json:"ResendCount"`
	FailCount int `gorm:"type:int" json:"FailCount"`
	DeliverCount int `gorm:"type:int" json:"DeliverCount"`
}

type EmailVerification struct{
	UserId string `gorm:"primary_key;type:varchar(100)"  validate:"required" json:"UserId"`
	VerificationToken string `gorm:"primary_key;varchar(100)" validate:"required" json:"VerificationToken"`
	CreatedAt string `gorm:"type:varchar(100)" json:"CreatedAt"`
}


// type AnonymousUser struct{
// 	UserId string `gorm:"primary_key;varchar(100)"  json:"UserId"`
// 	AnonymousId string `gorm:"primary_key;unique;varchar(100)"  json:"AnonymousId"`
// 	ChatId string `gorm:"primary_key;varchar(100)"  json:"ChatId"`
// 	Status string `gorm:"type:varchar(6)" json:"Status"`  // status for Liked/not liked/ none
// 	CreatedAt string `gorm:"type:varchar(100)" json:"CreatedAt"`
// }

type ForgetPasswordStatus struct{
	SessionId string `gorm:"primary_key;type:varchar(100)"  validate:"required" json:"SessionId"`
	VerificationStatus string `gorm:"type:varchar" json:"VerificationStatus"`
}
func init(){	
	// db:=database.GetDB()
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&OTPVerification{})
	// db.AutoMigrate(&EmailVerification{})
	// // db.AutoMigrate(&AnonymousUser{})
	// db.AutoMigrate(&ForgetPasswordStatus{})
}
