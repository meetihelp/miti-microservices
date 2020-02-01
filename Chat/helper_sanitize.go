package Chat

import(
	"gopkg.in/go-playground/validator.v9"
	// "gopkg.in/go-playground/validator"
	"fmt"
	"strings"
)


type SanatizeData interface{
	doSanitization() string
}


func(sendChatImageHeader SendChatImageHeader) doSanitization() string{
	if(strings.ToLower(sendChatImageHeader.AccessType)!="public" && strings.ToLower(sendChatImageHeader.AccessType)!="private"){
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

func(chatDetailDsData GetChatRequest) doSanitization() string {
	validate :=validator.New()
	err:= validate.Struct(chatDetailDsData)
	if err!=nil{
		fmt.Println(err.Error())
		return "Error"
	}
	return "Ok"
}

func(chatDetailData ChatDetailRequest) doSanitization() string{
	validate :=validator.New()
	err:= validate.Struct(chatDetailData)
	if err!=nil{
		fmt.Println(err.Error())
		return "Error"
	}
	return "Ok"	
}

func(sendMessageRequestData SendMessageRequestDS) doSanitization() string{
	validate :=validator.New()
	err:= validate.Struct(sendMessageRequestData)
	if err!=nil{
		fmt.Println(err.Error())
		return "Error"
	}
	return "Ok"	
}

func(actionMessageRequestData ActionMessageRequestDS) doSanitization() string{
	validate :=validator.New()
	err:= validate.Struct(actionMessageRequestData)
	if err!=nil{
		fmt.Println(err.Error())
		return "Error"
	}

	if(actionMessageRequestData.Phone==""){
		return "Error"
	}
	return "Ok"	
}

func Sanatize(s SanatizeData) string{
	return s.doSanitization()
}
