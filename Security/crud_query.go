package Security

import(
	database "miti-microservices/Database"
	"fmt"
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

func DeletePrimaryTrustChainDB(userId string,chainId string,phone string,requestId string,updatedAt string) string{
	db:=database.GetDB()
	primaryTrustChain:=PrimaryTrustChain{}
	db.Where("user_id=? AND chain_id=?",userId,chainId).Find(&primaryTrustChain)

	if(primaryTrustChain.Phone1==phone){
		fmt.Println("Deleting Phone1")
		primaryTrustChain.Phone1=""
		primaryTrustChain.Name1=""
	}else if(primaryTrustChain.Phone2==phone){
		fmt.Println("Deleting Phone2")
		primaryTrustChain.Phone2=""
		primaryTrustChain.Name2=""
	}else if(primaryTrustChain.Phone3==phone){
		fmt.Println("Deleting Phone3")
		primaryTrustChain.Phone3=""
		primaryTrustChain.Name3=""
	}else if(primaryTrustChain.Phone4==phone){
		fmt.Println("Deleting Phone4")
		primaryTrustChain.Phone4=""
		primaryTrustChain.Name4=""
	}else if(primaryTrustChain.Phone5==phone){
		fmt.Println("Deleting Phone5")
		primaryTrustChain.Phone5=""
		primaryTrustChain.Name5=""
	}else if(primaryTrustChain.Phone6==phone){
		fmt.Println("Deleting Phone6")
		primaryTrustChain.Phone6=""
		primaryTrustChain.Name6=""
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
		fmt.Print("Updating Primary Trust Chain:->")
		fmt.Println(primaryTrustChain)
		// db.Table("primary_trust_chains").Where("user_id=? AND chain_id=?",userId,chainId).Updates(&primaryTrustChain)
		db.Save(&primaryTrustChain)
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


func GetPrimaryTrustPhoneList(userId string) []string{
	db:=database.GetDB()
	primaryTrustChain:=PrimaryTrustChain{}
	db.Where("user_id=?",userId).Find(&primaryTrustChain)

	phoneList:=make([]string,0)
	if(primaryTrustChain.Phone1!=""){
		phoneList=append(phoneList,primaryTrustChain.Phone1)
	}
	if(primaryTrustChain.Phone2!=""){
		phoneList=append(phoneList,primaryTrustChain.Phone2)
	}
	if(primaryTrustChain.Phone3!=""){
		phoneList=append(phoneList,primaryTrustChain.Phone3)
	}
	if(primaryTrustChain.Phone4!=""){
		phoneList=append(phoneList,primaryTrustChain.Phone4)
	}
	if(primaryTrustChain.Phone5!=""){
		phoneList=append(phoneList,primaryTrustChain.Phone5)
	}
	if(primaryTrustChain.Phone6!=""){
		phoneList=append(phoneList,primaryTrustChain.Phone6)
	}

	return phoneList
}

func GetUserName(userId string) string{
	db:=database.GetDB()
	profile:=Profile{}
	db.Table("profiles").Where("user_id=?",userId).Find(&profile)
	return profile.Name
}

func GetPrimaryTrustChainDB(userId string) ([]PrimaryTrustChainList,string,string,string){
	db:=database.GetDB()
	primaryTrustChain:=PrimaryTrustChain{}
	db.Where("user_id=?",userId).Find(&primaryTrustChain)
	chainName:=primaryTrustChain.ChainName
	chainId:=primaryTrustChain.ChainId
	updatedAt:=primaryTrustChain.UpdatedAt
	primaryTrustChainList:=make([]PrimaryTrustChainList,0)
	primaryTrustChainListTemp:=PrimaryTrustChainList{}
	if(primaryTrustChain.Phone1!=""){
		primaryTrustChainListTemp.Phone=primaryTrustChain.Phone1
		primaryTrustChainListTemp.Name=primaryTrustChain.Name1
		primaryTrustChainList=append(primaryTrustChainList,primaryTrustChainListTemp)
	}
	if(primaryTrustChain.Phone2!=""){
		primaryTrustChainListTemp.Phone=primaryTrustChain.Phone2
		primaryTrustChainListTemp.Name=primaryTrustChain.Name2
		primaryTrustChainList=append(primaryTrustChainList,primaryTrustChainListTemp)
	}
	if(primaryTrustChain.Phone3!=""){
		primaryTrustChainListTemp.Phone=primaryTrustChain.Phone3
		primaryTrustChainListTemp.Name=primaryTrustChain.Name3
		primaryTrustChainList=append(primaryTrustChainList,primaryTrustChainListTemp)
	}
	if(primaryTrustChain.Phone4!=""){
		primaryTrustChainListTemp.Phone=primaryTrustChain.Phone4
		primaryTrustChainListTemp.Name=primaryTrustChain.Name4
		primaryTrustChainList=append(primaryTrustChainList,primaryTrustChainListTemp)
	}
	if(primaryTrustChain.Phone5!=""){
		primaryTrustChainListTemp.Phone=primaryTrustChain.Phone5
		primaryTrustChainListTemp.Name=primaryTrustChain.Name5
		primaryTrustChainList=append(primaryTrustChainList,primaryTrustChainListTemp)
	}
	if(primaryTrustChain.Phone6!=""){
		primaryTrustChainListTemp.Phone=primaryTrustChain.Phone6
		primaryTrustChainListTemp.Name=primaryTrustChain.Name6
		primaryTrustChainList=append(primaryTrustChainList,primaryTrustChainListTemp)
	}

	return primaryTrustChainList,chainName,chainId,updatedAt
}