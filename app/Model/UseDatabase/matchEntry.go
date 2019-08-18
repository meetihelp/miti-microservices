package UseDatabase

import(
	CD "app/Model/CreateDatabase"
	util "app/Utility"
	"time"
)

func Enter_Match_user(user_id1 string,user_id2 string){	

	chatID:=util.Generate_token()
	temp_user1:=util.Generate_token()
	temp_user2:=util.Generate_token()

	Enter_Anonymous_User(user_id1,temp_user2,chatID,"one-to-one",1)
	Enter_Anonymous_User(user_id2,temp_user1,chatID,"one-to-one",2)

}

func Enter_Anonymous_User(user_id string,temp_user_id string,chat_id string,chat_type string,user_index int){
	db:=GetDB()
	anonymousUser:=CD.AnonymousUser{}
	anonymousUser.User_id=user_id
	anonymousUser.Anonymous_id=temp_user_id
	anonymousUser.Chat_id=chat_id
	anonymousUser.CreatedAt=time.Now()
	anonymousUser.Status="None"

	chatDetail:=CD.ChatDetail{}
	chatDetail.Temp_User_id=temp_user_id
	chatDetail.Actual_User_id=user_id
	chatDetail.Chat_id=chat_id
	chatDetail.Chat_type=chat_type
	chatDetail.CreatedAt=anonymousUser.CreatedAt
	chatDetail.User_index=user_index

	db.Create(&anonymousUser)
	db.Create(&chatDetail)
}
