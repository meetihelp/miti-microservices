package main

import (
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"net/http"
	profile "app/Profile"
	apnaauth "app/Authentication"
	apnachat "app/Chat"
	util "app/Util"
	gps "app/GPS"
	event "app/Event"
	image "app/Image"
)

func test(w http.ResponseWriter,r *http.Request){
	util.Message(w,200)
}


func server(){
	fmt.Println("Server running.....")
	r := mux.NewRouter()
	r.HandleFunc("/", test).Methods("GET")

	r.HandleFunc("/createMatch",apnaauth.CreateMatch).Methods("GET")

	//Authentication related APIs
	r.HandleFunc("/loadingPage",apnaauth.LoadingPage).Methods("GET")
	r.HandleFunc("/register",apnaauth.Register).Methods("POST")
	r.HandleFunc("/verifyUser",apnaauth.VerifyUser).Methods("GET")
	r.HandleFunc("/verifyOTPUserverification",apnaauth.VerifyOTPUserverification).Methods("POST")
	r.HandleFunc("/login",apnaauth.Login).Methods("POST")
	r.HandleFunc("/logout",apnaauth.Logout).Methods("GET")
	r.HandleFunc("/forgetPassword",apnaauth.ForgetPassword).Methods("POST")
	r.HandleFunc("/verifyOTPForgetPassword",apnaauth.VerifyOTPForgetPassword).Methods("POST")
	r.HandleFunc("/updateForgetPassword",apnaauth.UpdateForgetPassword).Methods("POST")
	r.HandleFunc("/updatePassword",apnaauth.UpdatePassword).Methods("POST")
	
	
	


	//Chat related APIs
	r.HandleFunc("/getChatDetail",apnachat.GetChatDetailroute).Methods("POST")
	r.HandleFunc("/getChat",apnachat.GetChat).Methods("POST")
	r.HandleFunc("/chat",apnachat.ChatInsert).Methods("POST")
	r.HandleFunc("/getChatAfterIndex",apnachat.GetChatAfterIndex).Methods("POST")


	//Profile related APIs
	r.HandleFunc("/profileCreation",profile.ProfileCreation).Methods("POST")
	r.HandleFunc("/getQuestion",profile.GetQuestion).Methods("POST")
	r.HandleFunc("/insertQuestion",profile.InsertQuestion).Methods("POST")
	r.HandleFunc("/updateQuestionResponse",profile.UpdateQuestionResponse).Methods("POST")
	r.HandleFunc("/getProfile",profile.GetProfile).Methods("POST")
	http.Handle("/", r)
	
	//GPS related APIs
	r.HandleFunc("/updateUserLocation",gps.UpdateUserLocation).Methods("POST")
	r.HandleFunc("/getUserListByLocation",gps.GetUserListByLocation).Methods("POST")
	r.HandleFunc("/getEventListByLocation",gps.GetEventListByLocation).Methods("POST")

	//Event related APIs
	r.HandleFunc("/createEvent",event.CreateEvent).Methods("POST")
	r.HandleFunc("/getEventById",event.GetEventById).Methods("POST")

	//Image related APIs
	r.HandleFunc("/getImageById",image.GetImageById).Methods("POST")
	r.HandleFunc("/uploadProfilePic",image.UploadProfilePic).Methods("POST")
	r.HandleFunc("/getEventImageList",image.GetEventImageList).Methods("POST")
	r.HandleFunc("/getUserImageList",image.GetUserImageList).Methods("POST")
	if err := http.ListenAndServe("0.0.0.0:9000",nil); err != nil {
		log.Fatal(err)
	}
}