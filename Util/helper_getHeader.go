package Util

import(
	"net/http"
	"reflect"
)


func GetHeader(r *http.Request,data interface{}){
	val:=reflect.ValueOf(data).Elem()
	dataType:=reflect.TypeOf(data).Elem()
	header:=r.Header
	for i:=0 ;i<val.NumField();i++{
		fld:=val.Field(i)
		tag:=dataType.Field(i).Tag.Get("header")
		headerData,ok:= header[tag]
		if ok{
			switch fld.Kind(){
			case reflect.String:
				fld.SetString(headerData[0])
			}
		}
	}
}
