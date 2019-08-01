package CreateDatabase

import (
	// "log"
	// "fmt"
	"github.com/jinzhu/gorm"
 _ 	"github.com/jinzhu/gorm/dialects/postgres"
)

type Education struct{
	User_id string `gorm:"primary_key;type:varchar(100)"  validate:"required"`
	Degree string `gorm:"primary_key; type:varchar(20)" validate:"required"`
	College string `gorm:"type:varchar(20)" validate:"required"`
}

func createEducationTable(db *gorm.DB){
	db.AutoMigrate(&Education{})
}
