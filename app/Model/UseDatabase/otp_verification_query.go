package UseDatabase

import(
	// "time"
	CD "app/Model/CreateDatabase"
	util "app/Utility"
)

func Verify_OTP(user_id string,otp string) (bool){
	db:=GetDB()
	otp_verification:=CD.OTP_verification{}
	db.Where("user_id=? AND verification_otp=?",user_id,otp).First(&otp_verification)
	if otp_verification.User_id==""{
		return false
	}
	return true
}

func Enter_verification_otp(id string,otp string){
	db:=GetDB()
	otp_verification:=CD.OTP_verification{}
	otp_verification.User_id=id
	otp_verification.Verification_otp=otp
	// otp_verification.CreatedAt =time.Now()
	otp_verification.CreatedAt=util.GetTime()
	db.Create(&otp_verification)
}

func Get_otp_verification_count(id string)(int,string){
	count:=0
	otp_verification:=CD.OTP_verification{}
	db:=GetDB()
	db.Where("user_id=?",id).Find(&otp_verification).Count(&count)
	// return len(otp_verification),otp_verification.CreatedAt
	return count,otp_verification.CreatedAt
}
