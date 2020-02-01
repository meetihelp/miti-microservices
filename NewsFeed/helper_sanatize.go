package NewsFeed
import(
	"gopkg.in/go-playground/validator.v9"
	// "gopkg.in/go-playground/validator"
	"fmt"
)

type SanatizeData interface{
	doSanitization() string
}

func(getNewsFeedArticleData GetNewsArticleDS) doSanitization() string{
	validate:=validator.New()
	err:=validate.Struct(getNewsFeedArticleData)
	if err!=nil{
		fmt.Println(err.Error())
		return "Error"
	}
	// if(getNewsFeedArticleData.Label==""){
	// 	return "Error"
	// }
	return "Ok"
}


func Sanatize(s SanatizeData) string{
	return s.doSanitization()
}
