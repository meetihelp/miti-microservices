package Util

import(
	"net/http"
	"net"
)

func Get_IP_address(r *http.Request) string{
	ip,_,_ :=net.SplitHostPort(r.RemoteAddr)

	real_ip:=r.Header.Get("X-FORWARDED-FOR")
	if real_ip !=""{
		return real_ip
	} else{
		return ip
	}
}