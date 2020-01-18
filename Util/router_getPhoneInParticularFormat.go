package Util

import(
	"strings"
)
func GetPhoneInFormat(phone string) (string,string){
	temp:=strings.Replace(phone," ","",-1)
	temp=strings.Replace(temp,"\t","",-1)
	temp=strings.Replace(temp,"-","",-1)
	if(len(temp)<10){
		return "Error",""
	}
	temp=temp[len(temp)-10:]
	return "Ok",temp
}