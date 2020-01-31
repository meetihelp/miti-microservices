package SMS
import(
	"io/ioutil"
	"net/http"
	"net/url"
	"log"
	"fmt"
	util "miti-microservices/Util"
	// "bytes"
	"strings"
	"encoding/json"
)

// type MessageRequestSMS struct{
// 	Mobiles string `json:"mobiles"`
// 	AuthKey string `json:"authkey"`
// 	TemplateId string `json:"template_id"`
// 	Name string `json:"name"`
// }

type MessageRequestSMS struct{
	VAR1 string `json:"VAR1"`
}
func MessageRequestNotificaton(senderName string,senderPhone string,phone string) (*http.Response,error){
	fmt.Println("Message Request Sending by "+senderName+" by phone->"+senderPhone+" to phone->"+phone)
	

	base, err := url.Parse("https://api.msg91.com/api/v5/otp")
	if err != nil {
		return nil,err
	}

	q := url.Values{}

	authk:=GetAuth()
	if(authk==""){
		log.Println("Please set authkey for message")
		return nil,err
	}
	q.Add("authkey", authk)
	q.Add("template_id","5e3485ffd6fc0523784fb2b2")

	// q.Add("invisible", "1")
	q.Add("otp",strings.TrimSpace("1234"))
	q.Add("mobile",strings.Trim(phone,"'"))


	data:=MessageRequestSMS{}
	data.VAR1=senderName
	name,_:=json.Marshal(data)
	q.Add("extra_param",string(name))
	// q.Add("otp_expiry","10")
	base.RawQuery = q.Encode()
	client:=util.GetClient(2)
	fmt.Println(base.String());
	resp, err1:=client.Get(base.String())
	// data:=MessageRequestSMS{}
	// data.Mobiles=phone
	// data.AuthKey=authk
	// data.TemplateId="5e347dddd6fc050f2941ce54"
	// data.Name=senderName+",phone:"+senderPhone

	// msgByte,errMsg:=json.Marshal(data)
	// if(errMsg!=nil){
	// 	fmt.Println(errMsg)
	// }
	// msg:=bytes.NewReader(msgByte)


	// client:=util.GetClient(2)
	// resp, err1:=client.Post(base.String(),"application/json",msg)
	// fmt.Println(base.String())

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