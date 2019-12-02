package Profile
import(
	"gopkg.in/go-playground/validator.v9"
	"fmt"
)


type SanatizeData interface{
	doSanitization() string
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