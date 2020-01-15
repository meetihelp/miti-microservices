package Chat

import(
	"gopkg.in/go-playground/validator.v9"
	// "gopkg.in/go-playground/validator"
	"fmt"
)


type SanatizeData interface{
	doSanitization() string
}

func(sendChatImageHeader SendChatImageHeader) doSanitization() string{
	if(sendChatImageHeader.AccessType!="Public" && sendChatImageHeader.AccessType!="Private"){
		return "Error"
	}
	return "Ok"
}

func(chatData Chat) doSanitization() string {
	validate :=validator.New()
	err:= validate.Struct(chatData)
	if err!=nil{
		fmt.Println(err.Error())
		return "Error"
	}

	return "Ok"
}

func(chatRequestData ChatRequest) doSanitization() string {
	validate :=validator.New()
	err:= validate.Struct(chatRequestData)
	if err!=nil{
		fmt.Println(err.Error())
		return "Error"
	}

	return "Ok"
}

func(chatAfterTimeData ChatAfterTime) doSanitization() string {
	validate :=validator.New()
	err:= validate.Struct(chatAfterTimeData)
	if err!=nil{
		fmt.Println(err.Error())
		return "Error"
	}

	return "Ok"
}

func(chatDetailDsData ChatDetailDs) doSanitization() string {
	validate :=validator.New()
	err:= validate.Struct(chatDetailDsData)
	if err!=nil{
		fmt.Println(err.Error())
		return "Error"
	}
	return "Ok"
}

func Sanatize(s SanatizeData) string{
	return s.doSanitization()
}
