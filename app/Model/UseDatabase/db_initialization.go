package UseDatabase

import (
	"fmt"
	"os"
  	"github.com/jinzhu/gorm"
 _ 	"github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
const (
  dbUser = "postgres"
  dbName = "mytestdb"
  sslmode = "disable"
  password ="peetal"
)

func init() {
	db_credentials:=fmt.Sprintf("host=localhost port=5432 user=postgres dbname=mytestdb password=mypassword")
	conn ,err := gorm.Open("postgres",db_credentials)
	if err!=nil{
		fmt.Println(err.Error())
		os.Exit(3)
	} else{
		fmt.Println("Database Connecected Successfully ....")
	}
	db= conn
}

func GetDB() *gorm.DB{
	return db
}
