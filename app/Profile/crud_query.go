package Profile
import(
	"fmt"
   database "app/Database"
   "reflect"
)

func EnterProfileData(profileData Profile){
	fmt.Println("Enter_profile_data")
	db:=database.GetDB()
	db.Create(&profileData)
	questionResponse:=QuestionResponse{}
	questionResponse.UserId=profileData.UserId
	db.Create(&questionResponse)
}

func GetProfile(userId string) Profile{
	db:=database.GetDB()
	profile:=Profile{}
	db.Where("UserId=?",userId).First(&profile)
	return profile
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

func UpdateScore(userId string,score []int){
	db:=database.GetDB()
	profile:=Profile{}
	db.Where("UserId=?",userId).Find(&profile)
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


func InsertQuestionResponse(userId string,response map[string]int){
	db:=database.GetDB()

    keys := reflect.ValueOf(response).MapKeys()
	for _,key := range keys{
		db.Table("question_responses").Where("user_id=?",userId).Update(key.Interface().(string),response[key.Interface().(string)])
	}
}
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

