package Authentication
import(
	"fmt"
	"github.com/jinzhu/gorm"
 _ 	"github.com/jinzhu/gorm/dialects/postgres"
   	database "miti-microservices/Database"
    util "miti-microservices/Util"
)

func LoadingPageQuery(db *gorm.DB,id string) (bool,bool,int,bool){
	user:=User{}
	IsUserVerified:=false
	IsProfileCreated:=false
	Preferece:=0
	err:=db.Where("user_id=?",id).First(&user).Error
	if(gorm.IsRecordNotFoundError(err)){
		return IsUserVerified,IsProfileCreated,Preferece,false
	}
	if(err!=nil){
		return IsUserVerified,IsProfileCreated,Preferece,true	
	}
	if user.Status=="U"{
		IsUserVerified=false
	} else{
		IsUserVerified=true
	}

	if user.ProfileCreationStatus=="Y"{
		IsProfileCreated=true
	} else{
		IsProfileCreated=false
	}

	Preferece=user.PreferenceCreationStatus
	return IsUserVerified,IsProfileCreated,Preferece,false
}

func IsUserExist(db *gorm.DB,userData User) (bool,bool){
	tempUser:=User{}
	if userData.Phone!=""{
		db.Where("phone=?",userData.Phone).First(&tempUser)
		if tempUser.UserId!=""{
			return true,false
		}
	}

	if userData.Email!=""{
		db.Where("email=?",userData.Email).First(&tempUser)
		if tempUser.UserId!=""{
			return true,false
		}
	}

	return false,false
}

func EnterUserData(db *gorm.DB,userData LoginRequest) (string,bool){
	// userData.Password = util.GenerateEncryptedPassword(userData.Password)
	user:=User{}	
	err:=db.Where("phone=?",userData.Phone).First(&user).Error
	if(gorm.IsRecordNotFoundError(err)){
		user.UserId =util.GenerateToken()
		user.Status="U"
		user.ProfileCreationStatus="N"
		user.IPIPStatus=0
		user.PreferenceCreationStatus=0
		user.CreatedAt =util.GetTime()
		err=db.Create(&user).Error
		if(err!=nil){
			fmt.Println(err)
			return "",true
		}
		return user.UserId,false
	}
	if(err!=nil){
		fmt.Println(err)
		return "",true
	}
	return user.UserId,false
}

func InsertOTP(db *gorm.DB,userId string,sessionId string) (string,bool){
	otp:=util.GenerateOTP()
	otpVerification:=OTPVerification{}
	err:=db.Where("user_id=?",userId).Find(&otpVerification).Error
	if(otpVerification.UserId!=""){
		resendCount:=otpVerification.ResendCount+1
		err=db.Model(&otpVerification).Where("user_id = ?", userId).Update("resend_count", resendCount).Error
		if(err!=nil){
			return "",true	
		}
		return otp,false
	}
	if(err!=nil){
		return "",true
	}
	otpVerification.UserId=userId
	otpVerification.SessionId=sessionId
	otpVerification.OTP=otp
	otpVerification.ResendCount=0
	otpVerification.DeliverCount=0
	otpVerification.FailCount=0
	otpVerification.CreatedAt=util.GetTime()
	err=db.Create(&otpVerification).Error
	if(err!=nil){
		return "",true
	}
	return otp,false
}

func GetPhoneFromUserId(db *gorm.DB,userId string) (string,string,bool){
	// db:=database.GetDB()
	user:=User{}
	err:=db.Where("user_id=?",userId).Find(&user).Error
	if(err!=nil){
		return "","Error",true
	}
	if user.Phone!=""{
		return user.Phone,"Ok",false
	}
	return "","Error",false
}


func GetOTPDetails(db *gorm.DB,userId string) (OTPVerification,bool){
	otp:=OTPVerification{}
	err:=db.Where("user_id=?",userId).Find(&otp).Error
	if(err!=nil){
		return otp,true
	}
	return otp,false
}

func DeleteOTP(db *gorm.DB,id string) bool{
	err:=db.Where("user_id=?",id).Delete(&OTPVerification{}).Error
	if(err!=nil){
		return true
	}
	return false
}


func VerifyOTPDB(db *gorm.DB,userId string,otp string) (bool,int,bool){
	// db:=database.GetDB()
	otpVerification:=OTPVerification{}
	err:=db.Where("user_id=?",userId).First(&otpVerification).Error
	if(err!=nil){
		return false,otpVerification.FailCount,true
	}
	if(otpVerification.UserId!="" && otpVerification.OTP==otp){
		return true,otpVerification.FailCount,false
	}
	return false,otpVerification.FailCount,false
}


func ChangeVerificationStatus(db *gorm.DB,userId string)(bool){
	user:=User{}
	err:=db.Model(&user).Where("user_id=?",userId).Update("status","V").Error
	if(err!=nil){
		return true
	}

	return false
}


func UpdateFailCount(db *gorm.DB,userId string,failCount int) bool{
	// db:=database.GetDB()
	otp:=OTPVerification{}
	err:=db.Model(&otp).Where("user_id=?",userId).Update("fail_count",failCount+1).Error
	if(err!=nil){
		return true
	}
	return false
}

func UpdateProfileCreationStatus(db *gorm.DB,userId string) bool{
	user:=User{}
	err:=db.Model(&user).Where("user_id = ?", userId).Update("profile_creation_status", "Y").Error
	if(err!=nil){
		return true
	}
	return false
}

func UpdateIPIPStatus(db *gorm.DB,userId string,ipipStatus int) bool{
	user:=User{}
	err:=db.Where("user_id=?",userId).Find(&user).Error
	if(err!=nil){
		return true
	}
	if(ipipStatus>user.IPIPStatus){
		err=db.Model(&user).Where("user_id=?",userId).Update("ip_ip_status",ipipStatus).Error
		if(err!=nil){
			return true
		}
	}
	return false
	
}

func UpdatePreferencetatus(db *gorm.DB,userId string,preferenceStatus int) bool{
	user:=User{}
	err:=db.Model(&user).Where("user_id=?",userId).Update("preference_creation_status",preferenceStatus).Error
	if(err!=nil){
		return true
	}
	return false
}

//Last








// func EnterMatchUser(userId1 string,userId2 string){	

// 	chatID:=util.GenerateToken()
// 	tempUser1:=util.GenerateToken()
// 	tempUser2:=util.GenerateToken()

// 	EnterAnonymousUser(userId1,tempUser2,chatID,"OneToOne",1)
// 	EnterAnonymousUser(userId2,tempUser1,chatID,"OneToOne",2)

// 	EnterMatchData(userId1,userId2)
// }

// func EnterMatchData(userId1 string,userId2 string){
// 	db:=database.GetDB()
// 	match:=util.Match{}
// 	match.UserId1=userId1
// 	match.UserId2=userId2
// 	db.Create(&match)
// }
// func EnterAnonymousUser(userId string,tempUserId string,chatId string,chatType string,userIndex int){
// 	db:=database.GetDB()
// 	anonymousUser:=AnonymousUser{}
// 	anonymousUser.UserId=userId
// 	anonymousUser.AnonymousId=tempUserId
// 	anonymousUser.ChatId=chatId
// 	anonymousUser.CreatedAt=util.GetTime()
// 	anonymousUser.Status="None"

// 	chatDetail:=chat.ChatDetail{}
// 	chatDetail.TempUserId=tempUserId
// 	chatDetail.ActualUserId=userId
// 	chatDetail.ChatId=chatId
// 	chatDetail.ChatType=chatType
// 	chatDetail.CreatedAt=anonymousUser.CreatedAt
// 	chatDetail.UserIndex=userIndex

// 	db.Create(&anonymousUser)
// 	db.Create(&chatDetail)


// }




func CheckUserCredentials(userData User)(string,string){
	db:=database.GetDB()
	email:=userData.Email
	phone:=userData.Phone
	// password:=userData.Password

	user:=User{}
	if email!=""{
		db.Where("email=?",email).First(&user)
		// status:=util.ComaparePassword(user.Password,password)
		// if !status{
		// 	return "","WrongPassword"
		// }
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
		// status:=util.ComaparePassword(user.Password,password)
		if user.UserId==""{
			return user.UserId,"NoUser"
		}
		// if !status{
		// 	return "","WrongPassword"
		// }
	 
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
	db.Where("user_id=?",id).First(&user)
	// status:=util.ComaparePassword(user.Password,password)
	// if !status{
	// 	return "WrongPassword"
	// }
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
	db.Where("user_id=?",id).First(&user)
	if user.Status=="U"{
		return false
	} else{
		return true
	}
}

func GetUserDetail(userId string) (string,string){
	db:=database.GetDB()
	user:=User{}
	db.Where("user_id=?",userId).First(&user)
	return user.Email , user.Phone
}

func UpdatePasswordFunc(userId string,newPassword string){
	db:=database.GetDB()
	newPassword = util.GenerateEncryptedPassword(newPassword)

	user:=User{}
	db.Model(&user).Where("user_id = ?", userId).Update("password", newPassword)
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



func EnterVerificationOtp(id string,otp string){
	db:=database.GetDB()
	otpVerification:=OTPVerification{}
	otpVerification.UserId=id
	otpVerification.OTP=otp
	otpVerification.CreatedAt=util.GetTime()
	db.Create(&otpVerification)
}



func GetOtpVerificationCount(id string)(int,string){
	count:=0
	otpVerification:=[]OTPVerification{}
	db:=database.GetDB()
	db.Where("user_id=?",id).Find(&otpVerification).Count(&count)
	return count,otpVerification[count-1].CreatedAt
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
	db.Where("userId=?",id).Order("created_at").Find(&emailVerification).Count(&count)
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



func GetUserIdFromPhone(db *gorm.DB,phone string) (string,string){
	// db:=database.GetDB()
	user:=User{}
	db.Where("phone=?",phone).Find(&user)
	if user.UserId!=""{
		return user.UserId,"Ok"
	}
	return "","Error"
}







func UpdateDeliverCount(userId string,deliverCount int){
	db:=database.GetDB()
	otp:=OTPVerification{}
	db.Model(&otp).Where("user_id=?",userId).Update("deliver_count",deliverCount+1)
}
func InsertForgetPasswordStatus(sessionId string){
	db:=database.GetDB()
	forgetPasswordStatus:=ForgetPasswordStatus{}
	forgetPasswordStatus.SessionId=sessionId
	forgetPasswordStatus.VerificationStatus="U"
	db.Create(&forgetPasswordStatus)
}
func UpdateForgetPasswordStatus(sessionId string){
	db:=database.GetDB()
	forgetPasswordStatus:=ForgetPasswordStatus{}
	db.Model(&forgetPasswordStatus).Where("session_id = ?", sessionId).Update("verification_status", "V")
}

func CanUserUpdatePassword(sessionId string) string{
	db:=database.GetDB()
	forgetPasswordStatus:=ForgetPasswordStatus{}
	db.Where("session_id=? AND verification_status=?",sessionId,"V").Find(&forgetPasswordStatus)
	if forgetPasswordStatus.SessionId==""{
		return "Error"
	}
	return "Ok"
}

func DeleteForgetPasswordSession(sessionId string){
	db:=database.GetDB()
	db.Where("session_id=?",sessionId).Delete(&ForgetPasswordStatus{})
}

func IsProfileCreated(userId string) string{
	db:=database.GetDB()
	user:=User{}
	db.Where("user_id=?",userId).Find(&user)
	if(user.ProfileCreationStatus=="Y"){
		return "Ok"
	}
	return "Error"
}



// func GetTemporaryIdList(userId string) TempUserList{
// 	db:=database.GetDB()
// 	anonymousUser:=[]AnonymousUser{}
// 	db.Where("user_id=?",userId).Find(&anonymousUser)
// 	tempUser:=TempUserList{}
// 	tempUser.UserId=userId
// 	chatList:=[]ChatListElement{}
// 	for _,u :=range anonymousUser{
// 		c:=ChatListElement{}
// 		c.ChatId=u.ChatId
// 		c.TempId=u.AnonymousId
// 		chatList=append(chatList,c)

// 	}
// 	tempUser.ChatList=chatList
// 	return tempUser
// }

// func GetTempUserIdFromChatId(userId string,chatId string) string{
// 	db:=database.GetDB()
// 	anonymousUser:=AnonymousUser{}
// 	db.Where("user_id=? AND chat_id=?",userId,chatId).Find(&anonymousUser)
// 	return anonymousUser.AnonymousId
// }

func CheckPhoneNumberStatusDB(getPhoneStatusRequest GetPhoneStatusRequest) ([]int){
	db:=database.GetDB()
	phone:=getPhoneStatusRequest.PhoneList
	phoneStatusList:=make([]int,0)
	for _,p:=range phone{
		user:=User{}
		db.Table("users").Where("phone=?",p.Phone).Find(&user)
		if(user.UserId==""){
			phoneStatusList=append(phoneStatusList,0)
		}else{
			phoneStatusList=append(phoneStatusList,1)
		}
	}
	return phoneStatusList
}