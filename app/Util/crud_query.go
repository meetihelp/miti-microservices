package Util
import(
	database "app/Database"
	"github.com/nu7hatch/gouuid"
	"time"
	"net/http"
	"fmt"
)
func get_cookie() http.Cookie{
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
func Insert_session(User_id string,ip_address string) string{
	cookie:= get_cookie()
	session:=Session{}
	session.Session_id=cookie.Value
	session.User_id=User_id
	session.IP=ip_address
	session.CreatedAt =time.Now()
	db:=database.GetDB()
	db.Create(&session)
	fmt.Println("Session inserted in Session Table")
	// return cookie
	return cookie.Value
}

func Get_user_id_from_session(session_id string) (string,string){
	db:=database.GetDB()
	session:=Session{}
	db.Where("session_id=?",session_id).First(&session)
	if session.User_id==""{
		return "","ERROR"
	}
	return session.User_id,"OK"
}

func Delete_session(session_id string) (string){
	db:=database.GetDB()
	fmt.Println("Delete ",session_id)
	db.Where("session_id=?",session_id).Delete(&Session{})
	return "OK"
}