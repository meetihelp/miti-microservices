package main 

import(
	"net/http"
	"fmt"
	// CD "app/Model/CreateDatabase"
	database "app/Model/UseDatabase"
	util "app/Utility"
	// "io/ioutil"
	// "encoding/json"
	// "strings"
)

type Logout_header struct{
	Cookie string `header:"Miti-Cookie"`
}
func logout(w http.ResponseWriter,r *http.Request){
	// ip_address:=util.Get_IP_address(r)
	fmt.Println("Delete")
	logout_header:=Logout_header{}
	util.GetHeader(r,&logout_header)
	session_id:=logout_header.Cookie
	logout_status:=database.Delete_session(session_id)
	if logout_status=="OK"{
		util.Message(w,200)
		return
	}
}