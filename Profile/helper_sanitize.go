package Profile
import(
	"gopkg.in/go-playground/validator.v9"
	// "gopkg.in/go-playground/validator"
	"fmt"
)


type SanatizeData interface{
	doSanitization() string
}

func(updateIPIPRequest UpdateIPIPRequest) doSanitization() string{
	validate :=validator.New()
	err:= validate.Struct(updateIPIPRequest)
	if err!=nil{
		fmt.Println(err.Error())
		return "Error"
	}
	return "Ok"
}

func(preferenceRequestData UpdatePreferenceRequest) doSanitization() string{
	validate :=validator.New()
	err:= validate.Struct(preferenceRequestData)
	if err!=nil{
		fmt.Println(err.Error())
		return "Error"
	}
	return "Ok"
}

func(getProfileRequest GetProfileRequest) doSanitization() string{
	validate :=validator.New()
	err:= validate.Struct(getProfileRequest)
	if err!=nil{
		fmt.Println(err.Error())
		return "Error"
	}
	return "Ok"
}

func (profileData Profile) doSanitization() string{
	validate :=validator.New()
	err:= validate.Struct(profileData)
	if err!=nil{
		fmt.Println(err.Error())
		return "Error"
	}
	return "Ok"
}
func Sanatize(s SanatizeData) string{
	return s.doSanitization()
}
