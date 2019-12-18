package Util

import(
	"golang.org/x/crypto/bcrypt"
)

func GenerateEncryptedPassword(password string) string{
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func ComaparePassword(hashed_password string,password string) bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hashed_password),[]byte(password))
	if err!=nil{
		return false
	}

	return true
}