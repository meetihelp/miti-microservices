package main 

import(
	database "app/Model/UseDatabase"
	"net/http"
)

func createMatch(w http.ResponseWriter,r *http.Request){
	user:=database.GetAllUser()

	num_of_user:=len(user)

	for i:=0;i<num_of_user;i=i+1{
		// if i+1<num_of_user{
		// 	database.Enter_Match_user(user[i],user[i+1])
		// }
		for j:=i;j<num_of_user;j=j+1{
			database.Enter_Match_user(user[i],user[j])
		}
	}


}