package Authentication
import (
	// "log"
	// "fmt"
	database "app/Database"
	"time"
	// "github.com/jinzhu/gorm"
 // _ 	"github.com/jinzhu/gorm/dialects/postgres"
)


type User struct{
	User_id string `gorm:"primary_key;unique;varchar(100)"  json:"user_id"`
	Phone string `gorm:"type:varchar(13)" validate:"omitempty" json:"phone"`
	Email string  `gorm:"type:varchar(30)" validate:"omitempty,email" json:"email"`
	Password string `gorm:"type:varchar(100)" validate:"required" json:"password"`
	Status string `gorm:"type:varchar(3)" json:"Status"`
	CreatedAt time.Time `gorm:"type:time" json:"created_at"`
}
type OTP_verification struct{
	User_id string `gorm:"primary_key;type:varchar(100)"  validate:"required"`
	Verification_otp string `gorm:"primary_key;varchar(100)" validate:"required"`
	CreatedAt string `gorm:"type:varchar"`
}

type Email_verification struct{
	User_id string `gorm:"primary_key;type:varchar(100)"  validate:"required"`
	Verification_token string `gorm:"primary_key;varchar(100)" validate:"required"`
	CreatedAt time.Time `gorm:"type:time"`
}


type AnonymousUser struct{
	User_id string `gorm:"primary_key;varchar(100)"  json:"user_id"`
	Anonymous_id string `gorm:"primary_key;unique;varchar(100)"  json:"anonymous_user_id"`
	Chat_id string `gorm:"primary_key;varchar(100)"  json:"chat_id"`
	Status string `gorm:"type:varchar(6)" json:"status"`
	CreatedAt string `gorm:"type:varchar" json:"created_at"`
}

func init(){	
	db:=database.GetDB()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&OTP_verification{})
	db.AutoMigrate(&Email_verification{})
	db.AutoMigrate(&AnonymousUser{})
}
