package Utility

import(
	"github.com/nu7hatch/gouuid"
)

func Generate_token() string{
	u, _ := uuid.NewV4()
	token:=u.String()
	return token
}