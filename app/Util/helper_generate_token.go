package Util

import(
	"strings"
	"github.com/nu7hatch/gouuid"
)

func GenerateToken() string{
	u, _ := uuid.NewV4()
	token:=u.String()
	token= strings.Replace(token, "-", "", -1)

	return token
}