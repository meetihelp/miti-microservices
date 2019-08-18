package UseDatabase

import(
	CD "app/Model/CreateDatabase"
)

func GetChatMessages(chat_id string,offset int,num_of_rows int)([]CD.Chat){
	chat:=[]CD.Chat{}
	db:=GetDB()
	db.Offset(offset).Limit(num_of_rows).Where("chat_id=?",chat_id).Find(&chat)
	return chat
}

func Check_correct_chat(user_id string,chat_id string) string{
	db:=GetDB()
	chatDetail:=CD.ChatDetail{}
	db.Where("actual_user_id=? AND chat_id=?",user_id,chat_id).First(&chatDetail)
	if chatDetail.Chat_id==""{
		return "ERROR"
	}

	return "OK"
}

func GetChatDetail(user_id string) ([]CD.ChatDetail,string){
	db:=GetDB()
	chatDetail:=[]CD.ChatDetail{}
	db.Where("actual_user_id=?",user_id).Find(&chatDetail)
	return chatDetail,"OK"
}

func GetChatByMessageId(messageId []string)([]CD.Chat){
	chat:=[]CD.Chat{}
	db:=GetDB()
	temp_chat:=CD.Chat{}
	for _,id :=range messageId{
		db.Where("message_id=?",id).Find(&temp_chat)	
		chat=append(chat,temp_chat)
	}
	return chat
}


func GetUnreadMessage(user_id string)([]string){
	db:=GetDB()
	readBy:=[]CD.ReadBy{}
	db.Where("user_id=? AND status=unread",user_id).Find(&readBy)
	messageId:=make([]string,0)
	for _,data:=range readBy{
		messageId=append(messageId,data.Message_id)
	}
	return messageId
}
