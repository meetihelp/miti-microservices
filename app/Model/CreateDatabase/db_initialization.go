package CreateDatabase

import (
	"fmt"
	"os"
  	"github.com/jinzhu/gorm"
 _ 	"github.com/jinzhu/gorm/dialects/postgres"
)

// var db *gorm.DB
const (
  dbUser = "postgres"
  dbName = "mytestdb"
  sslmode = "disable"
  password ="peetal"
)

func db_init() *gorm.DB{
	db_credentials:=fmt.Sprintf("host=localhost port=5432 user=postgres dbname=mytestdb password=mypassword")
	conn ,err := gorm.Open("postgres",db_credentials)
	if err!=nil{
		fmt.Println(err.Error())
		os.Exit(3)
	} else{
		fmt.Println("Database Connecected Successfully ....")
	}
	return conn
}


func init(){
	db:=db_init()
	createUserTable(db)
	createProfileTable(db)
	createSessionTable(db)
	createEducationTable(db)
	createVerification_EmailTable(db)
}



/*func main(){
	create_database_init()
}*/
