package Util
import(
	database "app/Database"
	"github.com/nu7hatch/gouuid"
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

func InsertSessionValue(tempSession string,userId string,ipAddress string){
	session:=Session{}
	session.SessionId=tempSession
	session.UserId=userId
	session.IP=ipAddress
	// session.CreatedAt =time.Now()
	session.CreatedAt=GetTime()
	db:=database.GetDB()
	db.Create(&session)
	fmt.Println("Session inserted in Session Table")
}

func InsertTemporarySession(UserId string,ipAddress string) string{
	cookie:= getCookie()
	session:=TemporarySession{}
	session.TemporarySessionId=cookie.Value
	session.UserId=UserId
	session.IP=ipAddress
	// session.CreatedAt =time.Now()
	session.CreatedAt=GetTime()
	db:=database.GetDB()
	db.Create(&session)
	fmt.Println("Session inserted in User Verification Session Table")
	// return cookie
	return cookie.Value
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



func DeleteSession(sessionId string) (string){
	db:=database.GetDB()
	fmt.Println("Delete ",sessionId)
	db.Where("session_id=?",sessionId).Delete(&Session{})
	return "Ok"
}

func DeleteTemporarySession(sessionId string) (string){
	db:=database.GetDB()
	fmt.Println("Delete ",sessionId)
	db.Where("temporary_session_id=?",sessionId).Delete(&TemporarySession{})
	return "Ok"
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