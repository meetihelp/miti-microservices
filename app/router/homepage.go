package main

import (
	"net/http"
	"io"
)

type Data struct{
	temp string
}
func homepage(w http.ResponseWriter, r *http.Request){
	io.WriteString(w,"homepage")
}