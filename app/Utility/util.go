package Utility

import(
	"github.com/nu7hatch/gouuid"
	// "fmt"
)
func Generate_user_Id() string{
	u, _ := uuid.NewV4()
	return u.String()
}

func init(){
	
}