package UseDatabase

import(
	CD "app/Model/CreateDatabase"
)

func Get_user_detail(user_id string) (string,string){
	db:=GetDB()
	user:=CD.User{}
	db.Where("user_id=?",user_id).First(&user)
	return user.Email , user.Phone
}