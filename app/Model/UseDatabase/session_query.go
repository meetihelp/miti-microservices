package UseDatabase

import(
	CD "app/Model/CreateDatabase"
	"fmt"
)

func Get_user_id_from_session(session_id string) (string,string){
	db:=GetDB()
	session:=CD.Session{}
	db.Where("session_id=?",session_id).First(&session)
	if session.User_id==""{
		return "","ERROR"
	}
	return session.User_id,"OK"
}

func Delete_session(session_id string) (string){
	db:=GetDB()
	fmt.Println("Delete ",session_id)
	db.Where("session_id=?",session_id).Delete(&CD.Session{})
	return "OK"
}