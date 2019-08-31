package Authentication
import (
	"net/http"
	"fmt"
	util "app/Util"
	"io/ioutil"
	"encoding/json"
	// "strings"
)


func login(w http.ResponseWriter,r *http.Request){
	ip_address:=util.Get_IP_address(r)
	login_header:=Login_header{}
	util.GetHeader(r,&login_header)
	session_id:=login_header.Cookie
	fmt.Println(session_id)
	user_id,login_status:=util.Get_user_id_from_session(session_id)
	fmt.Println("session "+login_status)
	if login_status=="OK"{
		util.Message(w,200)
		return
	}

	//Read body data
	requestBody,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Could not read body")
		util.Message(w,1000)
		return 
	}

	//UNMARSHILING DATA
	user_data :=User{}
	err_user_data:=json.Unmarshal(requestBody,&user_data)
	if err_user_data!=nil{
		fmt.Println("Could not Unmarshall user data")
		util.Message(w,1001)
		return 
	}

	sanatization_status :=Sanatize(user_data)
	if sanatization_status =="ERROR"{
		fmt.Println("User data invalid")
		util.Message(w,1002)
		return
	}

	user_id,login_status=Check_user(user_data)
	if login_status=="WRONG_PASSWORD"{
		util.Message(w,1501)
		return
	}
	if login_status=="NO_USER"{
		util.Message(w,1501)
		return
	}
	if login_status=="UNVERIFIED"{
		cookie:=util.Insert_session(user_id,ip_address)
		// http.SetCookie(w,&cookie)
		w.Header().Set("miti-Cookie",cookie)
		util.Message(w,1005)
		return
	} 
	if login_status=="OK"{
		cookie:=util.Insert_session(user_id,ip_address)
		// http.SetCookie(w,&cookie)
		w.Header().Set("miti-Cookie",cookie)
		util.Message(w,200)
		return
	}
}
