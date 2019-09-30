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


func Profile_creation(w http.ResponseWriter, r *http.Request){
	header:=Profile_creation_Header{}
	util.GetHeader(r,&header)


	session_id:=header.Cookie

	user_id,d_err:=util.Get_user_id_from_session(session_id)
	if d_err=="ERROR"{
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
	profile_data:=Profile{}
	err_profile_data:=json.Unmarshal(requestBody,&profile_data)
	if err_profile_data!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}
	fmt.Println(profile_data.Name)
	profile_data.UserId=user_id
	sanatization_status:=Sanatize(profile_data)
	if sanatization_status =="ERROR"{
		fmt.Println("profile creation data invalid")
		util.Message(w,1002)
		return
	}
	// profile_data_handle(w,profile_data)
	Enter_profile_data(profile_data)

	util.Message(w,200)

}

// func profile_data_handle(w http.ResponseWriter,profile_data Profile){
// 	Enter_profile_data(profile_data)
// 	util.Message(w,200)
// 	return
// }

func GetQuestion(w http.ResponseWriter, r *http.Request){
	header:=Profile_creation_Header{}
	util.GetHeader(r,&header)


	session_id:=header.Cookie

	_,d_err:=util.Get_user_id_from_session(session_id)
	if d_err=="ERROR"{
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
	err_question_data:=json.Unmarshal(requestBody,&questionRequest)
	if err_question_data!=nil{
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
	header:=UpdateQuestionResponse_Header{}
	util.GetHeader(r,&header)


	session_id:=header.Cookie

	userId,d_err:=util.Get_user_id_from_session(session_id)
	if d_err=="ERROR"{
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
	

	responseWrapper:=ResponseWrapper{}
	// questionResponse.UserId=userId
	err_question_data:=json.Unmarshal(requestBody,&responseWrapper)
	fmt.Println(responseWrapper.IPIP)
	response:=[]Response{}
	response=responseWrapper.IPIP
	fmt.Println(response[0].QuestionId)
	if err_question_data!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}

	// sanatization_status:=Sanatize(questionResponse)
	// if sanatization_status =="ERROR"{
	// 	fmt.Println("profile creation data invalid")
	// 	util.Message(w,1002)
	// 	return
	// }

	InsertQuestionResponse(userId,responseWrapper.IPIP)

	// score:=GetScore(responseWrapper.IPIP)
	// UpdateScore(userId,score)
	util.Message(w,200)
}

func InsertQuestion(w http.ResponseWriter, r *http.Request){
	header:=InsertQuestion_Header{}
	util.GetHeader(r,&header)


	session_id:=header.Cookie

	_,d_err:=util.Get_user_id_from_session(session_id)
	if d_err=="ERROR"{
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
	err_question_data:=json.Unmarshal(requestBody,&question)
	if err_question_data!=nil{
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