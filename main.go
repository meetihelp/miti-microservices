package main

import (
	Api "miti-microservices/Api"
	"flag"
)

const(
	Domain ="http://localhost:9000"
)

func main(){
	runMethod:=flag.String("runmethod","Devlopment","This is flag for Devlopment or Production Code method")
	flag.Parse()
	// logMethod:=flag.String("log","Intensive","This is flag for intensive or extnsive logging")
	Api.Server(*runMethod)
	
}