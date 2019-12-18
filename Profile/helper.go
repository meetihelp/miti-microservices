package Profile

import(
	"net/http"
	// "io/ioutil"
	"encoding/json"	
	util "miti-microservices/Util"
	"log"
	// "reflect"
	// "fmt"
)

func SendQuestion(w http.ResponseWriter,question []Question){
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&SendQuestionContent{Code:200,Message:msg,Question:question}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}

func SendProfile(w http.ResponseWriter,profileResponse ProfileResponse){
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=ProfileResponseContent{Code:200,Message:msg,ProfileResponse:profileResponse}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}

func ConvertQuestionResponseToArray(questionResponse QuestionResponse) ([]int){
	response:=make([]int,51)
	response[1]=questionResponse.QuestionId1
	response[2]=questionResponse.QuestionId2
	response[3]=questionResponse.QuestionId3
	response[4]=questionResponse.QuestionId4
	response[5]=questionResponse.QuestionId5
	response[6]=questionResponse.QuestionId6
	response[7]=questionResponse.QuestionId7
	response[8]=questionResponse.QuestionId8
	response[9]=questionResponse.QuestionId9
	response[10]=questionResponse.QuestionId10
	response[11]=questionResponse.QuestionId11
	response[12]=questionResponse.QuestionId12
	response[13]=questionResponse.QuestionId13
	response[14]=questionResponse.QuestionId14
	response[15]=questionResponse.QuestionId15
	response[16]=questionResponse.QuestionId16
	response[17]=questionResponse.QuestionId17
	response[18]=questionResponse.QuestionId18
	response[19]=questionResponse.QuestionId19
	response[20]=questionResponse.QuestionId20
	response[21]=questionResponse.QuestionId21
	response[22]=questionResponse.QuestionId22
	response[23]=questionResponse.QuestionId23
	response[24]=questionResponse.QuestionId24
	response[25]=questionResponse.QuestionId25
	response[26]=questionResponse.QuestionId26
	response[27]=questionResponse.QuestionId27
	response[28]=questionResponse.QuestionId28
	response[29]=questionResponse.QuestionId29
	response[30]=questionResponse.QuestionId30
	response[31]=questionResponse.QuestionId31
	response[32]=questionResponse.QuestionId32
	response[33]=questionResponse.QuestionId33
	response[34]=questionResponse.QuestionId34
	response[35]=questionResponse.QuestionId35
	response[36]=questionResponse.QuestionId36
	response[37]=questionResponse.QuestionId37
	response[38]=questionResponse.QuestionId38
	response[39]=questionResponse.QuestionId39
	response[40]=questionResponse.QuestionId40
	response[41]=questionResponse.QuestionId41
	response[42]=questionResponse.QuestionId42
	response[43]=questionResponse.QuestionId43
	response[44]=questionResponse.QuestionId44
	response[45]=questionResponse.QuestionId45
	response[46]=questionResponse.QuestionId46
	response[47]=questionResponse.QuestionId47
	response[48]=questionResponse.QuestionId48
	response[49]=questionResponse.QuestionId49
	response[50]=questionResponse.QuestionId50
	return response
}