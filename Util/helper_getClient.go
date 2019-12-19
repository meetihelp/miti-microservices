package Util

import(
	"net/http"
	"time"
)

func GetClient(timeout int)http.Client{
	client := http.Client{
    	Timeout: time.Duration(timeout) * time.Second,
	}
	return client
}