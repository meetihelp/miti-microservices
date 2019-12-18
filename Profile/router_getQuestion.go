package Profile

import(
	"fmt"
	"net/http"
	// "log"
	"io/ioutil"
	// "strings"
	"encoding/json"
   util "miti-microservices/Util"
)

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