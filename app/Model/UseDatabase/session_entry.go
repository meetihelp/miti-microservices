package UseDatabase

import(
	CD "app/Model/CreateDatabase"
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
func Insert_session(User_id string,ip_address string) http.Cookie{
	cookie:= get_cookie()
	session:=CD.Session{}
	session.Session_id=cookie.Value
	session.User_id=User_id
	session.IP=ip_address
	session.CreatedAt =time.Now()
	fmt.Println("Session")
	fmt.Println(session)
	db:=GetDB()
	db.Create(&session)
	return cookie
}