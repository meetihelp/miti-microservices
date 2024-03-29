package Database

import (
	"fmt"
	"os"
  	"github.com/jinzhu/gorm"
 _ 	"github.com/jinzhu/gorm/dialects/postgres"
)




// func getPostgresPassword() string{
// 	return os.Getenv("postgres_password")
// }
func DBConnection() *gorm.DB{
	// db_credentials:=fmt.Sprintf("host=localhost port=5432 user=postgres dbname=mytestdb password=mypass")
	postgres_password:=getPostgresPassword()
	// db_credentials:=fmt.Sprintf("sslmode=require user=doadmin host=db-postgresql-blr1-60587-do-user-3919569-0.db.ondigitalocean.com dbname=defaultdb port=25060  sslmode=require sslrootcert=/home/gaurav/Downloads/do/ca-certificate.crt password="+postgres_password)
	db_credentials:=fmt.Sprintf("sslmode=require user=doadmin host=db-postgresql-blr1-60587-do-user-3919569-0.db.ondigitalocean.com dbname=production port=25061  sslmode=require sslrootcert=/home/gaurav/Downloads/do/ca-certificate.crt password="+postgres_password)
	conn ,err := gorm.Open("postgres",db_credentials)
	if err!=nil{
		fmt.Println(err.Error())
		fmt.Println("gaurav")
		os.Exit(3)
	} else{
		fmt.Println("Database Connecected Successfully ....")
	}
	return conn
}


