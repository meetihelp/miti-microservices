package Social

import(
	database "miti-microservices/Database"
	profile "miti-microservices/Profile"
	chat "miti-microservices/Chat"
	// "fmt"
	util "miti-microservices/Util"
	"github.com/jinzhu/gorm"
)

func PoolStatusDB(db *gorm.DB,userId string) (PoolStatus,bool){
	poolStatus:=PoolStatus{}
	err:=db.Where("user_id=?",userId).Find(&poolStatus).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return poolStatus,true
	}
	return poolStatus,false
}


func EnterInPooL(db *gorm.DB,userId string,pincode string,createdAt string,gender string,sex string) (PoolStatusHelper,bool){
	poolWait:=PoolWaiting{}
	poolStatus:=PoolStatus{}
	poolStatusResponse:=PoolStatusHelper{}

	err:=db.Where("user_id=?",userId).Find(&poolWait).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return poolStatusResponse,true
	}
	if(poolWait.UserId==""){
		poolWait.UserId=userId
		poolWait.Pincode=pincode
		poolWait.CreatedAt=createdAt
		poolWait.Gender=gender
		poolWait.Sex=sex
		err:=db.Create(&poolWait).Error
		if(err!=nil){
			return poolStatusResponse,true
		}
	}

	err=db.Where("user_id=?",userId).Find(&poolStatus).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return poolStatusResponse,true
	}
	if(poolStatus.UserId==""){
		poolStatus.UserId=userId
		poolStatus.Status="Waiting"
		poolStatus.CreatedAt=util.GetTime()
		err:=db.Create(&poolStatus).Error
		if(err!=nil){
			return poolStatusResponse,true
		}

	}
	
	poolStatusResponse.ChatId=poolStatus.ChatId
	poolStatusResponse.MatchUserId=poolStatus.MatchUserId
	poolStatusResponse.Status=poolStatus.Status
	return poolStatusResponse,false
}

func EnterInGroupPooL(db *gorm.DB,userId string,pincode string,interest string,createdAt string,gender string,sex string) (GroupPoolStatusHelper,bool){
	groupPoolStatusHelper:=GroupPoolStatusHelper{}
	poolWait:=GroupPoolWaiting{}
	poolWait.UserId=userId
	poolWait.Pincode=pincode
	poolWait.Interest=interest
	poolWait.CreatedAt=createdAt
	poolWait.Gender=gender
	poolWait.Sex=sex
	err:=db.Create(&poolWait).Error
	if(err!=nil){
		return groupPoolStatusHelper,true
	}

	groupPoolStatus:=GroupPoolStatus{}
	err=db.Where("user_id=? AND interest=?",userId,interest).Find(&groupPoolStatus).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return groupPoolStatusHelper,true
	}
	if(groupPoolStatus.UserId==""){
		groupPoolStatus.UserId=userId
		groupPoolStatus.Status="Waiting"
		groupPoolStatus.Interest=interest
	}
	groupPoolStatusHelper.ChatId=groupPoolStatus.ChatId
	groupPoolStatusHelper.Status=groupPoolStatus.Status

	return groupPoolStatusHelper,false
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


func GroupPoolStatusDB(db *gorm.DB,userId string) ([]string,[]GroupPoolStatusHelper,bool){
	groupPoolStatus:=[]GroupPoolStatus{}
	groupPoolStatusHelper:=[]GroupPoolStatusHelper{}
	var interest []string
	err:=db.Where("user_id=?",userId).Find(&groupPoolStatus).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return interest,groupPoolStatusHelper,true
	}
	
	for _,g:=range groupPoolStatusHelper{
		groupPoolStatusHelperTemp:=GroupPoolStatusHelper{}
		groupPoolStatusHelperTemp.ChatId=g.ChatId
		groupPoolStatusHelperTemp.Status=g.Status
		groupPoolStatusHelperTemp.Interest=g.Interest
		groupPoolStatusHelperTemp.CreatedAt=g.CreatedAt
		groupPoolStatusHelper=append(groupPoolStatusHelper,groupPoolStatusHelperTemp)
	}

	interest,dbError:=profile.GetUserInterest(db,userId)

	return interest,groupPoolStatusHelper,dbError
}

func GetGroupAvailabilty(db *gorm.DB,userId string,pincode string,interest string,requestId string) (string,string,bool){
	groupStats:=GroupStats{}
	group:=Group{}
	err:=db.Where("user_id=? AND pincode=? AND interest=? AND request_id=?",userId,pincode,interest,requestId).Find(&group).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return "","",true
	}
	if(group.UserId!=""){
		return "already",group.ChatId,false
	}
	err=db.Where("pincode=? AND interest=? AND number_of_member<?",pincode,interest,MAX_NUMBER_OF_MEMBER).Find(&groupStats).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return "","",true
	}
	if(groupStats.ChatId!=""){
		return groupStats.ChatId,"permanent",false
	}
	err=db.Where("pincode=? AND interest=? AND number_of_temporary_member<?",pincode,interest,MAX_NUMBER_OF_TEMPORARY_MEMBER).Find(&groupStats).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return "","",true
	}
	if(groupStats.ChatId==""){
		chatId,status,dbError:=CreateNewGroup(db,pincode,interest)
		if(!dbError){
			if(chatId==""){
				return "","None",false
			}else {
				return chatId,status,false
			}
		}else{
			return "","",true
		}
			
	}else{
		return groupStats.ChatId,"temporary",false
	}
}

func CreateNewGroup(db *gorm.DB,pincode string,interest string) (string,string,bool){
	createdAt:=util.GetTime()
	chatDetail:=ChatDetail{}
	chatDetail.ChatId=util.GenerateToken()
	chatDetail.ChatType="group"
	chatDetail.CreatedAt=createdAt
	chatDetail.Name=util.GetGroupName(interest)

	groupStats:=GroupStats{}
	err:=db.Where("pincode=? AND interest=? AND number_of_temporary_member=?",pincode,interest,MAX_NUMBER_OF_TEMPORARY_MEMBER).Find(&groupStats).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return "","",true
	}
	group:=[]Group{}
	chatId:=groupStats.ChatId
	err=db.Where("pincode=? AND interest=? AND chat_id=?",pincode,interest,chatId).Find(&group).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return "","",true
	}
	err=db.Table("group_stats").Where("chat_id=?",chatId).Update("number_of_temporary_member",0).Error
	if(err!=nil){
		return "","",true
	}
	count:=0


	for _,member:=range group{
		userId:=member.UserId
		chatId:=member.ChatId
		err=db.Table("group").Where("user_id=? AND interest=?",userId,interest).Updates(Group{ChatId:chatDetail.ChatId,Membership:"permanent",CreatedAt:createdAt}).Error
		if(err!=nil){
			return "","",true
		}
		err=db.Where("actual_user_id=? AND chat_id=?",userId,chatId).Delete(&ChatDetail{}).Error
		if(err!=nil){
			return "","",true
		}
		chatDetail.ActualUserId=userId
		err=db.Create(&chatDetail).Error
		if(err!=nil){
			return "","",true
		}
		count++;
	}

	newGroupStats:=GroupStats{}
	newGroupStats.ChatId=chatDetail.ChatId
	newGroupStats.NumberOfMember=count
	newGroupStats.NumberOfTemporaryMember=0
	newGroupStats.Interest=interest
	newGroupStats.Pincode=pincode
	newGroupStats.CreatedAt=createdAt
	err=db.Create(&newGroupStats).Error
	if(err!=nil){
		return "","",true
	}


	return chatDetail.ChatId,"permanent",false
}



func InsertInGroup(db *gorm.DB,chatId string,pincode string,userId string,membership string,interest string,requestId string) (GroupPoolStatusHelper,bool){
	groupPoolStatusHelper:=GroupPoolStatusHelper{}
	createdAt:=util.GetTime()
	group:=Group{}
	group.UserId=userId
	group.ChatId=chatId
	group.Interest=interest
	group.CreatedAt=createdAt
	group.Pincode=pincode
	group.RequestId=requestId
	group.Membership=membership
	groupTemp:=Group{}
	err:=db.Where("user_id=? AND interest=?",userId,interest).Find(&groupTemp).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return groupPoolStatusHelper,true
	}
	if(groupTemp.UserId==""){
		err=db.Create(&group).Error
		if(err!=nil){
			return groupPoolStatusHelper,true
		}

		chatDetails:=ChatDetail{}
		chatDetails.ActualUserId=userId
		chatDetails.ChatId=chatId
		chatDetails.CreatedAt=createdAt
		chatDetails.ChatType="Group"
		chatDetails.Name=util.GetGroupName(interest)
		err=db.Create(&chatDetails).Error
		if(err!=nil){
			return groupPoolStatusHelper,true
		}
	}else{
		err=db.Table("chat_details").Where("actual_user_id=? AND chat_id=?",userId,groupTemp.ChatId).Updates(chat.ChatDetail{ChatId:chatId,CreatedAt:createdAt}).Error
		if(err!=nil){
			return groupPoolStatusHelper,true
		}
		err=db.Table("groups").Where("user_id=? AND chat_id=?",userId,groupTemp.ChatId).Updates(group).Error
		if(err!=nil){
			return groupPoolStatusHelper,true
		}
	}
	groupStats:=GroupStats{}
	db.Where("chat_id=?",chatId).Find(&groupStats)
	if(membership=="temporary"){
		numberOfTemporaryMember:=groupStats.NumberOfTemporaryMember+1
		err=db.Model(&groupStats).Where("chat_id=?",chatId).Update("number_of_temporary_member",numberOfTemporaryMember).Error
		if(err!=nil){
			return groupPoolStatusHelper,true
		}
	}else if(membership=="permanent"){
		numberOfMember:=groupStats.NumberOfMember+1
		err=db.Model(&groupStats).Where("chat_id=?",chatId).Update("number_of_member",numberOfMember).Error
		if(err!=nil){
			return groupPoolStatusHelper,true
		}
	}

	groupPoolStatus:=GroupPoolStatus{}
	groupPoolStatus.UserId=userId
	groupPoolStatus.ChatId=chatId
	groupPoolStatus.CreatedAt=createdAt
	groupPoolStatus.Interest=interest
	groupPoolStatus.Status=membership

	groupPoolStatusTemp:=GroupPoolStatus{}
	err=db.Where("user_id=? AND interest=?",userId,interest).Find(&groupPoolStatusTemp).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return groupPoolStatusHelper,true
	}
	if(groupPoolStatusTemp.UserId==""){
		err=db.Create(&groupPoolStatus).Error
		if(err!=nil){
			return groupPoolStatusHelper,true
		}
	}else{
		err=db.Table("group_pool_statuses").Where("user_id=? AND interest=?",userId,interest).Updates(groupPoolStatus).Error
		if(err!=nil){
			return groupPoolStatusHelper,true
		}
	}

	
	groupPoolStatusHelper.ChatId=chatId
	groupPoolStatusHelper.Status=membership

	return groupPoolStatusHelper,false
}

func GetGroupPoolStatus(db *gorm.DB,userId string,pincode string,interest string) (GroupPoolStatusHelper,bool){
	group:=Group{}
	groupPoolStatusHelper:=GroupPoolStatusHelper{}
	err:=db.Where("user_id=? AND pincode=? AND interest=?",userId,pincode,interest).Find(&group).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return groupPoolStatusHelper,true
	}
	
	groupPoolStatusHelper.ChatId=group.ChatId
	groupPoolStatusHelper.Status=group.Membership
	groupPoolStatusHelper.Interest=group.Interest
	groupPoolStatusHelper.CreatedAt=group.CreatedAt
	return groupPoolStatusHelper,false
}