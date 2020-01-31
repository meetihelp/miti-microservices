package Util

import(
	"fmt"
	"strconv"
)

func APIHitLog(api string,ipAddress string,sessionId string){
	timeNow:=GetTime()
	statement:=api+","+ipAddress+","+sessionId+","+timeNow+","+"APIHitLog"
	fmt.Println(statement)
}

func SessionLog(api string,ipAddress string,sessionId string,status string){
	timeNow:=GetTime()
	statement:=api+","+ipAddress+","+sessionId+","+timeNow+",SessionLog_"+status
	fmt.Println(statement)
}


func TemporarySessionLog(api string,ipAddress string,sessionId string,status string){
	timeNow:=GetTime()
	statement:=api+","+ipAddress+","+sessionId+","+timeNow+",TemporarySessionLog_"+status
	fmt.Println(statement)
}

func BodyLog(api string,ipAddress string,sessionId string,body interface{}){
	timeNow:=GetTime()
	statement:=api+","+ipAddress+","+sessionId+","+timeNow+","+"BodyLog"+","
	fmt.Print(statement)
	fmt.Println(body)
}

func ResponseLog(api string,ipAddress string,sessionId string,code int,response interface{}){
	timeNow:=GetTime()
	statement:=api+","+ipAddress+","+sessionId+","+timeNow+","+"ResponseLog"+","+strconv.Itoa(code)+","
	fmt.Print(statement)
	fmt.Println(response)
}