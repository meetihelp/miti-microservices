package Profile

import(
	"net/http"
	// "io/ioutil"
	"encoding/json"	
	util "app/Util"
	"log"
	// "reflect"
	// "fmt"
)

func SendQuestion(w http.ResponseWriter,question []Question){
	w.Header().Set("Content-Type", "application/json")
	msg:=util.Get_message_decode(200)
	p:=&SendQuestion_Content{Code:200,Message:msg,Question:question}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}

// func GetData(r *http.Request,data interface{}){
// 	val:=reflect.ValueOf(data).Elem()
// 	data_type:=reflect.TypeOf(data).Elem()
// 	body,_:=ioutil.ReadAll(r.Body)
// 	fmt.Println(body)
// 	for i:=0 ;i<val.NumField();i++{
// 		fld:=val.Field(i)
// 		tag:=data_type.Field(i).Tag.Get("json")
// 		body_data,ok:= body[tag]
// 		if ok{
// 			switch fld.Kind(){
// 			case reflect.String:
// 				fld.SetString(body_data[0])
// 			}
// 		}
// 	}
// }