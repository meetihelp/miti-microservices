package Authentication

import(
	"net/http"
	"fmt"
	util "app/Util"
)


func Logout(w http.ResponseWriter,r *http.Request){
	fmt.Println("Delete")
	logoutHeader:=LogoutHeader{}
	util.GetHeader(r,&logoutHeader)
	sessionId:=logoutHeader.Cookie
	logoutStatus:=util.DeleteSession(sessionId)
	if logoutStatus=="Ok"{
		util.Message(w,200)
		return
	}
}