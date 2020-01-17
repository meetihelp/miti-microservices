package Social

import(
	database "miti-microservices/Database"
	"fmt"
	util "miti-microservices/Util"
)

func PoolStatusDB(userId string) PoolStatus{
	db:=database.GetDB()
	poolStatus:=PoolStatus{}
	err:=db.Where("user_id=?",userId).Find(&poolStatus).Error
	fmt.Print("PoolStatusDB:")
	fmt.Println(err)
	return poolStatus
}


func EnterInPooL(userId string,pincode string,createdAt string,gender string,sex string) PoolStatusHelper{
	db:=database.GetDB()
	poolWait:=PoolWaiting{}

	db.Where("user_id=?",userId).Find(&poolWait)
	if(poolWait.UserId==""){
		poolWait.UserId=userId
		poolWait.Pincode=pincode
		poolWait.CreatedAt=createdAt
		poolWait.Gender=gender
		poolWait.Sex=sex
		err:=db.Create(&poolWait).Error
		fmt.Print("EnterInPooL DB1:")
		fmt.Println(err)
	}

	poolStatus:=PoolStatus{}
	poolStatusResponse:=PoolStatusHelper{}
	db.Where("user_id=?",userId).Find(&poolStatus)
	if(poolStatus.UserId==""){
		poolStatus.UserId=userId
		poolStatus.Status="Waiting"
		poolStatus.CreatedAt=util.GetTime()
		err:=db.Create(&poolStatus).Error
		fmt.Print("EnterInPooL DB2:")
		fmt.Println(err)


	}

	poolStatusResponse.ChatId=poolStatus.ChatId
	poolStatusResponse.MatchUserId=poolStatus.MatchUserId
	poolStatusResponse.Status=poolStatus.Status
	return poolStatusResponse
}

func EnterInGroupPooL(userId string,pincode string,interest string,createdAt string,gender string,sex string) GroupPoolStatusHelper{
	db:=database.GetDB()
	poolWait:=GroupPoolWaiting{}
	poolWait.UserId=userId
	poolWait.Pincode=pincode
	poolWait.Interest=interest
	poolWait.CreatedAt=createdAt
	poolWait.Gender=gender
	poolWait.Sex=sex
	_=db.Create(&poolWait).Error

	groupPoolStatus:=GroupPoolStatus{}
	db.Where("user_id=? AND interest=?",userId,interest).Find(&groupPoolStatus)
	if(groupPoolStatus.UserId==""){
		groupPoolStatus.UserId=userId
		groupPoolStatus.Status="Waiting"
		groupPoolStatus.Interest=interest
	}
	groupPoolStatusHelper:=GroupPoolStatusHelper{}
	groupPoolStatusHelper.ChatId=groupPoolStatus.ChatId
	groupPoolStatusHelper.Status=groupPoolStatus.Status

	return groupPoolStatusHelper
}
func DeleteWaitPool(userId string) {
	db:=database.GetDB()
	db.Where("user_id=?",userId).Delete(&PoolWaiting{})
}

func DeleteWaitGroupPool(userId string) {
	db:=database.GetDB()
	db.Where("user_id=?",userId).Delete(&GroupPoolWaiting{})
}

func DeletePool(userId string,areaCode string,gender string){
	db:=database.GetDB()
	pool:=Pool{}
	db.Where("user_id=?",userId).Find(&pool)
	if(pool.UserId!=""){
		status:=EnterInPoolFromWait(areaCode,gender,1)
		Complementary_gender:="Male"
		if(gender=="Male"){
			Complementary_gender="Female"
		}
		if(status=="NA"){
			DeleteFromPoolHelper(areaCode,Complementary_gender,1)
		}
	}
}

func DeleteGroupPool(userId string,areaCode string,gender string){
	db:=database.GetDB()
	pool:=Pool{}
	db.Where("user_id=?",userId).Find(&pool)
	if(pool.UserId!=""){
		status:=EnterInGroupPoolFromWait(areaCode,gender,1)
		Complementary_gender:="Male"
		if(gender=="Male"){
			Complementary_gender="Female"
		}
		if(status=="NA"){
			DeleteFromGroupPoolHelper(areaCode,Complementary_gender,1)
		}
	}
}
func EnterInPoolFromWait(areaCode string,gender string,number_of_person int) string{
	return "Ok"
	//Code for checking if user is available to replace the person who cancelled the pooling
	//and replace the person
}

func DeleteFromPoolHelper(areaCode string,gender string,number_of_person int){
	//Delete a user from the pool table and put him/her back to wait pool table
}

func EnterInGroupPoolFromWait(areaCode string,gender string,number_of_person int) string{
	return "Ok"
}

func DeleteFromGroupPoolHelper(areaCode string,gender string,number_of_person int){

}


func GroupPoolStatusDB(userId string) []GroupPoolStatusHelper{
	db:=database.GetDB()
	groupPoolStatus:=[]GroupPoolStatus{}
	err:=db.Where("user_id=? AND interest=?",userId).Find(&groupPoolStatus).Error
	fmt.Print("PoolStatusDB:")
	fmt.Println(err)
	groupPoolStatusHelper:=[]GroupPoolStatusHelper{}
	for _,g:=range groupPoolStatusHelper{
		groupPoolStatusHelperTemp:=GroupPoolStatusHelper{}
		groupPoolStatusHelperTemp.ChatId=g.ChatId
		groupPoolStatusHelperTemp.Status=g.Status
		groupPoolStatusHelperTemp.Interest=g.Interest
		groupPoolStatusHelperTemp.CreatedAt=g.CreatedAt
		groupPoolStatusHelper=append(groupPoolStatusHelper,groupPoolStatusHelperTemp)
	}
	return groupPoolStatusHelper
}

func GetGroupAvailabilty(pincode string,interest string) (string,string){
	db:=database.GetDB()
	groupStats:=GroupStats{}
	db.Where("pincode=? AND interest=? AND number_of_temporary_member<?",pincode,interest,MAX_NUMBER_OF_TEMPORARY_MEMBER).Find(&groupStats)
	if(groupStats.ChatId==""){
		return "","None"
	}else{
		return groupStats.ChatId,"temporary"
	}
}

func InsertInGroup(chatId string,userId string,membership string,interest string) GroupPoolStatusHelper{
	db:=database.GetDB()
	createdAt:=util.GetTime()
	group:=Group{}
	group.UserId=userId
	group.ChatId=chatId
	group.Interest=interest
	group.CreatedAt=createdAt
	group.Membership=membership
	groupTemp:=Group{}
	db.Where("user_id=? AND chat_id=?",userId,chatId).Find(&groupTemp)
	if(groupTemp.UserId==""){
		db.Create(&group)

		chatDetails:=ChatDetail{}
		chatDetails.ActualUserId=userId
		chatDetails.ChatId=chatId
		chatDetails.CreatedAt=createdAt
		chatDetails.ChatType="Group"
		chatDetails.Name=util.GetGroupName(interest)
		db.Create(&chatDetails)
	}else{
		db.Table("chat_details").Where("user_id=? AND chat_id=?",userId,chatId).Update("chat_id=?",chatId)
		db.Table("group").Where("user_id=? AND chat_id=?",userId,chatId).Updates(group)
	}
	groupStats:=GroupStats{}
	db.Where("chat_id=?",chatId).Find(&groupStats)
	if(membership=="temporary"){
		numberOfTemporaryMember:=groupStats.NumberOfTemporaryMember+1
		db.Model(&groupStats).Where("chat_id=?",chatId).Update("number_of_temporary_member",numberOfTemporaryMember)
	}else if(membership=="permanent"){
		numberOfMember:=groupStats.NumberOfMember+1
		db.Model(&groupStats).Where("chat_id=?",chatId).Update("number_of_member",numberOfMember)
	}

	groupPoolStatus:=GroupPoolStatus{}
	groupPoolStatus.UserId=userId
	groupPoolStatus.ChatId=chatId
	groupPoolStatus.CreatedAt=createdAt
	groupPoolStatus.Interest=interest
	groupPoolStatus.Status=membership

	groupPoolStatusTemp:=GroupPoolStatus{}
	db.Where("user_id=? AND interest=?",userId,interest).Find(&groupPoolStatusTemp)
	if(groupPoolStatusTemp.UserId==""){
		db.Create(&groupPoolStatus)
	}else{
		db.Table("group_pool_status").Where("user_id=? AND interest=?",userId,interest).Updates(groupPoolStatus)
	}

	groupPoolStatusHelper:=GroupPoolStatusHelper{}
	groupPoolStatusHelper.ChatId=chatId
	groupPoolStatusHelper.Status=membership

	return groupPoolStatusHelper
}