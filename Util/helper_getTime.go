package Util

import(
	"time"
)

func GetTime() string{
	timeNow:=time.Now()
	temp:=timeNow.Format("2006-01-02 15:04:05")
	return temp
}

func GetDateFromTime(timestamp string) string{
	return timestamp[:10]
}