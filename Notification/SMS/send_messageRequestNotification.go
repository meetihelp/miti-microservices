package SMS
import(
	"io/ioutil"
	"net/http"
	"net/url"
	"log"
	"fmt"
	util "miti-microservices/Util"
	"bytes"
	"encoding/json"
)

type MessageRequestSMS struct{
	Mobiles string `json:"mobiles"`
	AuthKey string `json:"authkey"`
	TemplateId string `json:"template_id"`
	Name string `json:"name"`
}
func MessageRequestNotificaton(senderName string,senderPhone string,phone string) (*http.Response,error){
	fmt.Println("Message Request Sending by "+senderName+" by phone->"+senderPhone+" to phone->"+phone)
	

	base, err := url.Parse("api.msg91.com/api/v5/flow/?response=json")
	if err != nil {
		return nil,err
	}

	authk:=GetAuth()
	if(authk==""){
		log.Println("Please set authkey for message")
		return nil,err
	}

	data:=MessageRequestSMS{}
	data.Mobiles=phone
	data.AuthKey=authk
	data.TemplateId="5e347dddd6fc050f2941ce54"
	data.Name=senderName+",phone:"+senderPhone

	msgByte,errMsg:=json.Marshal(data)
	if(errMsg!=nil){
		fmt.Println(errMsg)
	}
	msg:=bytes.NewReader(msgByte)


	client:=util.GetClient(2)
	resp, err1:=client.Post(base.String(),"application/json",msg)
	fmt.Println(base.String())

	bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    bodyString := string(bodyBytes)
    fmt.Println(bodyString)

	if err1!=nil {
		log.Print(err)
	}
	return resp,err1
}