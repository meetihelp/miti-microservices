package main

import (
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"net/http"
	profile "app/Profile"
	apnaauth "app/Authentication"
	apnachat "app/Chat"
)

func server(){
	fmt.Println("Server running.....")
	r := mux.NewRouter()
	// r.HandleFunc("/", homepage).Methods("GET")
	r.HandleFunc("/register",apnaauth.Register).Methods("POST")
	r.HandleFunc("/generate_verification_email",apnaauth.Generate_verification_email).Methods("GET")
	r.HandleFunc("/generate_otp",apnaauth.Generate_otp).Methods("GET")
	r.HandleFunc("/verify_email",apnaauth.Verify_email).Methods("GET")
	r.HandleFunc("/verify_otp",apnaauth.Verify_otp).Methods("POST")
	r.HandleFunc("/login",apnaauth.Login).Methods("POST")
	r.HandleFunc("/logout",apnaauth.Logout).Methods("GET")
	r.HandleFunc("/update_password",apnaauth.Update_password).Methods("POST")
	// r.HandleFunc("/createMatch",createMatch).Methods("GET")

	r.HandleFunc("/getChatDetail",apnachat.GetChatDetailroute).Methods("POST")
	r.HandleFunc("/getChat",apnachat.GetChat).Methods("POST")
	// r.HandleFunc("/getUnreadChat",getUnreadChat).Methods("POST")
	r.HandleFunc("/chat",apnachat.Chatinsert).Methods("POST")
	r.HandleFunc("/getChatAfterIndex",apnachat.GetChatAfterIndex).Methods("POST")

	r.HandleFunc("/profileCreation",profile.Profile_creation).Methods("POST")
	r.HandleFunc("/getQuestion",profile.GetQuestion).Methods("GET")
	r.HandleFunc("/insertQuestion",profile.InsertQuestion).Methods("POST")
	r.HandleFunc("/updateQuestionResponse",profile.UpdateQuestionResponse).Methods("POST")
	http.Handle("/", r)
	if err := http.ListenAndServe("0.0.0.0:9000", nil); err != nil {
		log.Fatal(err)
	}
}