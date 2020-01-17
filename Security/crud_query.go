package Security

import(
	database "miti-microservices/Database"
	// util "miti-microservices/Util"
)

func UpdatePrimaryTrustChain(userId string,chainId string,phone string,name string,requestId string,updatedAt string) string{
	db:=database.GetDB()
	primaryTrustChain:=PrimaryTrustChain{}
	db.Where("user_id=? AND chain_id=?",userId,chainId).Find(&primaryTrustChain)
	if(primaryTrustChain.Phone1==""){
		primaryTrustChain.Phone1=phone
		primaryTrustChain.Name1=name
	}else if(primaryTrustChain.Phone2==""){
		primaryTrustChain.Phone2=phone
		primaryTrustChain.Name2=name
	}else if(primaryTrustChain.Phone3==""){
		primaryTrustChain.Phone3=phone
		primaryTrustChain.Name3=name
	}else if(primaryTrustChain.Phone4==""){
		primaryTrustChain.Phone4=phone
		primaryTrustChain.Name4=name
	}else if(primaryTrustChain.Phone5==""){
		primaryTrustChain.Phone5=phone
		primaryTrustChain.Name5=name
	}else if(primaryTrustChain.Phone6==""){
		primaryTrustChain.Phone6=phone
		primaryTrustChain.Name6=name
	}
	if(primaryTrustChain.UserId==""){
		primaryTrustChain.UserId=userId
		primaryTrustChain.ChainId=chainId
		primaryTrustChain.RequestId=requestId
		primaryTrustChain.UpdatedAt=updatedAt
		db.Create(&primaryTrustChain)
		return updatedAt
	}else if(primaryTrustChain.RequestId==requestId){
		return primaryTrustChain.UpdatedAt
	}else{
		primaryTrustChain.UpdatedAt=updatedAt
		primaryTrustChain.RequestId=requestId
		db.Table("primary_trust_chains").Where("user_id=?",userId).Updates(primaryTrustChain)
		return updatedAt	
	}
	
}

func DeletePrimaryTrustChainDB(userId string,phoneId int,requestId string,updatedAt string) string{
	db:=database.GetDB()
	primaryTrustChain:=PrimaryTrustChain{}
	db.Where("user_id=?",userId).Find(&primaryTrustChain)

	if(phoneId==1){
		primaryTrustChain.Phone1=""
	}else if(phoneId==2){
		primaryTrustChain.Phone2=""
	}else if(phoneId==3){
		primaryTrustChain.Phone3=""
	}else if(phoneId==4){
		primaryTrustChain.Phone4=""
	}else if(phoneId==5){
		primaryTrustChain.Phone5=""
	}else if(phoneId==6){
		primaryTrustChain.Phone6=""
	}
	if(primaryTrustChain.UserId==""){
		primaryTrustChain.UserId=userId
		primaryTrustChain.RequestId=requestId
		primaryTrustChain.UpdatedAt=updatedAt
		db.Create(&primaryTrustChain)
		return updatedAt
	}else if(primaryTrustChain.RequestId==requestId){
		return primaryTrustChain.UpdatedAt
	}else{
		primaryTrustChain.UpdatedAt=updatedAt
		primaryTrustChain.RequestId=requestId
		db.Table("primary_trust_chain").Where("user_id=?",userId).Updates(primaryTrustChain)
		return updatedAt	
	}
}

func InsertSecondaryTrustChain(userId string,chatId string,requestId string,createdAt string) string{
	db:=database.GetDB()
	secondaryTrustChain:=SecondaryTrustChain{}
	db.Where("user_id=? AND chat_id=?",userId,chatId).Find(&secondaryTrustChain)
	if(secondaryTrustChain.UserId==""){
		secondaryTrustChain.UserId=userId
		secondaryTrustChain.ChatId=chatId
		secondaryTrustChain.CreatedAt=createdAt
		secondaryTrustChain.RequestId=requestId
		db.Create(&secondaryTrustChain)
		return createdAt
	}else{
		return secondaryTrustChain.CreatedAt
	}
}

func DeleteSecondaryTrustChainDB(userId string,chatId string){
	db:=database.GetDB()
	_=db.Where("user_id=? AND chat_id=?",userId,chatId).Delete(SecondaryTrustChain{}).Error
}