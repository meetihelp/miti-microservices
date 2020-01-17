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

func GetChatByRequestId(userId string,requestId string)Chat{
	db:=database.GetDB()
	chat:=Chat{}
	db.Where("user_id=? AND request_id=?",userId,requestId).Find(&chat)
	return chat
}

func ChatInsertDB(chatData Chat,lastUpdate string) (Chat,[]Chat,int) {
	// index:=GetLastChatIndex(chatData.ChatId)
	// index=index+1
	// chatData.Index=index
	chat:=Chat{}
	db:=database.GetDB()
	unSyncedChat:=[]Chat{}
	chatId:=chatData.ChatId
	err:=db.Order("created_at desc").Where("chat_id=? AND created_at>?",chatId,lastUpdate).Find(&unSyncedChat).Error
	code:=200
	if(err!=nil){
		fmt.Print("ChatInsertDB Error 1")
		fmt.Println(err)
		code=1006
	}
	err=db.Where("user_id=? AND request_id=?",chatData.UserId,chatData.RequestId).Find(&chat).Error
	if(err!=nil){
		fmt.Print("ChatInsertDB Error 2")
		fmt.Println(err)
		// code=1006
	}
	fmt.Println("ChatInsertDB")
	fmt.Println(chat)
	fmt.Println(chatData)
	if(chat.UserId==""){
		code=200
		db.Create(&chatData)
		return chatData,unSyncedChat,code
	}else{
		code=200
		return chat,unSyncedChat,code
	}
	
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

func GetChatDetail(userId string,date string,numOfChat int) ([]string,[]ChatDetail,string){
	db:=database.GetDB()
	chatDetail:=[]ChatDetail{}
	db.Limit(numOfChat).Where("actual_user_id=? AND created_at>?",userId,date).Find(&chatDetail)
	userId2:=make([]string,0)
	for _,c:=range chatDetail{
		u:=[]ChatDetail{}
		db.Where("chat_id=?",c.ChatId).Find(&u)
		fmt.Println(u)
		if(len(u)==0){
			userId2=append(userId2,"")
		}else if(u[0].ActualUserId==userId){
			userId2=append(userId2,u[1].ActualUserId)	
		}else{
			userId2=append(userId2,u[0].ActualUserId)
		}
		// fmt.Println(u)
		// if(len(u)!=1){
		// 	userId2=append(userId2,"")
		// }else{
		// 	userId2=append(userId2,u[0].ActualUserId)	
		// }
		
	}
	return userId2,chatDetail,"Ok"
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

// func GetLastChatIndex(chatId string) int{
// 	db:=database.GetDB()
// 	chat:=Chat{}
// 	db.Order("index desc").Where("chat_id=?",chatId).First(&chat)
// 	fmt.Println(chat)
// 	if chat.ChatId==""{
// 		return 0
// 	}
// 	return chat.Index
// }

func GetChatAfterTimeMessages(chatId string, numOfChat int, createdAt string)([]Chat){
	db:=database.GetDB()
	chat:=[]Chat{}
	db.Order("created_at desc").Limit(numOfChat).Where("chat_id=? AND created_at>?",chatId,createdAt).Find(&chat)
	return chat
}

func GetTempUserIdFromChatId(userId string,chatId string) string{
	db:=database.GetDB()
	anonymousUser:=AnonymousUserHelper{}
	db.Table("anonymous_users").Where("user_id=? AND chat_id=?",userId,chatId).Find(&anonymousUser)
	return anonymousUser.AnonymousId
}

func InsertMessageRequestDB(userId string,phone string,requestId string,messageType string,messageContent string,createdAt string) string{
	db:=database.GetDB()
	messageRequest:=MessageRequest{}
	db.Where("user_id=? AND phone=?",userId,phone).Find(&messageRequest)
	if(messageRequest.UserId==""){
		userPhone:=""
		db.Table("users").Select("phone").Where("user_id=?",userId).Find(&userPhone)
		messageRequest.UserId=userId
		messageRequest.Phone=phone
		messageRequest.RequestId=requestId
		messageRequest.MessageType=messageType
		messageRequest.MessageContent=messageContent
		messageRequest.CreatedAt=createdAt
		db.Create(&messageRequest)
		return createdAt
	}else{
		return messageRequest.CreatedAt
	}
}

func GetMessageRequestDB(userId string) []MessageRequestDS{
	db:=database.GetDB()
	phone:=""
	db.Table("users").Select("phone").Where("user_id=?",userId).Find(&phone)
	messageRequest:=[]MessageRequest{}
	db.Where("phone=?",phone).Find(&messageRequest)
	messageRequestDS:=make([]MessageRequestDS,0)
	for _,mr:=range messageRequest{
		MRTemp:=MessageRequestDS{}
		MRTemp.Phone=mr.UserPhone
		MRTemp.UserId=mr.UserId
		MRTemp.CreatedAt=mr.CreatedAt
		messageRequestDS=append(messageRequestDS,MRTemp)
	}
	return messageRequestDS
}