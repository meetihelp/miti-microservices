
package Util
import(
	database "miti-microservices/Database"
	"github.com/nu7hatch/gouuid"
	"github.com/jinzhu/gorm"
	"time"
	"net/http"
	"fmt"
)
func getCookie() http.Cookie{
	expire := time.Now().Add(10 * time.Minute)
	cookie:=http.Cookie{}
	cookie.Name="cookie"
	temp, _ := uuid.NewV4()
	cookie.Value=temp.String()
	cookie.Expires=expire
	cookie.Path="/"
	cookie.MaxAge=90000
	return cookie
}
func InsertSession(UserId string,ipAddress string) string{
	cookie:= getCookie()
	session:=Session{}
	session.SessionId=cookie.Value
	session.UserId=UserId
	session.IP=ipAddress
	// session.CreatedAt =time.Now()
	session.CreatedAt=GetTime()
	db:=database.GetDB()
	db.Create(&session)
	fmt.Println("Session inserted in Session Table")
	// return cookie
	return cookie.Value
}

func InsertSessionValue(db *gorm.DB,tempSession string,userId string,ipAddress string) bool{
	session:=Session{}
	session.SessionId=tempSession
	session.UserId=userId
	session.IP=ipAddress
	// session.CreatedAt =time.Now()
	session.CreatedAt=GetTime()
	// db:=database.GetDB()
	err:=db.Create(&session).Error
	if(err!=nil){
		return true
	}
	return false
	// fmt.Println("Session inserted in Session Table")
}

func InsertTemporarySession(db *gorm.DB,UserId string,ipAddress string) (string,bool){
	cookie:= getCookie()
	session:=TemporarySession{}
	session.TemporarySessionId=cookie.Value
	session.UserId=UserId
	session.IP=ipAddress
	session.CreatedAt=GetTime()
	err:=db.Create(&session).Error
	if(err!=nil){
		return "",true
	}
	// fmt.Println("Session inserted in User Verification Session Table")
	return cookie.Value,false
}

func GetUserIdFromSession(sessionId string) (string,string){
	db:=database.GetDB()
	session:=Session{}
	db.Where("session_id=?",sessionId).First(&session)
	if session.UserId==""{
		return "","Error"
	}
	return session.UserId,"Ok"
}
func GetUserIdFromTemporarySession(sessionId string) (string,string){
	db:=database.GetDB()
	session:=TemporarySession{}
	fmt.Println(sessionId)
	db.Where("temporary_session_id=?",sessionId).First(&session)
	if session.UserId==""{
		return "","Error"
	}
	return session.UserId,"Ok"
}

func GetUserIdFromSession2(db *gorm.DB,sessionId string) (string,string){
	// db:=database.GetDB()
	session:=Session{}
	db.Where("session_id=?",sessionId).First(&session)
	if session.UserId==""{
		return "","Error"
	}
	return session.UserId,"Ok"
}
func GetUserIdFromTemporarySession2(db *gorm.DB,sessionId string) (string,string){
	// db:=database.GetDB()
	session:=TemporarySession{}
	fmt.Println(sessionId)
	db.Where("temporary_session_id=?",sessionId).First(&session)
	if session.UserId==""{
		return "","Error"
	}
	return session.UserId,"Ok"
}

func GetUserIdFromSession3(db *gorm.DB,sessionId string) (string,string,bool){
	session:=Session{}
	err:=db.Where("session_id=?",sessionId).First(&session).Error
	if gorm.IsRecordNotFoundError(err){
		return "","Error",false
	}
	if(err!=nil){
		return "","Error",true
	}
	return session.UserId,"Ok",false
}
func GetUserIdFromTemporarySession3(db *gorm.DB,sessionId string) (string,string,bool){
	session:=TemporarySession{}
	err:=db.Where("temporary_session_id=?",sessionId).First(&session).Error
	if gorm.IsRecordNotFoundError(err){
		return "","Error",false
	}
	if(err!=nil){
		return "","Error",true
	}
	return session.UserId,"Ok",false
}



func DeleteSession(sessionId string) (string){
	db:=database.GetDB()
	fmt.Println("Delete ",sessionId)
	db.Where("session_id=?",sessionId).Delete(&Session{})
	return "Ok"
}

func DeleteTemporarySession(db *gorm.DB,sessionId string) (bool){
	// db:=database.GetDB()
	fmt.Println("Delete ",sessionId)
	err:=db.Where("temporary_session_id=?",sessionId).Delete(&TemporarySession{}).Error
	if(err!=nil){
		return true
	}
	return false
}

// func InsertOtp(userId string,sessionId string) string{
// 	db:=database.GetDB()
// 	otp:=util.GenerateOtp()
// 	otpVerification:=OtpVerification{}
// 	otpVerification.UserId=userId
// 	otpVerification.SessionId=sessionId
// 	otpVerification.OTP=otp
// 	otpVerification.CreatedAt=GetTime()
// 	db.Create(&otpVerification)
// 	return otp
// }