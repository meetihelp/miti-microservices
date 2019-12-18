package SMS

import(
	"net/http"
	"net/url"
	"log"
	"os"
)

func GetAuth string(){
	return os.Getenv("msg91_authkey")
}
func SendSMS(phone string,otp string){
	base, err := url.Parse("https://api.msg91.com/api/v5/otp")
	if err != nil {
		return
	}
	q := url.Values{}
	q.Add("invisible", 1)
	q.Add("otp",otp)
	q.Add("mobile",phone)

	q.Add("authkey","308893A1u1gEJGa9U5df9fb61")
	q.Add("template_id","5dfa1cdbd6fc054db941c67a")
	q.Add("otp_expiry",10)
	base.RawQuery = q.Encode()
	resp, err1:=http.POST(base.String())
	if err1!=nil {
		log.Print(err)
	}
}

func ReSendSMS(phone string){
	base, err := url.Parse("https://api.msg91.com/api/v5/otp/retry")
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
	if err1!=nil {
		log.Print(err)
	}
}