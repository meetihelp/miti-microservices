package Authentication
import(
	"gopkg.in/go-playground/validator.v9"
	// "gopkg.in/go-playground/validator"
	"fmt"
)

type SanatizeData interface{
	doSanitization() string
}

func (userData User) doSanitization() string {
	validate :=validator.New()
	err:= validate.Struct(userData)
	if err!=nil{
		fmt.Println(err.Error())
		return "Error"
	}
	//WE CAN IMPROVE THIS FUCNTION BY CHECKING IF THE PHONE IS NUMERIC OR NOT
	if (userData.Phone =="") && (userData.Email==""){
		return "Error"
	}

	if (userData.Phone!="") && (len(userData.Phone)!=10){
		return "Error"
	}


	return "Ok"
}

func (otpVerification OTPVerification) doSanitization() string{
	// validate:=validator.New()
	// err:=validate.Struct(otpVerification)
	// if err!=nil{
	// 	fmt.Println(err.Error())
	// 	return "Error"
	// }
	return "Ok"
}

func (passwordChangeData PasswordChange) doSanitization() string{
	validate:=validator.New()
	err:=validate.Struct(passwordChangeData)
	if err!=nil{
		fmt.Println(err.Error())
		return "Error"
	}
	return "Ok"
}

func Sanatize(s SanatizeData) string{
	return s.doSanitization()
}
