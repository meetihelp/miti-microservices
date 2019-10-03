package Authentication
import(
	"net/http"
	util "app/Util"
)

func CreateMatch(w http.ResponseWriter,r *http.Request){
	user:=GetAllUser()

	num_of_user:=len(user)

	for i:=0;i<num_of_user;i=i+1{
		// if i+1<num_of_user{
		// 	database.Enter_Match_user(user[i],user[i+1])
		// }
		for j:=i+1;j<num_of_user;j=j+1{
			Enter_Match_user(user[i],user[j])
		}
	}

	util.Message(w,200)
}