package Authentication

// import(
// 	"net/http"
// 	"log"
// )


// func VerifyEmail(w http.ResponseWriter,r *http.Request){
// 	tokens, ok := r.URL.Query()["token"]
//     if !ok || len(tokens[0]) < 1 {
//         log.Println("Url Param 'token' is missing")
//         return
//     }
//     token:=tokens[0]
    

//     userId,emailVerify:=VerifyEmailFunc(token)
//     if emailVerify{
//     	//CHANGE STATUS OF USER TO VERIFIED
//     	ChangeVerificationStatus(userId)
//     }
// }