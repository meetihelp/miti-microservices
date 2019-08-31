package Authentication

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	// "io"
	"encoding/json"
   util "app/Util"
)

func register(w http.ResponseWriter, r *http.Request){
	//Get ip address of user
	ip_address:=util.Get_IP_address(r)
	fmt.Println("Registeration request from "+ip_address)

	//GET HEADER 
	header:=Register_Header{}
	util.GetHeader(r,&header)

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

	//SANITIZE USER AND PROFILE DATA
	sanatization_status :=Sanatize(user_data)
	if sanatization_status =="ERROR"{
		fmt.Println("User data invalid")
		util.Message(w,1002)
		return
	}

	user_id,ok:=user_data_handle(w,user_data)
	if ok{
		user_data.User_id=user_id
		cookie:=util.Insert_session(user_data.User_id,ip_address)
		w.Header().Set("miti-Cookie",cookie)
		util.Message(w,200)
	}
}

func user_data_handle(w http.ResponseWriter, user_data User) (string,bool){
	user_id,db_status:=Enter_user_data(user_data)
	if db_status ==1{
		log.Println("User Already exist")
		util.Message(w,1101)
		return user_id,false
	} else{
		log.Println("User data entered successfully")
		return user_id,true
	}
}



