package Util

import (
	// "io/ioutil"
	// "log"
	// "net/http"
    "io"
    "crypto/rand"
)
func EncodeToString(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
    b := make([]byte, max)
    n, err := io.ReadAtLeast(rand.Reader, b, max)
    if n != max {
        panic(err)
    }
    for i := 0; i < len(b); i++ {
        b[i] = table[int(b[i])%len(table)]
    }
    return string(b)
}
func GenerateOTP() string{
	ret := EncodeToString(5)
	return ret
}