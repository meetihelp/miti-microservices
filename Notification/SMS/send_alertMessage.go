package SMS

import(
	"fmt"
)

func AlertNotification(phoneList []string,name string,latitude string,longitude string) string{
	fmt.Print("Sending Alert Message to:")
	fmt.Println(phoneList)
	msg:=name+" is in danger at latitude:"+latitude+",longitude:"+longitude
	fmt.Println(msg)
	return "Ok"
}