package Image

import(
	"gopkg.in/go-playground/validator.v9"
	// "gopkg.in/go-playground/validator"
	"fmt"
	"strings"
)


type SanatizeData interface{
	doSanitization() string
}


func(uploadImageHeader UploadImageHeader) doSanitization() string{
	validate :=validator.New()
	err:= validate.Struct(uploadImageHeader)
	if err!=nil{
		fmt.Println(err.Error())
		return "Error"
	}
	if(strings.ToLower(uploadImageHeader.AccessType)!="public" && strings.ToLower(uploadImageHeader.AccessType)!="private"){
		return "Error"
	}
	return "Ok"
}

func(getImageByIdData GetImageByIdRequest) doSanitization() string{
	validate :=validator.New()
	err:= validate.Struct(getImageByIdData)
	if err!=nil{
		fmt.Println(err.Error())
		return "Error"
	}
	if(getImageByIdData.ImageId==""){
		return "Error"
	}	
	return "Ok"
}

func Sanatize(s SanatizeData) string{
	return s.doSanitization()
}