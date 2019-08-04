package UseDatabase

import(
	CD "app/Model/CreateDatabase"
)

func Change_Verification_Status(user_id string){
	db:=GetDB()
	user:=CD.User{}
	db.Model(&user).Where("user_id=?",user_id).Update("status","V")
}