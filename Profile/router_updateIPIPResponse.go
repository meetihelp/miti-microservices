package Profile

import(
	"fmt"
	"net/http"
	// "log"
	"io/ioutil"
	// "strings"
	"strconv"
	"encoding/json"
   util "miti-microservices/Util"
   auth "miti-microservices/Authentication"
)

func UpdateIPIPResponse(w http.ResponseWriter, r *http.Request){
	header:=UpdateQuestionResponseHeader{}
	util.GetHeader(r,&header)


	sessionId:=header.Cookie

	userId,dErr:=util.GetUserIdFromSession(sessionId)
	fmt.Print("UpdateIPIP Header:")
	fmt.Println(header)
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

	updateIPIPRequest:=UpdateIPIPRequest{}
	errQuestionData:=json.Unmarshal(requestBody,&updateIPIPRequest)
	if errQuestionData!=nil{
		fmt.Println("Could not Unmarshall profile data")
		util.Message(w,1001)
		return
	}

	fmt.Print("UpdateIPIP Body:")
	fmt.Println(updateIPIPRequest)
	
    data:=make(map[string]int)
    data["IPIP1"]=updateIPIPRequest.IPIP1
    data["IPIP2"]=updateIPIPRequest.IPIP2
    data["IPIP3"]=updateIPIPRequest.IPIP3
    data["IPIP4"]=updateIPIPRequest.IPIP4
    data["IPIP5"]=updateIPIPRequest.IPIP5
	UpdateIPIPResponseDB(userId,data,updateIPIPRequest.Page)
	ipipStatus:=updateIPIPRequest.Page+1
	auth.UpdateIPIPStatus(userId,ipipStatus)
	UpdateIPIPScore(userId)
	util.Message(w,200)
}

func getDataInQuestionResponseForm(questionResponse QuestionResponse,data map[string]int,page int) (QuestionResponse){
	// ipipStatus:=0
	for key,value:=range data{
		ipipResponseNo:=GetIPIPResonseNo(key)
		questionNo:=page*5+ipipResponseNo
		questionId:=strconv.Itoa(questionNo)
		key="QuestionId"+questionId
		switch key{
		case "QuestionId1":
			// ipipStatus=1
			questionResponse.QuestionId1=value
		case "QuestionId2":
			// ipipStatus=1
			questionResponse.QuestionId2=value
		case "QuestionId3":
			// ipipStatus=1
			questionResponse.QuestionId3=value
		case "QuestionId4":
			// ipipStatus=1
			questionResponse.QuestionId4=value
		case "QuestionId5":
			// ipipStatus=1
			questionResponse.QuestionId5=value
		case "QuestionId6":
			// ipipStatus=2
			questionResponse.QuestionId6=value
		case "QuestionId7":
			// ipipStatus=2
			questionResponse.QuestionId7=value
		case "QuestionId8":
			// ipipStatus=2
			questionResponse.QuestionId8=value
		case "QuestionId9":
			// ipipStatus=2
			questionResponse.QuestionId9=value
		case "QuestionId10":
			// ipipStatus=2
			questionResponse.QuestionId10=value
		case "QuestionId11":
			// ipipStatus=3
			questionResponse.QuestionId11=value
		case "QuestionId12":
			// ipipStatus=3
			questionResponse.QuestionId12=value
		case "QuestionId13":
			// ipipStatus=3
			questionResponse.QuestionId13=value
		case "QuestionId14":
			// ipipStatus=3
			questionResponse.QuestionId14=value
		case "QuestionId15":
			// ipipStatus=3
			questionResponse.QuestionId15=value
		case "QuestionId16":
			// ipipStatus=4
			questionResponse.QuestionId16=value
		case "QuestionId17":
			// ipipStatus=4
			questionResponse.QuestionId17=value
		case "QuestionId18":
			// ipipStatus=4
			questionResponse.QuestionId18=value
		case "QuestionId19":
			// ipipStatus=4
			questionResponse.QuestionId19=value
		case "QuestionId20":
			// ipipStatus=4
			questionResponse.QuestionId20=value
		case "QuestionId21":
			// ipipStatus=5
			questionResponse.QuestionId21=value
		case "QuestionId22":
			// ipipStatus=5
			questionResponse.QuestionId22=value
		case "QuestionId23":
			// ipipStatus=5
			questionResponse.QuestionId23=value
		case "QuestionId24":
			// ipipStatus=5
			questionResponse.QuestionId24=value
		case "QuestionId25":
			// ipipStatus=5
			questionResponse.QuestionId25=value
		case "QuestionId26":
			// ipipStatus=6
			questionResponse.QuestionId26=value
		case "QuestionId27":
			// ipipStatus=6
			questionResponse.QuestionId27=value
		case "QuestionId28":
			// ipipStatus=6
			questionResponse.QuestionId28=value
		case "QuestionId29":
			// ipipStatus=6
			questionResponse.QuestionId29=value
		case "QuestionId30":
			// ipipStatus=6
			questionResponse.QuestionId30=value
		case "QuestionId31":
			// ipipStatus=7
			questionResponse.QuestionId31=value
		case "QuestionId32":
			// ipipStatus=7
			questionResponse.QuestionId32=value
		case "QuestionId33":
			// ipipStatus=7
			questionResponse.QuestionId33=value
		case "QuestionId34":
			// ipipStatus=7
			questionResponse.QuestionId34=value
		case "QuestionId35":
			// ipipStatus=7
			questionResponse.QuestionId35=value
		case "QuestionId36":
			// ipipStatus=8
			questionResponse.QuestionId36=value
		case "QuestionId37":
			// ipipStatus=8
			questionResponse.QuestionId37=value
		case "QuestionId38":
			// ipipStatus=8
			questionResponse.QuestionId38=value
		case "QuestionId39":
			// ipipStatus=8
			questionResponse.QuestionId39=value
		case "QuestionId40":
			// ipipStatus=8
			questionResponse.QuestionId40=value
		case "QuestionId41":
			// ipipStatus=9
			questionResponse.QuestionId41=value
		case "QuestionId42":
			// ipipStatus=9
			questionResponse.QuestionId42=value
		case "QuestionId43":
			// ipipStatus=9
			questionResponse.QuestionId43=value
		case "QuestionId44":
			// ipipStatus=9
			questionResponse.QuestionId44=value
		case "QuestionId45":
			// ipipStatus=9
			questionResponse.QuestionId45=value
		case "QuestionId46":
			// ipipStatus=10
			questionResponse.QuestionId46=value
		case "QuestionId47":
			// ipipStatus=10
			questionResponse.QuestionId47=value
		case "QuestionId48":
			// ipipStatus=10
			questionResponse.QuestionId48=value
		case "QuestionId49":
			// ipipStatus=10
			questionResponse.QuestionId49=value
		case "QuestionId50":
			// ipipStatus=10
			questionResponse.QuestionId50=value
		}
	}
	return questionResponse
}

func GetIPIPResonseNo(ipip string) int{
	if(ipip=="IPIP1"){
		return 1
	}
	if(ipip=="IPIP2"){
		return 2
	}
	if(ipip=="IPIP3"){
		return 3
	}
	if(ipip=="IPIP4"){
		return 4
	}
	if(ipip=="IPIP5"){
		return 5
	}
	return -1
}