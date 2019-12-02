package Authentication
import(
	"fmt"
	"github.com/jinzhu/gorm"
 _ 	"github.com/jinzhu/gorm/dialects/postgres"
   	database "app/Database"
   	chat "app/Chat"
    util "app/Util"
)

func EnterMatchUser(userId1 string,userId2 string){	

	chatID:=util.GenerateToken()
	tempUser1:=util.GenerateToken()
	tempUser2:=util.GenerateToken()

	EnterAnonymousUser(userId1,tempUser2,chatID,"OneToOne",1)
	EnterAnonymousUser(userId2,tempUser1,chatID,"OneToOne",2)

}

func EnterAnonymousUser(userId string,tempUserId string,chatId string,chatType string,userIndex int){
	db:=database.GetDB()
	anonymousUser:=AnonymousUser{}
	anonymousUser.UserId=userId
	anonymousUser.AnonymousId=tempUserId
	anonymousUser.ChatId=chatId
	anonymousUser.CreatedAt=util.GetTime()
	anonymousUser.Status="None"

	chatDetail:=chat.ChatDetail{}
	chatDetail.TempUserId=tempUserId
	chatDetail.ActualUserId=userId
	chatDetail.ChatId=chatId
	chatDetail.ChatType=chatType
	chatDetail.CreatedAt=anonymousUser.CreatedAt
	chatDetail.UserIndex=userIndex

	db.Create(&anonymousUser)
	db.Create(&chatDetail)
}

func EnterUserData(userData User) (string,int){
	userData.Password = util.GenerateEncryptedPassword(userData.Password)
	
	db:=database.GetDB()
	//CHECK IF USER EMAIL ID OR PHONE ALREADY EXISTS
	checkingStatus:=IsUserExist(db,userData)
	if checkingStatus == true{
		return "",1
	}
	//GENERATE USER ID
	userData.UserId =util.GenerateToken()
	userData.Status="U"
	// userData.CreatedAt =time.Now()
	userData.CreatedAt =util.GetTime()
	//INSERT IN DATABASE
	db.Create(&userData)
	return userData.UserId,2
}

func IsUserExist(db *gorm.DB,userData User) bool{
	tempUser:=User{}
	if userData.Phone!=""{
		db.Where("phone=?",userData.Phone).First(&tempUser)
		if tempUser.UserId!=""{
			return true
		}
	}

	if userData.Email!=""{
		db.Where("email=?",userData.Email).First(&tempUser)
		if tempUser.UserId!=""{
			return true
		}
	}

	return false
}
func CheckUser(userData User)(string,string){
	db:=database.GetDB()
	email:=userData.Email
	phone:=userData.Phone
	password:=userData.Password

	user:=User{}
	if email!=""{
		db.Where("email=?",email).First(&user)
		status:=util.ComaparePassword(user.Password,password)
		if !status{
			return "","WrongPassword"
		}
		if user.UserId==""{
			return user.UserId,"NoUser"
		} 
		if user.UserId !="" && user.Status=="U"{
			return user.UserId,"Unverified"
		}

		return user.UserId,"Ok"
	}

	if phone!=""{
		db.Where("phone=?",phone).First(&user)
		status:=util.ComaparePassword(user.Password,password)
		if !status{
			return "","WrongPassword"
		}
		if user.UserId==""{
			return user.UserId,"NoUser"
		} 
		if user.UserId !="" && user.Status=="U"{
			return user.UserId,"Unverified"
		}

		return user.UserId,"Ok"

	}
	return "","Error"
}

func CheckUserById(id string,password string) string{
	db:=database.GetDB()
	user:=User{}
	db.Where("userId=?",id).First(&user)
	status:=util.ComaparePassword(user.Password,password)
	if !status{
		return "WrongPassword"
	}
	if user.UserId==""{
		return "NoUser"
	} 
	if user.UserId !="" && user.Status=="U"{
		return "Unverified"
	}

	return "Ok"
}
func IsUserVerified(id string) bool{
	db:=database.GetDB()
	user:=User{}
	db.Where("userId=?",id).First(&user)
	if user.Status=="U"{
		return false
	} else{
		return true
	}
}

func GetUserDetail(userId string) (string,string){
	db:=database.GetDB()
	user:=User{}
	db.Where("userId=?",userId).First(&user)
	return user.Email , user.Phone
}

func UpdatePasswordFunc(userId string,newPassword string){
	db:=database.GetDB()
	newPassword = util.GenerateEncryptedPassword(newPassword)

	user:=User{}
	db.Model(&user).Where("userId = ?", userId).Update("password", newPassword)
}

func GetAllUser() ([]string){
	db:=database.GetDB()
	user:=[]User{}
	db.Find(&user)

	UserList:=make([]string,0)
	for _,id := range user{
		UserList=append(UserList,id.UserId)
	}
	return UserList
}

func VerifyOTP(userId string,otp string) (bool){
	db:=database.GetDB()
	otpVerification:=OTPVerification{}
	db.Where("userId=? AND verification_otp=?",userId,otp).First(&otpVerification)
	if otpVerification.UserId==""{
		return false
	}
	return true
}

func EnterVerificationOtp(id string,otp string){
	db:=database.GetDB()
	otpVerification:=OTPVerification{}
	otpVerification.UserId=id
	otpVerification.VerificationOtp=otp
	otpVerification.CreatedAt=util.GetTime()
	db.Create(&otpVerification)
}

func GetOtpVerificationCount(id string)(int,string){
	count:=0
	otpVerification:=OTPVerification{}
	db:=database.GetDB()
	db.Where("userId=?",id).Find(&otpVerification).Count(&count)
	return count,otpVerification.CreatedAt
}
func ChangeVerificationStatus(userId string){
	db:=database.GetDB()
	user:=User{}
	db.Model(&user).Where("userId=?",userId).Update("status","V")
}
func EnterEmailVerification(id string,token string){
	db:=database.GetDB()
	emailVerification:=EmailVerification{}
	emailVerification.UserId=id
	emailVerification.VerificationToken=token
	emailVerification.CreatedAt=util.GetTime()
	db.Create(&emailVerification)
}


func GetEmailVerificationCount(id string)(int,string){
	count:=0
	emailVerification:=EmailVerification{}
	db:=database.GetDB()
	db.Where("userId=?",id).Find(&emailVerification).Count(&count)
	fmt.Println(count)
	return count,emailVerification.CreatedAt

}


func DeleteAllEmailVerification(id string){
	db:=database.GetDB()
	db.Where("userId=?",id).Delete(&EmailVerification{})
}

func VerifyEmailFunc(token string) (string,bool){
	db:=database.GetDB()
	emailVerification:=EmailVerification{}
	db.Where("verification_token=?",token).First(&emailVerification)
	if emailVerification.UserId==""{
		return "",false
	}
	return emailVerification.UserId,true
}