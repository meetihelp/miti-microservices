package UseDatabase

import (
	
	CD "app/Model/CreateDatabase"
	util "app/Utility"
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
		// err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
		status:=util.Comapare_password(user.Password,password)
		if !status{
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
		// err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
		status:=util.Comapare_password(user.Password,password)
		if !status{
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
	status:=util.Comapare_password(user.Password,password)
	if !status{
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
	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(new_Password), bcrypt.DefaultCost)
	// new_Password = string(hashedPassword)
	new_Password = util.Generate_encrypted_password(new_Password)

	user:=CD.User{}
	db.Model(&user).Where("user_id = ?", user_id).Update("password", new_Password)
}

func GetAllUser() ([]string){
	db:=GetDB()
	user:=[]CD.User{}
	db.Find(&user)

	UserList:=make([]string,0)
	for _,id := range user{
		UserList=append(UserList,id.User_id)
	}
	return UserList
}