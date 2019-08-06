package Utility

import(
	"golang.org/x/crypto/bcrypt"
)

func Generate_encrypted_password(password string) string{
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func Comapare_password(hashed_password string,password string) bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hashed_password),[]byte(password))
	if err!=nil{
		return false
	}

	return true
}