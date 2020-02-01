package Util

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
   profile "miti-microservices/Profile"
	auth "miti-microservices/Authentication"
	chat "miti-microservices/Chat"
	util "miti-microservices/Util"
	gps "miti-microservices/GPS"
	// event "miti-microservices/Event"
	feed "miti-microservices/NewsFeed"
	image "miti-microservices/Media/Image"
	social "miti-microservices/Social"
	security "miti-microservices/Security"
	privacy "miti-microservices/Privacy"
	database "miti-microservices/Database"
)

type DeleteProfileData struct{
	Phone string `json:"Phone"`
}

func DeleteProfile(w http.ResponseWriter, r *http.Request){
	db:=database.DBConnection()
	
	requestBody,err:=ioutil.ReadAll(r.Body)
	if (err!=nil){
		fmt.Println("Delete Profile Data Line 19")
	}

	data:=DeleteProfileData{}
	err=json.Unmarshal(requestBody,&data)
	if(err!=nil){
		fmt.Println(err)
	}

	phone:=data.Phone
	userId,_:=auth.GetUserIdFromPhone(db,phone)

	err=db.Where("user_id=?",userId).Delete(&privacy.BoardContent{}).Error
	if(err!=nil){
		fmt.Print("Line 32")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&privacy.Board{}).Error
	if(err!=nil){
		fmt.Print("Line 38")
		fmt.Println(err)
	}	

	err=db.Where("actual_user_id=?",userId).Delete(&chat.ChatDetail{}).Error
	if(err!=nil){
		fmt.Print("Line 44")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&chat.Chat{}).Error
	if(err!=nil){
		fmt.Print("Line 50")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&social.GroupPoolStatus{}).Error
	if(err!=nil){
		fmt.Print("Line 56")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&social.Group{}).Error
	if(err!=nil){
		fmt.Print("Line 62")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&profile.Interest{}).Error
	if(err!=nil){
		fmt.Print("Line 68")
		fmt.Println(err)
	}

	err=db.Where("user_id1=? OR user_id2=?",userId,userId).Delete(&profile.Match{}).Error
	if(err!=nil){
		fmt.Print("Line 74")
		fmt.Println(err)
	}

	err=db.Where("sender_user_id=? or phone=?",userId,phone).Delete(&chat.MessageRequest{}).Error
	if(err!=nil){
		fmt.Print("Line 80")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&feed.NewsFeedReaction{}).Error
	if(err!=nil){
		fmt.Print("Line 32")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&auth.OTPVerification{}).Error
	if(err!=nil){
		fmt.Print("Line 32")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&social.PoolLog{}).Error
	if(err!=nil){
		fmt.Print("Line 32")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&social.PoolStatus{}).Error
	if(err!=nil){
		fmt.Print("Line 32")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&social.PoolWaiting{}).Error
	if(err!=nil){
		fmt.Print("Line 32")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&social.Pool{}).Error
	if(err!=nil){
		fmt.Print("Line 32")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&security.PrimaryTrustChain{}).Error
	if(err!=nil){
		fmt.Print("Line 32")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&profile.Profile{}).Error
	if(err!=nil){
		fmt.Print("Line 32")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&profile.QuestionResponse{}).Error
	if(err!=nil){
		fmt.Print("Line 32")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&util.Session{}).Error
	if(err!=nil){
		fmt.Print("Line 32")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&util.TemporarySession{}).Error
	if(err!=nil){
		fmt.Print("Line 32")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&gps.UserCurrentLocation{}).Error
	if(err!=nil){
		fmt.Print("Line 32")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&feed.UserFeedStatus{}).Error
	if(err!=nil){
		fmt.Print("Line 32")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&image.UserImage{}).Error
	if(err!=nil){
		fmt.Print("Line 32")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&social.UserPool{}).Error
	if(err!=nil){
		fmt.Print("Line 32")
		fmt.Println(err)
	}

	err=db.Where("user_id=?",userId).Delete(&auth.User{}).Error
	if(err!=nil){
		fmt.Print("Line 32")
		fmt.Println(err)
	}
}