package Util

import(
	"net/http"
)

func GetResponseFormatHeader(w http.ResponseWriter,header map[string]string) http.ResponseWriter{
	for key,value:=range header{
		w.Header().Set(key, value)
	}
	return w
}