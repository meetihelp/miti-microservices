package Profile

import(
	"fmt"
	"net/http"
	// "log"
	"io/ioutil"
	// "strings"
	"encoding/json"
   util "app/Util"
)


func ProfileCreation(w http.ResponseWriter, r *http.Request){
	header:=ProfileCreationHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie

	userId,dErr:=util.GetUserIdFromSession(sessionId)
	if dErr=="Error"{
		fmt.Println("Session Does not exist")
		util.Message(w,1003)
		return
	}


	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}
	profileData:=Profile{}
	errProfileData:=json.Unmarshal(requestBody,&profileData)
	if errProfileData!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}
	fmt.Println(profileData.Name)
	profileData.UserId=userId
	sanatizationStatus:=Sanatize(profileData)
	if sanatizationStatus =="Error"{
		fmt.Println("profile creation data invalid")
		util.Message(w,1002)
		return
	}
	// profile_data_handle(w,profile_data)
	EnterProfileData(profileData)

	util.Message(w,200)

}

// func profile_data_handle(w http.ResponseWriter,profile_data Profile){
// 	Enter_profile_data(profile_data)
// 	util.Message(w,200)
// 	return
// }

func GetQuestion(w http.ResponseWriter, r *http.Request){
	header:=ProfileCreationHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie

	_,dErr:=util.GetUserIdFromSession(sessionId)
	if dErr=="Error"{
		fmt.Println("Session Does not exist")
		util.Message(w,1003)
		return
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}

	questionRequest:=QuestionRequest{}
	errQuestionData:=json.Unmarshal(requestBody,&questionRequest)
	if errQuestionData!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}
	// sanatization_status:=Sanatize(questionRequest)
	// if sanatization_status =="ERROR"{
	// 	fmt.Println("profile creation data invalid")
	// 	util.Message(w,1002)
	// 	return
	// }
	question:=GetQuestionFromDB(questionRequest.Offset,questionRequest.NumOfQuestion)

	SendQuestion(w,question)
}


func UpdateQuestionResponse(w http.ResponseWriter, r *http.Request){
	header:=UpdateQuestionResponseHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie

	userId,dErr:=util.GetUserIdFromSession(sessionId)
	if dErr=="Error"{
		fmt.Println("Session Does not exist")
		util.Message(w,1003)
		return
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}
	
	var data map[string]int
	if err := json.Unmarshal(requestBody, &data); err != nil {
        panic(err)
    }
    fmt.Println(data)
	// responseWrapper:=ResponseWrapper{}
	// err_question_data:=json.Unmarshal(requestBody,&responseWrapper)
	// fmt.Println(responseWrapper.IPIP)
	// response:=[]Response{}
	// response=responseWrapper.IPIP
	// fmt.Println(response[0].QuestionId)
	// if err_question_data!=nil{
	// 	fmt.Println("Could not Unmarshall profile data")
	// 	util.Message(w,1001)
	// 	return
	// }

	// sanatization_status:=Sanatize(questionResponse)
	// if sanatization_status =="ERROR"{
	// 	fmt.Println("profile creation data invalid")
	// 	util.Message(w,1002)
	// 	return
	// }

	InsertQuestionResponse(userId,data)

	// score:=GetScore(responseWrapper.IPIP)
	// UpdateScore(userId,score)
	util.Message(w,200)
}

func InsertQuestion(w http.ResponseWriter, r *http.Request){
	header:=InsertQuestionHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie

	_,dErr:=util.GetUserIdFromSession(sessionId)
	if dErr=="Error"{
		fmt.Println("Session Does not exist")
		util.Message(w,1003)
		return
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}

	question:=Question{}
	errQuestionData:=json.Unmarshal(requestBody,&question)
	if errQuestionData!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}
	// sanatization_status:=Sanatize(question)
	// if sanatization_status =="ERROR"{
	// 	fmt.Println("profile creation data invalid")
	// 	util.Message(w,1002)
	// 	return
	// }

	InsertQuestionInDB(question.Content,question.Type,question.Factor)
	util.Message(w,200)
}

func GetProfile(w http.ResponseWriter, r *http.Request){
	//CHECK IF USER IS AUTHORIZED TO SEE THE PROFILE
	header:=InsertQuestionHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie
	userId,dErr:=util.GetUserIdFromSession(sessionId)
	if dErr=="Error"{
		fmt.Println("Session Does not exist")
		util.Message(w,1003)
		return
	}

	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}
	
	profileRequest:=ProfileRequest{}
	profileRequestErr:=json.Unmarshal(requestBody,&profileRequest)
	if profileRequestErr!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}
	profileViewAuthorization:=ProfileViewAuthorization(userId,profileRequest.UserId)
	if profileViewAuthorization=="Error"{
		util.Message(w,2001)
		return
	}

	//RETURN PROFILE
	profileResponse:=GetProfileDB(profileRequest.UserId)
	SendProfile(w,profileResponse)

}