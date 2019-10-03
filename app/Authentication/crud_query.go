package Authentication
import(
	"fmt"
	"time"
	// "golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
 _ 	"github.com/jinzhu/gorm/dialects/postgres"
   database "app/Database"
   chat "app/Chat"
    util "app/Util"
)

func Enter_Match_user(user_id1 string,user_id2 string){	

	chatID:=util.Generate_token()
	temp_user1:=util.Generate_token()
	temp_user2:=util.Generate_token()

	Enter_Anonymous_User(user_id1,temp_user2,chatID,"one-to-one",1)
	Enter_Anonymous_User(user_id2,temp_user1,chatID,"one-to-one",2)

}

func Enter_Anonymous_User(user_id string,temp_user_id string,chat_id string,chat_type string,user_index int){
	db:=database.GetDB()
	anonymousUser:=AnonymousUser{}
	anonymousUser.User_id=user_id
	anonymousUser.Anonymous_id=temp_user_id
	anonymousUser.Chat_id=chat_id
	// anonymousUser.CreatedAt=time.Now()
	anonymousUser.CreatedAt=util.GetTime()
	anonymousUser.Status="None"

	chatDetail:=chat.ChatDetail{}
	chatDetail.Temp_User_id=temp_user_id
	chatDetail.Actual_User_id=user_id
	chatDetail.Chat_id=chat_id
	chatDetail.Chat_type=chat_type
	chatDetail.CreatedAt=anonymousUser.CreatedAt
	chatDetail.User_index=user_index

	db.Create(&anonymousUser)
	db.Create(&chatDetail)
}

func Enter_user_data(user_data User) (string,int){
	user_data.Password = util.Generate_encrypted_password(user_data.Password)
	
	db:=database.GetDB()
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

func is_user_exist(db *gorm.DB,user_data User) bool{
	temp_user:=User{}
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
func Check_user(user_data User)(string,string){
	db:=database.GetDB()
	email:=user_data.Email
	phone:=user_data.Phone
	password:=user_data.Password

	user:=User{}
	if email!=""{
		db.Where("email=?",email).First(&user)
		// err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
		status:=util.Comapare_password(user.Password,password)
		if !status{
			return "","WRONG_PASSWORD"
		}
		if user.User_id==""{
			return user.User_id,"NO_USER"
		} 
		if user.User_id !="" && user.Status=="U"{
			return user.User_id,"UNVERIFIED"
		}

		return user.User_id,"OK"
	}

	if phone!=""{
		db.Where("phone=?",phone).First(&user)
		// err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
		status:=util.Comapare_password(user.Password,password)
		if !status{
			// fmt.Println(err.Error())
			return "","WRONG_PASSWORD"
		}
		if user.User_id==""{
			return user.User_id,"NO_USER"
		} 
		if user.User_id !="" && user.Status=="U"{
			return user.User_id,"UNVERIFIED"
		}

		return user.User_id,"OK"

	}
	return "","ERROR"
}

func Check_user_by_id(id string,password string) string{
	db:=database.GetDB()
	user:=User{}
	db.Where("user_id=?",id).First(&user)
	status:=util.Comapare_password(user.Password,password)
	if !status{
		// fmt.Println(err.Error())
		return "WRONG_PASSWORD"
	}
	if user.User_id==""{
		return "NO_USER"
	} 
	if user.User_id !="" && user.Status=="U"{
		return "UNVERIFIED"
	}

	return "OK"
}
func IsUserVerified(id string) bool{
	db:=database.GetDB()
	user:=User{}
	db.Where("user_id=?",id).First(&user)
	if user.Status=="U"{
		return false
	} else{
		return true
	}
}

func Get_user_detail(user_id string) (string,string){
	db:=database.GetDB()
	user:=User{}
	db.Where("user_id=?",user_id).First(&user)
	return user.Email , user.Phone
}

func Update_Password(user_id string,new_Password string){
	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(new_Password), bcrypt.DefaultCost)
	// new_Password = string(hashedPassword)
	db:=database.GetDB()
	new_Password = util.Generate_encrypted_password(new_Password)

	user:=User{}
	db.Model(&user).Where("user_id = ?", user_id).Update("password", new_Password)
}

func GetAllUser() ([]string){
	db:=database.GetDB()
	user:=[]User{}
	db.Find(&user)

	UserList:=make([]string,0)
	for _,id := range user{
		UserList=append(UserList,id.User_id)
	}
	return UserList
}

func Verify_OTP(user_id string,otp string) (bool){
	db:=database.GetDB()
	otp_verification:=OTP_verification{}
	db.Where("user_id=? AND verification_otp=?",user_id,otp).First(&otp_verification)
	if otp_verification.User_id==""{
		return false
	}
	return true
}

func Enter_verification_otp(id string,otp string){
	db:=database.GetDB()
	otp_verification:=OTP_verification{}
	otp_verification.User_id=id
	otp_verification.Verification_otp=otp
	// otp_verification.CreatedAt =time.Now()
	otp_verification.CreatedAt=util.GetTime()
	db.Create(&otp_verification)
}

func Get_otp_verification_count(id string)(int,string){
	count:=0
	otp_verification:=OTP_verification{}
	db:=database.GetDB()
	db.Where("user_id=?",id).Find(&otp_verification).Count(&count)
	// return len(otp_verification),otp_verification.CreatedAt
	return count,otp_verification.CreatedAt
}
func Change_Verification_Status(user_id string){
	db:=database.GetDB()
	user:=User{}
	db.Model(&user).Where("user_id=?",user_id).Update("status","V")
}
func Enter_email_verification(id string,token string){
	db:=database.GetDB()
	email_verification:=Email_verification{}
	email_verification.User_id=id
	email_verification.Verification_token=token
	email_verification.CreatedAt=time.Now()
	db.Create(&email_verification)
}

func Get_Email_verification_count(id string)(int,time.Time){
	count:=0
	email_verification:=Email_verification{}
	db:=database.GetDB()
	db.Where("user_id=?",id).Find(&email_verification).Count(&count)
	fmt.Println(email_verification)
	// return len(email_verification),email_verification.CreatedAt
	fmt.Println(count)
	return count,email_verification.CreatedAt

}


func Delete_all_email_verification(id string){
	db:=database.GetDB()
	db.Where("user_id=?",id).Delete(&Email_verification{})
}

func Verify_Email(token string) (string,bool){
	db:=database.GetDB()
	email_verification:=Email_verification{}
	db.Where("verification_token=?",token).First(&email_verification)
	if email_verification.User_id==""{
		return "",false
	}
	return email_verification.User_id,true
}