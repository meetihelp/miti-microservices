package Util

import(
	"net/http"
	"net"
)

func GetIPAddress(r *http.Request) string{
	ip,_,_ :=net.SplitHostPort(r.RemoteAddr)

	realIp:=r.Header.Get("X-FORWARDED-FOR")
	if realIp !=""{
		return realIp
	} else{
		return ip
	}
}