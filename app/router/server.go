package main

import (
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"net/http"
)

func server(){
	fmt.Println("Server running.....")
	r := mux.NewRouter()
	r.HandleFunc("/", homepage).Methods("GET")
	r.HandleFunc("/register",register).Methods("POST")
	r.HandleFunc("/profile_creation",profile_creation).Methods("POST")
	r.HandleFunc("/generate_verification_email",generate_verification_email).Methods("GET")
	r.HandleFunc("/generate_otp",generate_otp).Methods("GET")
	r.HandleFunc("/verify_email",verify_email).Methods("GET")
	r.HandleFunc("/verify_otp",verify_otp).Methods("POST")
	r.HandleFunc("/login",login).Methods("POST")
	r.HandleFunc("/logout",logout).Methods("GET")
	r.HandleFunc("/update_password",update_password).Methods("POST")
	r.HandleFunc("/createMatch",createMatch).Methods("GET")
	r.HandleFunc("/getChatDetail",getChatDetail).Methods("POST")
	r.HandleFunc("/getAllChat",getAllChat).Methods("POST")
	r.HandleFunc("/getUnreadChat",getUnreadChat).Methods("POST")
	r.HandleFunc("/chat",chat).Methods("POST")
	http.Handle("/", r)
	if err := http.ListenAndServe("0.0.0.0:9000", nil); err != nil {
		log.Fatal(err)
	}
}