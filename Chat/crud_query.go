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
		user:=[]ChatDetail{}
		db.Where("chat_id=?",c.ChatId).Find(&user)
		// fmt.Println(u)
		for _,u:=range user{
			if(u.ActualUserId!=userId){
				userId2=append(userId2,u.ActualUserId)
			}
		}
		// if(len(u)==0){
		// 	userId2=append(userId2,"")
		// }else if(u[0].ActualUserId==userId){
		// 	userId2=append(userId2,u[1].ActualUserId)	
		// }else{
		// 	userId2=append(userId2,u[0].ActualUserId)
		// }
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

func InsertMessageRequestDB(userId string,senderName string,senderPhone string,phone string,requestId string,messageType string,messageContent string,createdAt string) string{
	db:=database.GetDB()
	messageRequest:=MessageRequest{}
	db.Where("sender_user_id=? AND phone=?",userId,phone).Find(&messageRequest)
	if(messageRequest.SenderUserId==""){
		messageRequest.SenderUserId=userId
		messageRequest.SenderName=senderName
		messageRequest.SenderPhone=senderPhone
		messageRequest.Phone=phone
		messageRequest.RequestId=requestId
		messageRequest.MessageId=util.GenerateToken()
		messageRequest.MessageType=messageType
		messageRequest.MessageContent=messageContent
		messageRequest.CreatedAt=createdAt
		messageRequest.Status="Wait"
		db.Create(&messageRequest)
		return createdAt
	}else{
		return messageRequest.CreatedAt
	}
}

func GetMessageRequestDB(userId string,createdAt string) []MessageRequestDS{
	db:=database.GetDB()
	phone:=GetUserPhone(userId)
	messageRequest:=make([]MessageRequest,0)
	db.Where("phone=? AND created_at>?",phone,createdAt).Find(&messageRequest)
	messageRequestDS:=make([]MessageRequestDS,0)
	for _,mr:=range messageRequest{
		MRTemp:=MessageRequestDS{}
		MRTemp.Phone=mr.SenderPhone
		MRTemp.Name=mr.SenderName
		// MRTemp.UserId=mr.UserId
		MRTemp.CreatedAt=mr.CreatedAt
		messageRequestDS=append(messageRequestDS,MRTemp)
	}
	return messageRequestDS
}

func GetUserPhone(userId string) string{
	db:=database.GetDB()
	user:=User{}
	db.Table("users").Select("phone").Where("user_id=?",userId).Find(&user)
	fmt.Println("GetuserPhone:->"+user.Phone)
	return user.Phone
}

func UpdateMessageRequestDB(phone string,senderPhone string,action string,actionRequestId string,updatedAt string) (string,string,MessageRequest){
	db:=database.GetDB()
	messageRequest:=MessageRequest{}
	db.Where("sender_phone=? AND phone=?",senderPhone,phone).Find(&messageRequest)
	if(messageRequest.SenderPhone==""){
		return "","",messageRequest
	}else if(messageRequest.ActionRequestId==actionRequestId){
		return messageRequest.SenderUserId,messageRequest.UpdatedAt,messageRequest
	}else{
		messageRequest.Status=action
		messageRequest.UpdatedAt=updatedAt
		messageRequest.ActionRequestId=actionRequestId
		db.Save(&messageRequest)
		return messageRequest.SenderUserId,messageRequest.UpdatedAt,messageRequest
	}
}

func InsertChatDetail(userId1 string,userId2 string,requestId string) (int,string){
	db:=database.GetDB()
	chatDetail:=ChatDetail{}
	chatDetail.CreatedAt=util.GetTime()
	chatDetail.ChatId=util.GenerateToken()
	chatDetail.ChatType="one-to-one"
	chatDetail.RequestId=requestId
	code:=200

	chatDetailTemp1:=ChatDetail{}
	db.Where("actual_user_id=? AND request_id=?",userId1,requestId).Find(&chatDetailTemp1)
	chatDetailTemp2:=ChatDetail{}
	db.Where("actual_user_id=? AND request_id=?",userId2,requestId).Find(&chatDetailTemp2)

	if(chatDetailTemp1.ChatId!=""){
		chatDetail.ChatId=chatDetailTemp1.ChatId
	}else if(chatDetailTemp2.ChatId!=""){
		chatDetail.ChatId=chatDetailTemp2.ChatId
	}


	if(chatDetailTemp1.ChatId==""){
		chatDetail.ActualUserId=userId1
		err:=db.Create(&chatDetail).Error
		if(err!=nil){
			fmt.Print("InsertChatDetail Error 1:")
			fmt.Println(err)
			code=1006
		}
	}
	if(chatDetailTemp2.ChatId==""){
		chatDetail.ActualUserId=userId2
		err:=db.Create(&chatDetail).Error
		if(err!=nil){
			fmt.Print("InsertChatDetail Error 2:")
			fmt.Println(err)
			code=1006
		}
	}
	

	return code,chatDetail.ChatId
}

func IsPhoneNumberExist(phone string) string{
	db:=database.GetDB()
	user:=User{}
	db.Table("users").Where("phone=?",phone).Find(&user)
	if(user.UserId==""){
		return "No"
	}
	return "Yes"
}

func InsertIntoChatFromMessageRequest(chatId string,requestId string,messageRequest MessageRequest) int{
	db:=database.GetDB()

	userId:=messageRequest.SenderUserId
	messageId:=messageRequest.MessageId
	messageContent:=messageRequest.MessageContent
	messageType:=messageRequest.MessageType
	createdAt:=messageRequest.CreatedAt
	chat:=Chat{}
	db.Where("user_id=? AND chat_id=? AND message_id=?",userId,chatId,messageId).Find(&chat)
	code:=200
	if(chat.ChatId==""){
		chat.UserId=userId
		chat.ChatId=chatId
		chat.RequestId=requestId
		chat.MessageId=messageId
		chat.MessageContent=messageContent
		chat.MessageType=messageType
		chat.CreatedAt=createdAt
		err:=db.Create(&chat).Error
		if(err!=nil){
			fmt.Print("InsertIntoChatFromMessageRequest Error :")
			fmt.Println(err)
			code=1006
		}

	}

	return code
}
// func UpdateMessageRequestDB(userid string,userPhone string,action string,actionRequestId string,updatedAt string)(string,string){
// 	db:=database.GetDB()
// 	messageRequest:=MessageRequest{}

// }