package UseDatabase

import(
	CD "app/Model/CreateDatabase"
)

func Enter_verification_email(id string,token string){
	db:=GetDB()
	email_verification:=CD.Email_verification{}
	email_verification.User_id=id
	email_verification.Verification_token=token
	db.Create(&email_verification)
}