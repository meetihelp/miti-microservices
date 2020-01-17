package SMS
import(
	"fmt"
)
func MessageRequestNotificaton(senderName string,senderPhone string,phone string){
	fmt.Println("Message Request Sending by "+senderName+" by phone->"+senderPhone+" to phone->"+phone)
}