package CreateDatabase

import (
	// "log"
	// "fmt"
	"time"
	"github.com/jinzhu/gorm"
 _ 	"github.com/jinzhu/gorm/dialects/postgres"
)


type User struct{
	User_id string `gorm:"primary_key;unique;varchar(100)"  json:"user_id"`
	Phone string `gorm:"type:varchar(13)" validate:"omitempty" json:"phone"`
	Email string  `gorm:"type:varchar(30)" validate:"omitempty,email" json:"email"`
	Password string `gorm:"type:varchar(100)" validate:"required" json:"password"`
	Status string `gorm:"type:varchar(3)" json:"Status"`
	CreatedAt time.Time `gorm:"type:time" json:"created_at"`
}
func createUserTable(db *gorm.DB){	
	db.AutoMigrate(&User{})
}
