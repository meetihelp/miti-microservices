package Profile
import(
	"fmt"
   database "miti-microservices/Database"
   security "miti-microservices/Security"
   // util "miti-microservices/Util"
   // "reflect"
   
)

func GetProfileDB(userId string) Profile{
	db:=database.GetDB()
	profile:=Profile{}
	db.Where("user_id=?",userId).Find(&profile)
	return profile
}

func CheckIPIPStatus(userId string) int{
	db:=database.GetDB()
	user:=User{}
	db.Table("users").Where("user_id=?",userId).Find(&user)
	return user.IPIPStatus
}
func EnterProfileData(profileData Profile){
	fmt.Println("Enter_profile_data")
	db:=database.GetDB()
	tempProfile:=Profile{}
	db.Where("user_id=?",profileData.UserId).Find(&tempProfile)
	if(tempProfile.UserId==""){
		db.Create(&profileData)
		questionResponse:=QuestionResponse{}
		questionResponse.UserId=profileData.UserId
		db.Create(&questionResponse)
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
	}
}

// func GetProfile(userId string) Profile{
// 	db:=database.GetDB()
// 	profile:=Profile{}
// 	db.Where("UserId=?",userId).First(&profile)
// 	return profile
// }
func GetAuthorizedProfileDB(userId string) ProfileResponse{
	db:=database.GetDB()
	profile:=Profile{}
	db.Where("user_id=?",userId).First(&profile)
	profileResponse:=ProfileResponse{}
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
	db.Where("user_id=?",userId).Find(&interest)
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
	return profileResponse
}

func GetUnAuthorizedProfileDB(userId string) ProfileResponse{
	db:=database.GetDB()

	profileResponse:=ProfileResponse{}
	interest:=Interest{}
	db.Where("user_id=?",userId).Find(&interest)
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
	return profileResponse
}

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


func UpdateIPIPResponseDB(userId string,response map[string]int) int{
	db:=database.GetDB()
	questionResponse:=QuestionResponse{}
	db.Where("user_id=?",userId).Find(&questionResponse)
	flag:=1
	if(questionResponse.UserId==""){
		flag=0
	}
	ipipStatus,questionResponse:=getDataInQuestionResponseForm(questionResponse,response)
	if(flag==0){
		db.Create(&questionResponse)
	}else{
		db.Model(&questionResponse).Where("user_id=?",userId).Update(questionResponse)	
	}
	
	return ipipStatus
}

func UpdatePreferecePResponseDB(userId string,response PreferenceRequest) int{
	db:=database.GetDB()
	interest:=Interest{}
	db.Where("user_id=?",userId).Find(&interest)
	if(interest.UserId==""){
		interest.UserId=userId
		db.Create(&interest)

	}
	preferenceStatus,interest:=getDataInInterestForm(interest,response)
	db.Model(&interest).Where("user_id=?",userId).Update(interest)
	return preferenceStatus
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

func ProfileViewAuthorization(userId1 string,userId2 string) string{
	db:=database.GetDB()
	match:=Match{}
	db.Where("user_id1=? AND user_id2=?",userId1,userId2).First(&match)
	if (match.UserId1!="" && match.Like1=="Like" && match.Like2=="Like"){
		return "Ok"
	}
	db.Where("user_id1=? AND user_id2=?",userId2,userId1).First(&match)
	if (match.UserId1!="" && match.Like1=="Like" && match.Like2=="Like"){
		return "Ok"
	}
	return "Error"
}

func UpdateIPIPScore(userId string){
	db:=database.GetDB()
	questionResponse:=QuestionResponse{}
	db.Where("user_id=?",userId).Find(&questionResponse)
	score:=CalculateIPIPScore(questionResponse)
	UpdateScore(userId,score)
}

func CalculateIPIPScore(questionResponse QuestionResponse) ([] int){
	db:=database.GetDB()
	response:=ConvertQuestionResponseToArray(questionResponse)
	score:=make([]int,5)
	question:=[]Question{}
	db.Find(&question)
	for _,q:=range question{
		score[q.Type]=score[q.Type]+q.Factor*response[q.Id]
	}
	return score
}
func UpdateScore(userId string,score []int){
	db:=database.GetDB()
	profile:=Profile{}
	db.Where("user_id=?",userId).Find(&profile)
	score[0]=profile.Extraversion+score[0]
	score[1]=profile.Agreeableness+score[1]
	score[2]=profile.Conscientiousness+score[2]
	score[3]=profile.EmotionalStability+score[3]
	score[4]=profile.Intellect+score[4]
	db.Table("profiles").Where("UserId=?",userId).Updates(Profile{Extraversion:score[0],
		Agreeableness:score[1],Conscientiousness:score[2],EmotionalStability:score[3],Intellect:score[4]})
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
