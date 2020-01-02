package Chat

import(
	util "miti-microservices/Util"
	database "miti-microservices/Database"
	// "time"
	"fmt"
)

func GetChatMessages(chatId string,offset int,numOfRows int)([]Chat){
	chat:=[]Chat{}
	db:=database.GetDB()
	db.Offset(offset).Order("created_at").Limit(numOfRows).Where("chat_id=?",chatId).Find(&chat)
	return chat
}

func ChatInsertDB(chatData Chat) {
	index:=GetLastChatIndex(chatData.ChatId)
	index=index+1
	chatData.Index=index
	db:=database.GetDB()
	db.Create(&chatData)
}
func CheckCorrectChat(userId string,chatId string) string{
	db:=database.GetDB()
	chatDetail:=ChatDetail{}
	db.Where("actual_user_id=? AND chat_id=?",userId,chatId).First(&chatDetail)
	if chatDetail.ChatId==""{
		return "Error"
	}

	return "Ok"
}

func GetChatDetail(userId string,offset int,numOfChat int) ([]ChatDetail,string){
	db:=database.GetDB()
	chatDetail:=[]ChatDetail{}
	db.Offset(offset).Limit(numOfChat).Where("actual_user_id=?",userId).Find(&chatDetail)
	return chatDetail,"Ok"
}

func GetChatByMessageId(messageId []string)([]Chat){
	chat:=[]Chat{}
	db:=database.GetDB()
	tempChat:=Chat{}
	for _,id :=range messageId{
		db.Where("message_id=?",id).Find(&tempChat)	
		chat=append(chat,tempChat)
	}
	return chat
}


func GetUnreadMessage(userId string)([]string){
	db:=database.GetDB()
	readBy:=ReadBy{}
	userList:=GetAnonymousList(userId)
	messageId:=make([]string,0)
	for _,user:= range userList{
		db.Where("user_id=? AND status=unfetched",user).Find(&readBy)	
		messageId=append(messageId,readBy.MessageId)
	}
	return messageId
}

func GetAnonymousList(userId string)([]string){
	db:=database.GetDB()
	chatDetail:=[]ChatDetail{}
	db.Where("actual_user_id=?",userId).Find(&chatDetail)
	userList:=make([]string,0)
	for _,data:=range chatDetail{
		userList=append(userList,data.TempUserId)
	}
	return userList
}

func EnterReadBy(userList []string,messageId string){
	db:=database.GetDB()
	readBy:=ReadBy{}
	// readBy.ReadAt=time.Now()
	readBy.ReadAt=util.GetTime()
	readBy.MessageId=messageId
	readBy.Status="unfetched"
	for _,user := range userList{
		readBy.UserId=user
		db.Create(&readBy)
	}
}

func ChangeFetchStatus(userId string,messageId []string){
	db:=database.GetDB()
	db.Table("ReadBy").Where("user_id=? AND message_id IN (?)",userId,messageId).Update("status","fetched")
}


func GetUserListFromChatId(chatId string)([]string){
	db:=database.GetDB()
	chatDetail:=[]ChatDetail{}
	db.Where("chat_id=?",chatId).Find(&chatDetail)
	userList:=make([]string,0)
	for _,data:= range chatDetail{
		userList=append(userList,data.TempUserId)
	}
	return userList
}

func UpdateChatTime(chatId string, lastUpdate string) error{
	db:=database.GetDB()
	db.Table("chat_details").Where("chat_id=?",chatId).Update("last_update",lastUpdate)
	return nil
}

func GetLastChatIndex(chatId string) int{
	db:=database.GetDB()
	chat:=Chat{}
	db.Order("index desc").Where("chat_id=?",chatId).First(&chat)
	fmt.Println(chat)
	if chat.ChatId==""{
		return 0
	}
	return chat.Index
}

func GetChatAfterIndexMessages(chatId string, offset int, numOfChat int, index int)([]Chat){
	db:=database.GetDB()
	chat:=[]Chat{}
	db.Order("created_at").Limit(numOfChat).Where("chat_id=? AND index>?",chatId,index).Find(&chat)
	return chat
}