package UseDatabase

import(
	CD "app/Model/CreateDatabase"
)

func Enter_verification_otp(id string,otp int){
	db:=GetDB()
	otp_verification:=CD.OTP_verification{}
	otp_verification.User_id=id
	otp_verification.Verification_otp=otp
	db.Create(&otp_verification)
}