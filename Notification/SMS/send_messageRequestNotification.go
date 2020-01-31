package SMS
import(
	"net/http"
	"net/url"
	"log"
	"fmt"
	util "miti-microservices/Util"
)
func MessageRequestNotificaton(senderName string,senderPhone string,phone string) (*http.Response,error){
	fmt.Println("Message Request Sending by "+senderName+" by phone->"+senderPhone+" to phone->"+phone)
	if(len(senderName)>9){
		senderName=senderName[:9]
	}

	base, err := url.Parse("")
	if err != nil {
		return nil,err
	}
	q := url.Values{}
	q.Add("invisible", "1")
	q.Add("otp",senderName)
	q.Add("mobile",phone)
	authk:=GetAuth()
	if(authk==""){
		log.Println("Please set authkey for message")
		return nil,err
	}
	q.Add("authkey", authk)
	q.Add("template_id","5e3472f8d6fc055db360e0d7")
	q.Add("otp_expiry","10")
	base.RawQuery = q.Encode()
	client:=util.GetClient(2)
	resp, err1:=client.Post(base.String(),"",nil)
	fmt.Println(resp)
	if err1!=nil {
		log.Print(err)
	}
	return resp,err1
}