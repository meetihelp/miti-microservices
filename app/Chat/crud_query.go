package Chat

import(
	database "app/Database"
	"time"
	"fmt"
)

func GetChatMessages(chat_id string,offset int,num_of_rows int)([]Chat){
	chat:=[]Chat{}
	db:=database.GetDB()
	db.Offset(offset).Order("created_at").Limit(num_of_rows).Where("chat_id=?",chat_id).Find(&chat)
	return chat
}

func Check_correct_chat(user_id string,chat_id string) string{
	db:=database.GetDB()
	chatDetail:=ChatDetail{}
	db.Where("actual_user_id=? AND chat_id=?",user_id,chat_id).First(&chatDetail)
	if chatDetail.Chat_id==""{
		return "ERROR"
	}

	return "OK"
}

func GetChatDetail(user_id string,offset int,num_of_chat int) ([]ChatDetail,string){
	db:=database.GetDB()
	chatDetail:=[]ChatDetail{}
	db.Offset(offset).Limit(num_of_chat).Where("actual_user_id=?",user_id).Find(&chatDetail)
	return chatDetail,"OK"
}

func GetChatByMessageId(messageId []string)([]Chat){
	chat:=[]Chat{}
	db:=database.GetDB()
	temp_chat:=Chat{}
	for _,id :=range messageId{
		db.Where("message_id=?",id).Find(&temp_chat)	
		chat=append(chat,temp_chat)
	}
	return chat
}


func GetUnreadMessage(user_id string)([]string){
	db:=database.GetDB()
	readBy:=ReadBy{}
	user_list:=GetAnonymousList(user_id)
	messageId:=make([]string,0)
	for _,user:= range user_list{
		db.Where("user_id=? AND status=unfetched",user).Find(&readBy)	
		messageId=append(messageId,readBy.Message_id)
	}
	return messageId
}

func GetAnonymousList(user_id string)([]string){
	db:=database.GetDB()
	chatDetail:=[]ChatDetail{}
	db.Where("actual_user_id=?",user_id).Find(&chatDetail)
	user_list:=make([]string,0)
	for _,data:=range chatDetail{
		user_list=append(user_list,data.Temp_User_id)
	}
	return user_list
}

func EnterReadBy(user_list []string,message_id string){
	db:=database.GetDB()
	readBy:=ReadBy{}
	readBy.ReadAt=time.Now()
	readBy.Message_id=message_id
	readBy.Status="unfetched"
	for _,user := range user_list{
		readBy.User_id=user
		db.Create(&readBy)
	}
}

func ChangeFetchStatus(user_id string,message_id []string){
	db:=database.GetDB()
	db.Table("ReadBy").Where("user_id=? AND message_id IN (?)",user_id,message_id).Update("status","fetched")
}


func GetUserListFromChatId(chat_id string)([]string){
	db:=database.GetDB()
	chatDetail:=[]ChatDetail{}
	db.Where("chat_id=?",chat_id).Find(&chatDetail)
	user_list:=make([]string,0)
	for _,data:= range chatDetail{
		user_list=append(user_list,data.Temp_User_id)
	}
	return user_list
}

func UpdateChatTime(chat_id string, lastupdate string) error{
	db:=database.GetDB()
	db.Table("chat_details").Where("chat_id=?",chat_id).Update("last_update",lastupdate)
	return nil
}

func GetLastChatIndex(chat_id string) int{
	db:=database.GetDB()
	chat:=Chat{}
	db.Order("index desc").Where("chat_id=?",chat_id).First(&chat)
	fmt.Println(chat)
	if chat.Chat_id==""{
		return 0
	}
	return chat.Index
}

func GetChatAfterIndexMessages(chat_id string, offset int, num_of_chat int, index int)([]Chat){
	db:=database.GetDB()
	chat:=[]Chat{}
	db.Order("created_at").Where("chat_id=? AND index>?",chat_id,index).Find(&chat)
	return chat
}