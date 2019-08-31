package Authentication

import(
	"net/http"
	// "fmt"
	"log"
	// util "app/Utility"
)


func verify_email(w http.ResponseWriter,r *http.Request){
	tokens, ok := r.URL.Query()["token"]
    if !ok || len(tokens[0]) < 1 {
        log.Println("Url Param 'token' is missing")
        return
    }
    token:=tokens[0]
    

    user_id,email_verify:=Verify_Email(token)
    if email_verify{
    	//CHANGE STATUS OF USER TO VERIFIED
    	Change_Verification_Status(user_id)
    }
}