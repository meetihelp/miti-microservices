package UseDatabase

import (
	"golang.org/x/crypto/bcrypt"
	CD "app/Model/CreateDatabase"
)

func Check_user(user_data CD.User)(string,string){
	db:=GetDB()
	email:=user_data.Email
	phone:=user_data.Phone
	password:=user_data.Password
	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user_data.Password), bcrypt.DefaultCost)
	// password := string(hashedPassword)

	user:=CD.User{}
	if email!=""{
		db.Where("email=?",email).First(&user)
		err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
		if err!=nil{
			return "","WRONG_PASSWORD"
		}
		if user.User_id==""{
			return user.User_id,"NO_USER"
		} 
		if user.User_id !="" && user.Status=="U"{
			return user.User_id,"UNVERIFIED"
		}

		return user.User_id,"OK"
	}

	if phone!=""{
		db.Where("phone=? AND password=?",phone,password).First(&user)
		err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
		if err!=nil{
			return "","WRONG_PASSWORD"
		}
		if user.User_id==""{
			return user.User_id,"NO_USER"
		} 
		if user.User_id !="" && user.Status=="U"{
			return user.User_id,"UNVERIFIED"
		}

		return user.User_id,"OK"

	}
	return "","ERROR"
}