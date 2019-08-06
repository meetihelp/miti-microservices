package UseDatabase

import (
	"golang.org/x/crypto/bcrypt"
	CD "app/Model/CreateDatabase"
	// "fmt"
)

func Check_user(user_data CD.User)(string,string){
	db:=GetDB()
	email:=user_data.Email
	phone:=user_data.Phone
	password:=user_data.Password

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
		db.Where("phone=?",phone).First(&user)
		err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
		if err!=nil{
			// fmt.Println(err.Error())
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

func Check_user_by_id(id string,password string) string{
	db:=GetDB()
	user:=CD.User{}
	db.Where("user_id=?",id).First(&user)
	err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
	if err!=nil{
		// fmt.Println(err.Error())
		return "WRONG_PASSWORD"
	}
	if user.User_id==""{
		return "NO_USER"
	} 
	if user.User_id !="" && user.Status=="U"{
		return "UNVERIFIED"
	}

	return "OK"
}
func IsUserVerified(id string) bool{
	db:=GetDB()
	user:=CD.User{}
	db.Where("user_id=?",id).First(&user)
	if user.Status=="U"{
		return false
	} else{
		return true
	}
}

func Get_user_detail(user_id string) (string,string){
	db:=GetDB()
	user:=CD.User{}
	db.Where("user_id=?",user_id).First(&user)
	return user.Email , user.Phone
}

func Update_Password(user_id string,new_Password string){
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(new_Password), bcrypt.DefaultCost)
	new_Password = string(hashedPassword)

	user:=CD.User{}
	db.Model(&user).Where("user_id = ?", user_id).Update("password", new_Password	)
}