package Util

import(
	"time"
)

func GetTime() string{
	time_now:=time.Now()
	temp:=time_now.Format("2006-01-02 15:04:05")
	return temp
}