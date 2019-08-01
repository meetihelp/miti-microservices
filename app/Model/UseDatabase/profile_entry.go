package UseDatabase

import(
	"fmt"
	// "golang.org/x/crypto/bcrypt"
	// "github.com/jinzhu/gorm"
 // _ 	"github.com/jinzhu/gorm/dialects/postgres"
   CD "app/Model/CreateDatabase"
    // ut "app/Utility"
)

func Enter_profile_data(profile_data CD.Profile){
	fmt.Println("Enter_profile_data")
	db:=GetDB()
	//INSERT IN DATABASE
	db.Create(&profile_data)
}