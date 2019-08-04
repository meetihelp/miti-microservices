package main

import (
	"net/http"
	// "io"
	"encoding/json"
	"fmt"
)

type Data struct{
	temp string `json:"data"`
}
func homepage(w http.ResponseWriter, r *http.Request){
	data:=Data{temp:"gaurav kumar jha"}
	d,_:=json.Marshal(data)
	 w.Header().Set("Content-Type", "application/json")
    enc := json.NewEncoder(w)
	err := enc.Encode(d)
	fmt.Println(err)
}