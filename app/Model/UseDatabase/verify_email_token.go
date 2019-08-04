package UseDatabase

import(
	CD "app/Model/CreateDatabase"
)

func Verify_Email(token string) (string,bool){
	db:=GetDB()
	email_verification:=CD.Email_verification{}
	db.Where("verification_token=?",token).First(&email_verification)
	if email_verification.User_id==""{
		return "",false
	}
	return email_verification.User_id,true
}