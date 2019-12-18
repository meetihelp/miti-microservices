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
	// data:=QuestionResponse{}
	if err := json.Unmarshal(requestBody, &data); err != nil {
        panic(err)
    }
    
	InsertQuestionResponse(userId,data)
	UpdateIPIPScore(userId)
	util.Message(w,200)
}

func getDataInQuestionResponseForm(questionResponse QuestionResponse,data map[string]int) QuestionResponse{
	for key,value:=range data{
		switch key{
		case "QuestionId1":
			questionResponse.QuestionId1=value
		case "QuestionId2":
			questionResponse.QuestionId2=value
		case "QuestionId3":
			questionResponse.QuestionId3=value
		case "QuestionId4":
			questionResponse.QuestionId4=value
		case "QuestionId5":
			questionResponse.QuestionId5=value
		case "QuestionId6":
			questionResponse.QuestionId6=value
		case "QuestionId7":
			questionResponse.QuestionId7=value
		case "QuestionId8":
			questionResponse.QuestionId8=value
		case "QuestionId9":
			questionResponse.QuestionId9=value
		case "QuestionId10":
			questionResponse.QuestionId10=value
		case "QuestionId11":
			questionResponse.QuestionId11=value
		case "QuestionId12":
			questionResponse.QuestionId12=value
		case "QuestionId13":
			questionResponse.QuestionId13=value
		case "QuestionId14":
			questionResponse.QuestionId14=value
		case "QuestionId15":
			questionResponse.QuestionId15=value
		case "QuestionId16":
			questionResponse.QuestionId16=value
		case "QuestionId17":
			questionResponse.QuestionId17=value
		case "QuestionId18":
			questionResponse.QuestionId18=value
		case "QuestionId19":
			questionResponse.QuestionId19=value
		case "QuestionId20":
			questionResponse.QuestionId20=value
		case "QuestionId21":
			questionResponse.QuestionId21=value
		case "QuestionId22":
			questionResponse.QuestionId22=value
		case "QuestionId23":
			questionResponse.QuestionId23=value
		case "QuestionId24":
			questionResponse.QuestionId24=value
		case "QuestionId25":
			questionResponse.QuestionId25=value
		case "QuestionId26":
			questionResponse.QuestionId26=value
		case "QuestionId27":
			questionResponse.QuestionId27=value
		case "QuestionId28":
			questionResponse.QuestionId28=value
		case "QuestionId29":
			questionResponse.QuestionId29=value
		case "QuestionId30":
			questionResponse.QuestionId30=value
		case "QuestionId31":
			questionResponse.QuestionId31=value
		case "QuestionId32":
			questionResponse.QuestionId32=value
		case "QuestionId33":
			questionResponse.QuestionId33=value
		case "QuestionId34":
			questionResponse.QuestionId34=value
		case "QuestionId35":
			questionResponse.QuestionId35=value
		case "QuestionId36":
			questionResponse.QuestionId36=value
		case "QuestionId37":
			questionResponse.QuestionId37=value
		case "QuestionId38":
			questionResponse.QuestionId38=value
		case "QuestionId39":
			questionResponse.QuestionId39=value
		case "QuestionId40":
			questionResponse.QuestionId40=value
		case "QuestionId41":
			questionResponse.QuestionId41=value
		case "QuestionId42":
			questionResponse.QuestionId42=value
		case "QuestionId43":
			questionResponse.QuestionId43=value
		case "QuestionId44":
			questionResponse.QuestionId44=value
		case "QuestionId45":
			questionResponse.QuestionId45=value
		case "QuestionId46":
			questionResponse.QuestionId46=value
		case "QuestionId47":
			questionResponse.QuestionId47=value
		case "QuestionId48":
			questionResponse.QuestionId48=value
		case "QuestionId49":
			questionResponse.QuestionId49=value
		case "QuestionId50":
			questionResponse.QuestionId50=value
		}
	}
	return questionResponse
}