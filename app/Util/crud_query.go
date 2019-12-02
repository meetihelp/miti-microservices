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

func InsertUserVerificationSession(UserId string,ipAddress string) string{
	cookie:= getCookie()
	session:=UserVerificationSession{}
	session.UserVerificationSessionId=cookie.Value
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
func GetUserIdFromUserVerificationSession(sessionId string) (string,string){
	db:=database.GetDB()
	session:=UserVerificationSession{}
	fmt.Println(sessionId)
	db.Where("user_verification_session_id=?",sessionId).First(&session)
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

func DeleteUserVerificationSession(sessionId string) (string){
	db:=database.GetDB()
	fmt.Println("Delete ",sessionId)
	db.Where("user_verification_session_id=?",sessionId).Delete(&UserVerificationSession{})
	return "Ok"
}