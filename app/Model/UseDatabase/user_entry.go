package UseDatabase

import(
	// "fmt"
	"time"
	// "golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
 _ 	"github.com/jinzhu/gorm/dialects/postgres"
   CD "app/Model/CreateDatabase"
    util "app/Utility"
)

func Enter_user_data(user_data CD.User) (string,int){
	user_data.Password = util.Generate_encrypted_password(user_data.Password)
	
	db:=GetDB()
	//CHECK IF USER EMAIL ID OR PHONE ALREADY EXISTS
	checking_status:=is_user_exist(db,user_data)
	if checking_status == true{
		return "",1
	}
	//GENERATE USER ID
	user_data.User_id =util.Generate_token()
	user_data.Status="U"
	user_data.CreatedAt =time.Now()
	//INSERT IN DATABASE
	db.Create(&user_data)
	return user_data.User_id,2
}

func is_user_exist(db *gorm.DB,user_data CD.User) bool{
	temp_user:=CD.User{}
	if user_data.Phone!=""{
		db.Where("phone=?",user_data.Phone).First(&temp_user)
		if temp_user.User_id!=""{
			return true
		}
	}

	if user_data.Email!=""{
		db.Where("email=?",user_data.Email).First(&temp_user)
		if temp_user.User_id!=""{
			return true
		}
	}

	return false
}
