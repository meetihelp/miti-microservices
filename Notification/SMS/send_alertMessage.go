package SMS

import(
	"fmt"
)

func AlertNotification(phoneList []string,name string,latitude string,longitude string) string{
	fmt.Print("Sending Alert Message to:")
	fmt.Println(phoneList)
	return "Ok"
}