package Chat

import(
	util "miti-microservices/Util"
	database "miti-microservices/Database"
	profile "miti-microservices/Profile"
	"github.com/jinzhu/gorm"
)

func GetChatDetail(db *gorm.DB,userId string,date string,numOfChat int) ([]string,[]ChatDetail,string,bool){
	// db:=database.GetDB()
	chatDetail:=[]ChatDetail{}
	userId2:=make([]string,0)
	err:=db.Limit(numOfChat).Where("actual_user_id=? AND created_at>?",userId,date).Find(&chatDetail).Error
	if(gorm.IsRecordNotFoundError(err)){
		return userId2,chatDetail,"Ok",false
	}else if(err!=nil){
		return userId2,chatDetail,"Error",true
	}
	
	for _,c:=range chatDetail{
		user:=[]ChatDetail{}
		err:=db.Where("chat_id=?",c.ChatId).Find(&user).Error
		if(gorm.IsRecordNotFoundError(err)){

		}else if(err!=nil){
			return userId2,chatDetail,"Error",true
		}
		flag:=0
		for _,u:=range user{
			if(u.ActualUserId!=userId && flag==0){
				userId2=append(userId2,u.ActualUserId)
				flag=1
			}
		}
		if(flag==0){
			userId2=append(userId2,"")
		}
		
	}
	return userId2,chatDetail,"Ok",false
}

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

func ChatInsertDB(db *gorm.DB,chatData Chat,lastUpdate string) (Chat,[]Chat,bool) {
	chat:=Chat{}
	unSyncedChat:=[]Chat{}
	chatId:=chatData.ChatId
	err:=db.Order("created_at desc").Where("chat_id=? AND created_at>?",chatId,lastUpdate).Find(&unSyncedChat).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return chat,unSyncedChat,true
	}
	err=db.Where("user_id=? AND request_id=?",chatData.UserId,chatData.RequestId).Find(&chat).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return chat,unSyncedChat,true
	}
	if(chat.UserId==""){
		err=db.Create(&chatData).Error
		if(err!=nil){
			return chatData,unSyncedChat,true
		}else{
			return chatData,unSyncedChat,false
		}
		
	}else{
		return chatData,unSyncedChat,false
	}
	
}
func CheckCorrectChat(db *gorm.DB,userId string,chatId string) (string,bool){
	// db:=database.GetDB()
	chatDetail:=ChatDetail{}
	err:=db.Where("actual_user_id=? AND chat_id=?",userId,chatId).First(&chatDetail).Error
	if(gorm.IsRecordNotFoundError(err)){
		return "Error",false
	}else if(err!=nil){
		return "Error",true
	}
	return "Ok",false
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

func UpdateChatTime(db *gorm.DB,chatId string, lastUpdate string) bool{
	// db:=database.GetDB()
	err:=db.Table("chat_details").Where("chat_id=?",chatId).Update("last_update",lastUpdate).Error
	if(err!=nil){
		return true
	}

	return false
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

func GetChatAfterTimeMessages(db *gorm.DB,chatId string, numOfChat int, createdAt string)([]Chat,bool){
	// db:=database.GetDB()
	chat:=[]Chat{}
	err:=db.Order("created_at desc").Limit(numOfChat).Where("chat_id=? AND created_at>?",chatId,createdAt).Find(&chat).Error
	if(gorm.IsRecordNotFoundError(err)){
		return chat,false
	}else if(err!=nil){
		return chat,true
	}
	return chat,false
}

func GetTempUserIdFromChatId(userId string,chatId string) string{
	db:=database.GetDB()
	anonymousUser:=AnonymousUserHelper{}
	db.Table("anonymous_users").Where("user_id=? AND chat_id=?",userId,chatId).Find(&anonymousUser)
	return anonymousUser.AnonymousId
}

func InsertMessageRequestDB(db *gorm.DB,userId string,senderName string,senderPhone string,phone string,requestId string,messageType string,messageContent string,createdAt string) (string,bool){
	messageRequest:=MessageRequest{}
	err:=db.Where("sender_user_id=? AND phone=?",userId,phone).Find(&messageRequest).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return "",true
	}
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
		err:=db.Create(&messageRequest).Error
		if(err!=nil){
			return "",true
		}
		return createdAt,false
	}else{
		return messageRequest.CreatedAt,false
	}
}

func GetMessageRequestDB(db *gorm.DB,userId string,createdAt string) ([]MessageRequestDS,bool){
	messageRequest:=make([]MessageRequest,0)
	messageRequestDS:=make([]MessageRequestDS,0)
	phone,dbError:=GetUserPhone(db,userId)
	if(dbError==true){
		return messageRequestDS,true
	}
	
	err:=db.Where("phone=? AND created_at>?",phone,createdAt).Find(&messageRequest).Error
	if(gorm.IsRecordNotFoundError(err)){
		return messageRequestDS,false
	}
	if(err!=nil){
		return messageRequestDS,true
	}
	
	for _,mr:=range messageRequest{
		MRTemp:=MessageRequestDS{}
		MRTemp.Phone=mr.SenderPhone
		MRTemp.Name=mr.SenderName
		// MRTemp.UserId=mr.UserId
		MRTemp.CreatedAt=mr.CreatedAt
		messageRequestDS=append(messageRequestDS,MRTemp)
	}
	return messageRequestDS,false
}

func GetUserPhone(db *gorm.DB,userId string) (string,bool){
	user:=User{}
	err:=db.Table("users").Select("phone").Where("user_id=?",userId).Find(&user).Error
	if(gorm.IsRecordNotFoundError(err)){
		return user.Phone,false
	}
	if(err!=nil){
		return user.Phone,true
	}
	return user.Phone,false
}

func UpdateMessageRequestDB(db *gorm.DB,phone string,senderPhone string,action string,actionRequestId string,updatedAt string) (string,string,MessageRequest,bool){
	messageRequest:=MessageRequest{}
	err:=db.Where("sender_phone=? AND phone=?",senderPhone,phone).Find(&messageRequest).Error
	if(gorm.IsRecordNotFoundError(err)){
		return "","",messageRequest,false
	}
	if(err!=nil){
		return "","",messageRequest,true
	}

	if(messageRequest.ActionRequestId==actionRequestId){
		return messageRequest.SenderUserId,messageRequest.UpdatedAt,messageRequest,false
	}else{
		messageRequest.Status=action
		messageRequest.UpdatedAt=updatedAt
		messageRequest.ActionRequestId=actionRequestId
		err:=db.Save(&messageRequest).Error
		if(err!=nil){
			return "","",messageRequest,true
		}
		return messageRequest.SenderUserId,messageRequest.UpdatedAt,messageRequest,false
	}
}

func InsertChatDetailMessageRequest(db *gorm.DB,userId1 string,userId2 string,requestId string) (string,bool){
	chatDetail:=ChatDetail{}
	chatDetail.CreatedAt=util.GetTime()
	chatDetail.ChatId=util.GenerateToken()
	chatDetail.ChatType="one-to-one"
	chatDetail.RequestId=requestId

	chatDetailTemp1:=ChatDetail{}
	err:=db.Where("actual_user_id=? AND request_id=?",userId1,requestId).Find(&chatDetailTemp1).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return "",true
	}
	chatDetailTemp2:=ChatDetail{}
	err=db.Where("actual_user_id=? AND request_id=?",userId2,requestId).Find(&chatDetailTemp2).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return "",true
	}

	if(chatDetailTemp1.ChatId!=""){
		chatDetail.ChatId=chatDetailTemp1.ChatId
	}else if(chatDetailTemp2.ChatId!=""){
		chatDetail.ChatId=chatDetailTemp2.ChatId
	}


	if(chatDetailTemp1.ChatId==""){
		chatDetail.ActualUserId=userId1
		name,dbError:=profile.GetUserName(db,userId2)
		chatDetail.Name=name
		if(!dbError){
			err:=db.Create(&chatDetail).Error
			if(err!=nil){
				return "",true
			}	
		}else{
			return "",true
		}
		
	}
	if(chatDetailTemp2.ChatId==""){
		chatDetail.ActualUserId=userId2
		name,dbError:=profile.GetUserName(db,userId1)
		chatDetail.Name=name
		if(!dbError){
			err:=db.Create(&chatDetail).Error
			if(err!=nil){
				return "",true
			}	
		}else{
			return "",true
		}
	}
	

	return chatDetail.ChatId,false
}

func IsPhoneNumberExist(db *gorm.DB,phone string) (string,bool){
	user:=User{}
	err:=db.Table("users").Where("phone=?",phone).Find(&user).Error
	if(gorm.IsRecordNotFoundError(err)){
		return "No",false
	}
	if(err!=nil){
		return "No",true
	}
	return "Yes",false
}

func InsertIntoChatFromMessageRequest(db *gorm.DB,chatId string,requestId string,messageRequest MessageRequest) bool{
	userId:=messageRequest.SenderUserId
	messageId:=messageRequest.MessageId
	messageContent:=messageRequest.MessageContent
	messageType:=messageRequest.MessageType
	createdAt:=messageRequest.CreatedAt
	chat:=Chat{}
	err:=db.Where("user_id=? AND chat_id=? AND message_id=?",userId,chatId,messageId).Find(&chat).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return true
	}
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
			return true
		}

	}

	return false
}
// func UpdateMessageRequestDB(userid string,userPhone string,action string,actionRequestId string,updatedAt string)(string,string){
// 	db:=database.GetDB()
// 	messageRequest:=MessageRequest{}

// }

