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
	r.HandleFunc("/send_verification_link",send_verification_link).Methods("GET")
	// r.HandleFunc("/login",router.login).Methods("POST")
	http.Handle("/", r)
	if err := http.ListenAndServe("0.0.0.0:9000", nil); err != nil {
		log.Fatal(err)
	}
}