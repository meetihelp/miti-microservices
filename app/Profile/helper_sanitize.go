package Profile
import(
	"gopkg.in/go-playground/validator.v9"
	"fmt"
)


type Sanatize_Data interface{
	do_sanitization() string
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
func Sanatize(s Sanatize_Data) string{
	return s.do_sanitization()
}