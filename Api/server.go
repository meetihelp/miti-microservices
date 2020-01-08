package API

import (
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"net/http"
	profile "miti-microservices/Profile"
	apnaauth "miti-microservices/Authentication"
	apnachat "miti-microservices/Chat"
	util "miti-microservices/Util"
	gps "miti-microservices/GPS"
	event "miti-microservices/Event"
	newsfeed "miti-microservices/NewsFeed"
	image "miti-microservices/Media/Image"
	// sms "miti-microservices/Notification/SMS"
)

func test(w http.ResponseWriter,r *http.Request){
	util.Message(w,200)
}


func Server(){
	fmt.Println("Server running.....")
	r := mux.NewRouter()
	r.HandleFunc("/", test).Methods("GET")

	r.HandleFunc("/createMatch",apnaauth.CreateMatch).Methods("GET")

	//Authentication related APIs
	r.HandleFunc("/loadingPage",apnaauth.LoadingPage).Methods("GET")
	r.HandleFunc("/register",apnaauth.Register).Methods("POST")
	r.HandleFunc("/generateOTP",apnaauth.VerifyUser).Methods("GET")
	r.HandleFunc("/verifyOTP",apnaauth.VerifyOTPUserverification).Methods("POST")
	// r.HandleFunc("/verifyUser",apnaauth.VerifyUser).Methods("GET")
	// r.HandleFunc("/verifyOTPUserverification",apnaauth.VerifyOTPUserverification).Methods("POST")
	r.HandleFunc("/login",apnaauth.Login).Methods("POST")
	r.HandleFunc("/logout",apnaauth.Logout).Methods("GET")
	r.HandleFunc("/forgetPassword",apnaauth.ForgetPassword).Methods("POST")
	r.HandleFunc("/verifyOTPForgetPassword",apnaauth.VerifyOTPForgetPassword).Methods("POST")
	r.HandleFunc("/updateForgetPassword",apnaauth.UpdateForgetPassword).Methods("POST")
	r.HandleFunc("/updatePassword",apnaauth.UpdatePassword).Methods("POST")
	r.HandleFunc("/resendOTP",apnaauth.ReSendOTP).Methods("GET")
	r.HandleFunc("/otpStatus",apnaauth.OTPStatus).Methods("GET")
	r.HandleFunc("/getTemporaryUserId",apnaauth.GetTemporaryUserId).Methods("GET")
	
	
	


	//Chat related APIs
	r.HandleFunc("/getChatDetail",apnachat.GetChatDetailroute).Methods("POST")
	r.HandleFunc("/getChat",apnachat.GetChat).Methods("POST")
	r.HandleFunc("/chat",apnachat.ChatInsert).Methods("POST")
	r.HandleFunc("/getChatAfterIndex",apnachat.GetChatAfterIndex).Methods("POST")


	//Profile related APIs
	r.HandleFunc("/profileCreation",profile.ProfileCreation).Methods("POST")
	r.HandleFunc("/getQuestion",profile.GetQuestion).Methods("POST")
	r.HandleFunc("/insertQuestion",profile.InsertQuestion).Methods("POST")
	r.HandleFunc("/updateIPIPResponse",profile.UpdateIPIPResponse).Methods("POST")
	r.HandleFunc("/getProfile",profile.GetProfile).Methods("POST")
	r.HandleFunc("/updatePreference",profile.UpdatePreference).Methods("Post")
	http.Handle("/", r)
	
	//GPS related APIs
	r.HandleFunc("/updateUserLocation",gps.UpdateUserLocation).Methods("POST")
	r.HandleFunc("/getUserListByLocation",gps.GetUserListByLocation).Methods("POST")
	r.HandleFunc("/getEventListByLocation",gps.GetEventListByLocation).Methods("POST")

	//Event related APIs
	r.HandleFunc("/createEvent",event.CreateEvent).Methods("POST")
	r.HandleFunc("/getEventById",event.GetEventById).Methods("POST")

	// //Image related APIs
	r.HandleFunc("/uploadImage",image.UploadImage).Methods("POST")
	// r.HandleFunc("/getImageById",image.GetImageById).Methods("POST")
	// r.HandleFunc("/uploadProfilePic",image.UploadProfilePic).Methods("POST")
	// r.HandleFunc("/getEventImageList",image.GetEventImageList).Methods("POST")
	// r.HandleFunc("/getUserImageList",image.GetUserImageList).Methods("POST")


	//NewsFeed related APIs
	r.HandleFunc("/getNewsArticleList",newsfeed.GetNewsArticle).Methods("POST")
	r.HandleFunc("/getNewsArticle",newsfeed.GetNewsFeedArticle).Methods("POST")
	// r.HandleFunc("/getNewsArticle",newsfeed.GetNewsArticle).Methods("POST")
	// r.HandleFunc("/getNewsFeedSummary",newsfeed.GetNewsFeedSummary).Methods("POST")
	// r.HandleFunc("/getNewsFeedArticle",newsfeed.GetNewsFeedArticle).Methods("POST")
	// r.HandleFunc("/newsFeedReaction",newsfeed.UpdateNewsFeedReaction).Methods("POST")
	
	if err := http.ListenAndServe("0.0.0.0:8000",nil); err != nil {
		log.Fatal(err)
	}
}
