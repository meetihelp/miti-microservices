package Database

import (
	"fmt"
	"os"
  	"github.com/jinzhu/gorm"
 _ 	"github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB


func init() {
	db_credentials:=fmt.Sprintf("host=localhost port=5432 user=postgres dbname=mytestdb password=mypass")
	conn ,err := gorm.Open("postgres",db_credentials)
	if err!=nil{
		fmt.Println(err.Error())
		fmt.Println("gaurav")
		os.Exit(3)
	} else{
		fmt.Println("Database Connecected Successfully ....")
	}
	db= conn
}

func GetDB() *gorm.DB{
	return db
}
