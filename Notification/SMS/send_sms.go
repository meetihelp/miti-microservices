package SMS

import(
	"net/http"
	"net/url"
	"log"
	"fmt"
	"os"
	"strings"
	util "miti-microservices/Util"
	
)

func GetAuth() string{
	return os.Getenv("msg91_authkey")
}
func SendSMS(phone string,otp string) (*http.Response,error){
	base, err := url.Parse("https://api.msg91.com/api/v5/otp")
	// base, err := url.Parse("")
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
	q.Add("template_id","5dfa1cdbd6fc054db941c67a")

	// q.Add("invisible", "1")
	q.Add("otp",strings.TrimSpace(otp))
	q.Add("mobile",strings.TrimSpace(phone))
	// q.Add("otp_expiry","10")
	base.RawQuery = q.Encode()
	client:=util.GetClient(2)
	fmt.Println(base.String());
	resp, err1:=client.Get(base.String())
	fmt.Println(resp.Body)
	if err1!=nil {
		log.Print(err)
	}
	return resp,err1
}

func ReSendSMSHelper(phone string){
	base, err := url.Parse("http://api.msg91.com/api/v5/otp/retry")
	// base, err := url.Parse("")
	if err != nil {
		return
	}
	q := url.Values{}
	authk:=GetAuth()
	if(authk==""){
		log.Println("Please set authkey for message")
		return
	}
	q.Add("authkey", authk)
	q.Add("mobile",phone)
	base.RawQuery = q.Encode()
	resp, err1:=http.Get(base.String())
	fmt.Println(resp.Body)
	if err1!=nil {
		log.Print(err)
	}
}

