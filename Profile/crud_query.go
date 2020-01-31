package Profile
import(
   database "miti-microservices/Database"
   security "miti-microservices/Security"
   "github.com/jinzhu/gorm"
   
)

func EnterProfileData(db *gorm.DB,profileData Profile) bool{
	tempProfile:=Profile{}
	err:=db.Where("user_id=?",profileData.UserId).Find(&tempProfile).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return true
	}
	if(tempProfile.UserId==""){
		err=db.Create(&profileData).Error
		if(err!=nil){
			return true
		}
		questionResponse:=QuestionResponse{}
		questionResponse.UserId=profileData.UserId
		err=db.Create(&questionResponse).Error
		if(err!=nil){
			return true
		}
	}else{
		if(profileData.Name!=""){
			tempProfile.Name=profileData.Name
		}
		if(profileData.DateOfBirth!=""){
			tempProfile.DateOfBirth=profileData.DateOfBirth
		}
		if(profileData.Job!=""){
			tempProfile.Job=profileData.Job
		}
		if(profileData.Gender!=""){
			tempProfile.Gender=profileData.Gender
		}
		if(profileData.Language!=""){
			tempProfile.Language=profileData.Language
		}
		if(profileData.Country!=""){
			tempProfile.Country=profileData.Country
		}
		if(profileData.Sex!=""){
			tempProfile.Sex=profileData.Sex
		}
		if(profileData.RelationshipStatus!=""){
			tempProfile.RelationshipStatus=profileData.RelationshipStatus
		}
		if(profileData.ParentsAddress!=""){
			tempProfile.ParentsAddress=profileData.ParentsAddress
		}

		err=db.Save(&tempProfile).Error
		if(err!=nil){
			return true
		}
	}

	return false
}

func UpdateIPIPResponseDB(db *gorm.DB,userId string,response map[string]int,page int) bool {
	questionResponse:=QuestionResponse{}
	err:=db.Where("user_id=?",userId).Find(&questionResponse).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return true
	}
	flag:=1
	if(questionResponse.UserId==""){
		flag=0
	}
	questionResponse=getDataInQuestionResponseForm(questionResponse,response,page)
	if(flag==0){
		err=db.Create(&questionResponse).Error
		if(err!=nil){
			return true
		}
	}else{
		err=db.Model(&questionResponse).Where("user_id=?",userId).Update(questionResponse).Error
		if(err!=nil){
			return true
		}
	}

	return false
	
}

func CalculateIPIPScore(db *gorm.DB,questionResponse QuestionResponse) ([] int,bool){
	response:=ConvertQuestionResponseToArray(questionResponse)
	score:=make([]int,5)
	question:=[]Question{}
	err:=db.Find(&question).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return score,true
	}
	for _,q:=range question{
		score[q.Type]=score[q.Type]+q.Factor*response[q.Id]
	}
	return score,false
}

func UpdateScore(db *gorm.DB,userId string,score []int) bool{
	profile:=Profile{}
	err:=db.Where("user_id=?",userId).Find(&profile).Error
	if(err!=nil){
		return true
	}
	score[0]=profile.Extraversion+score[0]
	score[1]=profile.Agreeableness+score[1]
	score[2]=profile.Conscientiousness+score[2]
	score[3]=profile.EmotionalStability+score[3]
	score[4]=profile.Intellect+score[4]
	err=db.Table("profiles").Where("user_id=?",userId).Updates(Profile{Extraversion:score[0],
		Agreeableness:score[1],Conscientiousness:score[2],EmotionalStability:score[3],Intellect:score[4]}).Error
	if(err!=nil){
		return true
	}
	return false
}


func UpdateIPIPScore(db *gorm.DB,userId string) bool{
	questionResponse:=QuestionResponse{}
	err:=db.Where("user_id=?",userId).Find(&questionResponse).Error
	if(err!=nil){
		return true
	}
	score,dbError:=CalculateIPIPScore(db,questionResponse)
	if(dbError){
		return true
	}
	dbError=UpdateScore(db,userId,score)
	return dbError
}

func UpdatePreferecePResponseDB(db *gorm.DB,userId string,response UpdatePreferenceRequest) (int,bool){
	interest:=Interest{}
	err:=db.Where("user_id=?",userId).Find(&interest).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return 0,true
	}
	if(interest.UserId==""){
		interest.UserId=userId
		err:=db.Create(&interest).Error
		if(err!=nil){
			return 0,true
		}

	}
	preferenceStatus,interest:=getDataInInterestForm(interest,response)
	err=db.Model(&interest).Where("user_id=?",userId).Update(interest).Error
	if(err!=nil){
		return 0,true
	}
	return preferenceStatus,false
}

func ProfileViewAuthorization(db *gorm.DB,userId1 string,userId2 string) (string,bool){
	match:=Match{}
	err:=db.Where("user_id1=? AND user_id2=?",userId1,userId2).First(&match).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return "Error",true
	}
	if (match.UserId1!="" && match.Like1=="Like" && match.Like2=="Like"){
		return "Authorized",false
	}
	err=db.Where("user_id1=? AND user_id2=?",userId2,userId1).First(&match).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return "Error",true
	}
	if (match.UserId1!="" && match.Like1=="Like" && match.Like2=="Like"){
		return "Authorized",false
	}

	//Check for UnAuthorized profile
	return "UnAuthorized",false
}

func GetUnAuthorizedProfileDB(db *gorm.DB,userId string) (ProfileResponse,bool){
	profileResponse:=ProfileResponse{}
	interest:=Interest{}
	err:=db.Where("user_id=?",userId).Find(&interest).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return profileResponse,true
	}
	profileResponse.InterestIndoorPassive1=interest.InterestIndoorPassive1
	profileResponse.InterestIndoorPassive2=interest.InterestIndoorPassive2
	profileResponse.InterestOutdoorPassive1=interest.InterestOutdoorPassive1
	profileResponse.InterestOutdoorPassive2=interest.InterestOutdoorPassive2
	profileResponse.InterestIndoorActive1=interest.InterestIndoorActive1
	profileResponse.InterestIndoorActive2=interest.InterestIndoorActive2
	profileResponse.InterestOutdoorActive1=interest.InterestOutdoorActive1
	profileResponse.InterestOutdoorActive2=interest.InterestOutdoorActive2
	profileResponse.InterestOthers1=interest.InterestOthers1
	profileResponse.InterestOthers2=interest.InterestOthers2
	profileResponse.InterestIdeology1=interest.InterestIdeology1
	profileResponse.InterestIdeology2=interest.InterestIdeology2
	return profileResponse,false
}

func GetAuthorizedProfileDB(db *gorm.DB,userId string) (ProfileResponse,bool){
	profile:=Profile{}
	profileResponse:=ProfileResponse{}
	err:=db.Where("user_id=?",userId).First(&profile).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return profileResponse,true
	}
	profileResponse.UserId=profile.UserId
	profileResponse.Name=profile.Name
	profileResponse.DateOfBirth=profile.DateOfBirth
	profileResponse.Job=profile.Job
	// profileResponse.ProfilePicURL=profile.ProfilePicURL
	profileResponse.ProfilePicID=profile.ProfilePicID
	profileResponse.Gender=profile.Gender
	profileResponse.Language=profile.Language
	profileResponse.Country=profile.Country
	profileResponse.Sex=profile.Sex
	profileResponse.RelationshipStatus=profile.RelationshipStatus


	interest:=Interest{}
	err=db.Where("user_id=?",userId).Find(&interest).Error
	if(err!=nil && !gorm.IsRecordNotFoundError(err)){
		return profileResponse,true
	}
	profileResponse.InterestIndoorPassive1=interest.InterestIndoorPassive1
	profileResponse.InterestIndoorPassive2=interest.InterestIndoorPassive2
	profileResponse.InterestOutdoorPassive1=interest.InterestOutdoorPassive1
	profileResponse.InterestOutdoorPassive2=interest.InterestOutdoorPassive2
	profileResponse.InterestIndoorActive1=interest.InterestIndoorActive1
	profileResponse.InterestIndoorActive2=interest.InterestIndoorActive2
	profileResponse.InterestOutdoorActive1=interest.InterestOutdoorActive1
	profileResponse.InterestOutdoorActive2=interest.InterestOutdoorActive2
	profileResponse.InterestOthers1=interest.InterestOthers1
	profileResponse.InterestOthers2=interest.InterestOthers2
	profileResponse.InterestIdeology1=interest.InterestIdeology1
	profileResponse.InterestIdeology2=interest.InterestIdeology2


	// profileResponse.ParentsAddress=profile.ParentsAddress
	// _=reflect.Copy(profileResponse,profile)
	return profileResponse,false
}

func InsertIntoMatch(db *gorm.DB,userId1 string,userId2 string) bool{
	match:=Match{}
	match.UserId1=userId1
	match.UserId2=userId2
	match.Like1="Like"
	match.Like2="Like"
	err:=db.Create(&match).Error
	if(err!=nil){
		return true
	}
	return false
}

func CheckIPIPStatus(db *gorm.DB,userId string) (int,bool){
	user:=User{}
	err:=db.Table("users").Where("user_id=?",userId).Find(&user).Error
	if(err!=nil){
		return 0,true
	}
	return user.IPIPStatus,false
}

func GetProfileDB(db *gorm.DB,userId string) (Profile,bool){
	profile:=Profile{}
	err:=db.Where("user_id=?",userId).Find(&profile).Error
	if(err!=nil){
		return profile,true
	}
	return profile,false
}
//LAST







// func GetProfile(userId string) Profile{
// 	db:=database.GetDB()
// 	profile:=Profile{}
// 	db.Where("UserId=?",userId).First(&profile)
// 	return profile
// }




func GetUserIdByName(Offset int,numOfProfile int,name string) ([]string){
	db:=database.GetDB()
	profile:=[]Profile{}
	db.Offset(Offset).Limit(numOfProfile).Where("Name=?",name).Find(&profile)
	userId:=make([]string,0)
	for _,data:=range profile{
		userId=append(userId,data.UserId)
	}
	return userId
}


func GetQuestionById(questionId int) string{
	db:=database.GetDB()
	question:=Question{}
	db.Where("QuestionId=?",questionId).Find(&question)
	return question.Content
}

func InsertQuestionInDB(content string,TypeofQuestion int,factor int){
	db:=database.GetDB()
	question:=Question{Content:content,Type:TypeofQuestion,Factor:factor}
	db.Create(&question)
}






// func UpdatePreferecePResponseDB(userId string,response map[string]string) int{
// 	db:=database.GetDB()
// 	interest:=Interest{}
// 	db.Where("user_id=?",userId).Find(&interest)
// 	if(interest.UserId==""){
// 		interest.UserId=userId
// 		db.Create(&interest)

// 	}
// 	preferenceStatus,interest:=getDataInInterestForm(interest,response)
// 	db.Model(&interest).Where("user_id=?",userId).Update(interest)
// 	return preferenceStatus
// }
func UpdateJob(userId string,job string){
	db:=database.GetDB()
	db.Table("profiles").Where("UserId=?",userId).Update("Job",job)
}

func UpdateName(userId string,name string){
	db:=database.GetDB()
	db.Table("profiles").Where("UserId=?",userId).Update("Name",name)
}

func UpdatePicURL(userId string,url string){
	db:=database.GetDB()
	db.Table("profiles").Where("UserId=?",userId).Update("PicUrl",url)
}
func UpdateDOB(userId string,dob string){
	db:=database.GetDB()
	db.Table("profiles").Where("UserId=?",userId).Update("DataOfBirth",dob)	
}

func UpdateLanguage(userId string,language string){
	db:=database.GetDB()
	db.Table("profiles").Where("UserId=?",userId).Update("Language",language)
}

func GetQuestionFromDB(offset int,numOfQuestion int)([]Question){
	db:=database.	GetDB()
	question:=[]Question{}
	db.Offset(offset).Limit(numOfQuestion).Find(&question)
	return question
}

func GetScore(response []Response) ([]int){
	db:=database.GetDB()
	score:=make([]int,5)
	for _,data:=range response{
		question:=Question{}
		db.Where("QuestionId=?",data.QuestionId).Find(&question)
		score[question.Factor]=score[question.Factor]+question.Type*data.Response
	}
	return score
}







func UpdateExtraversionScore(userId string,score int){
	db:=database.GetDB()
	profile:=Profile{}
	db.Where("UserId=?",userId).Find(&profile)
	newScore:=profile.Extraversion+score
	db.Table("profiles").Where("UserId=?",userId).Update("Extraversion",newScore)
}

func UpdateAgreeablenessScore(userId string,score int){
	db:=database.GetDB()
	profile:=Profile{}
	db.Where("UserId=?",userId).Find(&profile)
	newScore:=profile.Agreeableness+score
	db.Table("profiles").Where("UserId=?",userId).Update("Agreeableness",newScore)
}

func UpdateConscientiousnessScore(userId string,score int){
	db:=database.GetDB()
	profile:=Profile{}
	db.Where("UserId=?",userId).Find(&profile)
	newScore:=profile.Conscientiousness+score
	db.Table("profiles").Where("UserId=?",userId).Update("Conscientiousness",newScore)
}

func UpdateEmotionalStabilityScore(userId string,score int){
	db:=database.GetDB()
	profile:=Profile{}
	db.Where("UserId=?",userId).Find(&profile)
	newScore:=profile.EmotionalStability+score
	db.Table("profiles").Where("UserId=?",userId).Update("EmotionalStability",newScore)
}

func UpdateIntellectScore(userId string,score int){
	db:=database.GetDB()
	profile:=Profile{}
	db.Where("UserId=?",userId).Find(&profile)
	newScore:=profile.EmotionalStability+score
	db.Table("profiles").Where("UserId=?",userId).Update("Intellect",newScore)
}

func UpdateProfilePicDB(userId string,imageId string){
	db:=database.GetDB()
	db.Table("profiles").Where("user_id=?",userId).Update("profile_pic_id",imageId)
}

func UpdateMatchDB(userId1 string,profileReactionData ProfileReactionRequest){
	userId2:=profileReactionData.UserId
	reaction:=profileReactionData.Reaction
	db:=database.GetDB()
	_=db.Table("matches").Where("user_id1=? AND user_id2=?",userId1,userId2).Update("like1",reaction).Error
	_=db.Table("matches").Where("user_id1=? AND user_id2=?",userId2,userId1).Update("like2",reaction).Error

}

func EnterStatusDB(createStatusData Status) Status{
	db:=database.GetDB()
	statusResponse:=Status{}
	userId:=createStatusData.UserId
	requestId:=createStatusData.RequestId
	db.Table("status").Where("actual_user_id=? AND request_id=?",userId,requestId).Find(&statusResponse)
	if(statusResponse.UserId==""){
		db.Create(&createStatusData)
		return createStatusData
	}else{
		return statusResponse
	}
}

func IsChatIdOfUser(userId string,chatId string) string{
	db:=database.GetDB()
	chatDetail:=ChatDetail{}
	db.Table("chat_details").Where("user_id=? AND chat_id=?",userId,chatId).Find(&chatDetail)
	if(chatDetail.ActualUserId==""){
		return "Error"
	}else{
		return "Ok"
	}
}

func EnterCheckInterestDB(checkInterest CheckInterest){
	db:=database.GetDB()
	db.Create(&checkInterest)
}

func GetCheckInterestDB(userId string)([]GetCheckInterestList){
	db:=database.GetDB()
	checkInterestList:=make([]GetCheckInterestList,0)
	checkInterest:=[]CheckInterest{}
	db.Where("user_id2=?",userId).Find(&checkInterest)
	for _,c:=range checkInterest{
		checkInterest:=GetCheckInterestList{}
		checkInterest.UserId=c.UserId1
		checkInterest.Interest=c.Interest
		checkInterest.CreatedAt=c.CreatedAt
		checkInterestList=append(checkInterestList,checkInterest)
	}
	return checkInterestList
}



func GetStatusDB(chatId string) []StatusResponse{
	db:=database.GetDB()
	secondaryTrustChain:=[]security.SecondaryTrustChain{}
	db.Where("chat_id=?",chatId).Find(&secondaryTrustChain)
	statusList:=make([]StatusResponse,0)
	for _,chain:=range secondaryTrustChain{
		status:=Status{}
		userId:=chain.UserId
		db.Where("user_id=? AND active_status=?",userId,"active").Find(&status)
		statusResponseTemp:=StatusResponse{}
		if(status.UserId!=""){
			statusResponseTemp.UserId=status.UserId
			statusResponseTemp.CreatedAt=status.CreatedAt
			statusResponseTemp.StatusContent=status.StatusContent
			statusList=append(statusList,statusResponseTemp)
		}
	}
	return statusList
}

func GetUserInterest(userId string) []string{
	db:=database.GetDB()
	interest:=Interest{}
	db.Where("user_id=?",userId).Find(&interest)

	interestList:=make([]string,0)
	interestList=append(interestList,interest.InterestOutdoorPassive1)
	interestList=append(interestList,interest.InterestOutdoorPassive2)
	interestList=append(interestList,interest.InterestOutdoorActive1)
	interestList=append(interestList,interest.InterestOutdoorActive2)
	interestList=append(interestList,interest.InterestIndoorPassive1)
	interestList=append(interestList,interest.InterestIndoorPassive2)
	interestList=append(interestList,interest.InterestIndoorActive1)
	interestList=append(interestList,interest.InterestIndoorActive2)
	interestList=append(interestList,interest.InterestOthers1)
	interestList=append(interestList,interest.InterestOthers2)
	interestList=append(interestList,interest.InterestIdeology1)
	interestList=append(interestList,interest.InterestIdeology2)

	return interestList
}

func GetUserName(db *gorm.DB,userId string) (string,bool){
	profile:=Profile{}
	err:=db.Table("profiles").Where("user_id=?",userId).Find(&profile).Error
	if(gorm.IsRecordNotFoundError(err)){
		return profile.Name,false
	}
	if(err!=nil){
		return profile.Name,true
	}
	return profile.Name,false
}