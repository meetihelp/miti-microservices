package CreateDatabase

import(
	"gopkg.in/go-playground/validator.v9"
	"fmt"
)


type Sanatize_Data interface{
	do_sanitization() string
}

func (user_data User) do_sanitization() string {
	validate :=validator.New()
	err:= validate.Struct(user_data)
	if err!=nil{
		fmt.Println(err.Error())
		return "ERROR"
	}
	//WE CAN IMPROVE THIS FUCNTION BY CHECKING IF THE PHONE IS NUMERIC OR NOT
	if (user_data.Phone =="") && (user_data.Email==""){
		return "ERROR"
	}

	if (user_data.Phone!="") && (len(user_data.Phone)!=10){
		return "ERROR"
	}


	return "OK"
}

func (profile_data Profile) do_sanitization() string{
	validate :=validator.New()
	err:= validate.Struct(profile_data)
	if err!=nil{
		fmt.Println(err.Error())
		return "ERROR"
	}
	return "OK"
}

func (otp_verification OTP_verification) do_sanitization() string{
	validate:=validator.New()
	err:=validate.Struct(otp_verification)
	if err!=nil{
		fmt.Println(err.Error())
		return "ERROR"
	}
	return "OK"
}

func Sanatize(s Sanatize_Data) string{
	return s.do_sanitization()
}