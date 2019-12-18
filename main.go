package main

import (
	Api "miti-microservices/Api"
)

const(
	Domain ="http://localhost:9000"
)

func main(){
	Api.Server()
	
}