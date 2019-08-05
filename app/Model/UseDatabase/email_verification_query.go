package UseDatabase

import(
	"time"
	CD "app/Model/CreateDatabase"
)


func Enter_email_verification(id string,token string){
	db:=GetDB()
	email_verification:=CD.Email_verification{}
	email_verification.User_id=id
	email_verification.Verification_token=token
	email_verification.CreatedAt=time.Now()
	db.Create(&email_verification)
}

func Get_Email_verification_count(id string)(int,time.Time){
	email_verification:=CD.Email_verification{}
	db:=GetDB()
	db.Where("user_id=?",id).Find(&email_verification)
	// return len(email_verification),email_verification.CreatedAt
	return 2,email_verification.CreatedAt

}


func Delete_all_email_verification(id string){
	db:=GetDB()
	db.Where("user_id=?",id).Delete(&CD.Email_verification{})
}

func Verify_Email(token string) (string,bool){
	db:=GetDB()
	email_verification:=CD.Email_verification{}
	db.Where("verification_token=?",token).First(&email_verification)
	if email_verification.User_id==""{
		return "",false
	}
	return email_verification.User_id,true
}