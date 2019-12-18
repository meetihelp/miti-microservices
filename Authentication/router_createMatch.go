package Authentication
import(
	"net/http"
	util "miti-microservices/Util"
)

func CreateMatch(w http.ResponseWriter,r *http.Request){
	user:=GetAllUser()

	numOfUser:=len(user)

	for i:=0;i<numOfUser;i=i+1{
		for j:=i+1;j<numOfUser;j=j+1{
			EnterMatchUser(user[i],user[j])
		}
	}

	util.Message(w,200)
}