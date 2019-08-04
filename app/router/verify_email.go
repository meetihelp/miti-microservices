package main

import(
	"net/http"
	// "fmt"
	"log"
	database "app/Model/UseDatabase"
	// util "app/Utility"
)


func verify_email(w http.ResponseWriter,r *http.Request){
	tokens, ok := r.URL.Query()["token"]
    if !ok || len(tokens[0]) < 1 {
        log.Println("Url Param 'token' is missing")
        return
    }
    token:=tokens[0]
    

    user_id,verify:=database.Verify_Email(token)
    if verify{
    	//CHANGE STATUS OF USER TO VERIFIED
    	database.Change_Verification_Status(user_id)
    }
}