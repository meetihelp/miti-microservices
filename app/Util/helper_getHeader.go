package Util

import(
	"net/http"
	"reflect"
)


func GetHeader(r *http.Request,data interface{}){
	val:=reflect.ValueOf(data).Elem()
	data_type:=reflect.TypeOf(data).Elem()
	header:=r.Header
	for i:=0 ;i<val.NumField();i++{
		fld:=val.Field(i)
		tag:=data_type.Field(i).Tag.Get("header")
		header_data,ok:= header[tag]
		if ok{
			switch fld.Kind(){
			case reflect.String:
				fld.SetString(header_data[0])
			}
		}
	}
}