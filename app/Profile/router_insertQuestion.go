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