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
	// event "miti-microservices/Event"
	newsfeed "miti-microservices/NewsFeed"
	image "miti-microservices/Media/Image"
	social "miti-microservices/Social"
	security "miti-microservices/Security"
	privacy "miti-microservices/Privacy"
	others "miti-microservices/Others"
	// sms "miti-microservices/Notification/SMS"
	"os"
)

func test(w http.ResponseWriter,r *http.Request){
	util.Message(w,200)
}


func Server(runMethod string){
	fmt.Println("Server running.....")
	r := mux.NewRouter()
	r.HandleFunc("/", test).Methods("GET")

	// r.HandleFunc("/createMatch",apnaauth.CreateMatch).Methods("GET")

	//Authentication related APIs
	r.HandleFunc("/auth/loadingPage",apnaauth.LoadingPage).Methods("GET")
	r.HandleFunc("/auth/login",apnaauth.Login).Methods("POST")
	r.HandleFunc("/auth/otpStatus",apnaauth.OTPStatus).Methods("GET")
	r.HandleFunc("/auth/generateOTP",apnaauth.VerifyUser).Methods("GET")
	r.HandleFunc("/auth/verifyOTP",apnaauth.VerifyOTPUserverification).Methods("POST")
	r.HandleFunc("/auth/resendOTP",apnaauth.ReSendOTP).Methods("GET")
	// r.HandleFunc("/register",apnaauth.Register).Methods("POST")
	// r.HandleFunc("/generateOTP",apnaauth.VerifyUser).Methods("GET")
	
	
	// r.HandleFunc("/verifyOTPUserverification",apnaauth.VerifyOTPUserverification).Methods("POST")
	
	// r.HandleFunc("/logout",apnaauth.Logout).Methods("GET")
	// r.HandleFunc("/forgetPassword",apnaauth.ForgetPassword).Methods("POST")
	// r.HandleFunc("/verifyOTPForgetPassword",apnaauth.VerifyOTPForgetPassword).Methods("POST")
	// r.HandleFunc("/updateForgetPassword",apnaauth.UpdateForgetPassword).Methods("POST")
	// r.HandleFunc("/updatePassword",apnaauth.UpdatePassword).Methods("POST")
	
	
	r.HandleFunc("/profile/getTemporaryUserId",apnaauth.GetTemporaryUserId).Methods("GET")
	// r.HandleFunc("/getPhoneStatus",apnaauth.GetPhoneStatus).Methods("POST")
	
	
	


	//Chat related APIs
	r.HandleFunc("/chat/getChatDetail",apnachat.GetChatDetailroute).Methods("POST")
	// r.HandleFunc("/getChat",apnachat.GetChat).Methods("POST")
	r.HandleFunc("/chat/chat",apnachat.ChatInsert).Methods("POST")
	r.HandleFunc("/chat/getChatAfterIndex",apnachat.GetChatAfterIndex).Methods("POST")
	r.HandleFunc("/chat/sendChatImage",apnachat.SendChatImage).Methods("POST")
	r.HandleFunc("/chat/sendMessageRequest",apnachat.SendMessageRequest).Methods("POST")
	r.HandleFunc("/chat/getMessageRequest",apnachat.GetMessageRequest).Methods("POST")
	r.HandleFunc("/chat/actionMessageRequest",apnachat.ActionMessageRequest).Methods("POST")


	//Profile related APIs
	r.HandleFunc("/profile/profileCreation",profile.ProfileCreation).Methods("POST")
	r.HandleFunc("/getQuestion",profile.GetQuestion).Methods("POST")
	r.HandleFunc("/insertQuestion",profile.InsertQuestion).Methods("POST")
	r.HandleFunc("/profile/updateIPIPResponse",profile.UpdateIPIPResponse).Methods("POST")
	r.HandleFunc("/profile/getProfile",profile.GetProfile).Methods("POST")
	r.HandleFunc("/profile/updatePreference",profile.UpdatePreference).Methods("Post")
	r.HandleFunc("/profileReaction",profile.ProfileReaction).Methods("POST")
	r.HandleFunc("/createStatus",profile.CreateStatus).Methods("POST")
	r.HandleFunc("/getStatus",profile.GetStatus).Methods("POST")
	r.HandleFunc("/checkInterest",profile.CheckInterestRouter).Methods("POST")
	r.HandleFunc("/getCheckInterest",profile.GetCheckInterestRouter).Methods("GET")
	

	//Security Related APIs
	r.HandleFunc("/security/createPrimaryTrustChain",security.CreatePrimaryTrustChain).Methods("POST")
	r.HandleFunc("/security/createSecondaryTrustChain",security.CreateSecondaryTrustChain).Methods("POST")
	r.HandleFunc("/security/deletePrimaryTrustChain",security.DeletePrimaryTrustChain).Methods("POST")
	r.HandleFunc("/security/deletePrimaryTrustChain",security.DeleteSecondaryTrustChain).Methods("POST")
	r.HandleFunc("/security/getPrimaryTrustChain",security.GetPrimaryTrustChain).Methods("GET")
	r.HandleFunc("/security/alertMessage",security.AlertMessage).Methods("POST")

	//Privacy Related APIs
	r.HandleFunc("/diary/uploadBoardContent",privacy.UploadBoardContent).Methods("POST")
	r.HandleFunc("/diary/shareBoard",privacy.ShareBoard).Methods("POST")
	r.HandleFunc("/diary/shareBoardContnet",privacy.ShareBoardContent).Methods("POST")
	r.HandleFunc("/diary/getBoardContent",privacy.GetBoardContent).Methods("POST")
	
	
	//GPS related APIs
	r.HandleFunc("/profile/updateUserLocation",gps.UpdateUserLocation).Methods("POST")
	// r.HandleFunc("/getUserListByLocation",gps.GetUserListByLocation).Methods("POST")
	// r.HandleFunc("/getEventListByLocation",gps.GetEventListByLocation).Methods("POST")
	// r.HandleFunc("/updateUserLocation",gps.UpdateUserLocation).Methods("POST")

	//Event related APIs
	// r.HandleFunc("/createEvent",event.CreateEvent).Methods("POST")
	// r.HandleFunc("/getEventById",event.GetEventById).Methods("POST")

	// //Image related APIs
	r.HandleFunc("/image/uploadImage",image.UploadImage).Methods("POST")
	r.HandleFunc("/image/getImageById",image.GetImageById).Methods("POST")
	// r.HandleFunc("/uploadProfilePic",image.UploadProfilePic).Methods("POST")
	// r.HandleFunc("/getEventImageList",image.GetEventImageList).Methods("POST")
	// r.HandleFunc("/getUserImageList",image.GetUserImageList).Methods("POST")


	//NewsFeed related APIs
	r.HandleFunc("/feed/getNewsArticleList",newsfeed.GetNewsArticle).Methods("POST")
	// r.HandleFunc("/getNewsArticle",newsfeed.GetNewsFeedArticle).Methods("POST")
	r.HandleFunc("/feed/newsFeedReaction",newsfeed.UpdateNewsFeedReaction).Methods("POST")
	r.HandleFunc("/feed/uploadNewsFeedImage",newsfeed.UploadNewsFeedImage).Methods("POST")
	// r.HandleFunc("/getNewsArticle",newsfeed.GetNewsArticle).Methods("POST")
	// r.HandleFunc("/getNewsFeedSummary",newsfeed.GetNewsFeedSummary).Methods("POST")
	// r.HandleFunc("/getNewsFeedArticle",newsfeed.GetNewsFeedArticle).Methods("POST")
	// r.HandleFunc("/newsFeedReaction",newsfeed.UpdateNewsFeedReaction).Methods("POST")

	//Social related APIs
	r.HandleFunc("/social/getPoolStatus",social.PoolStatusRouter).Methods("POST")
	r.HandleFunc("/social/getInPool",social.GetInPool).Methods("GET")
	// r.HandleFunc("/cancelPool",social.CancelPoolRouter).Methods("GET")
	r.HandleFunc("/social/getInGroupPool",social.GetInGroupPool).Methods("POST")
	r.HandleFunc("/social/groupPoolStatus",social.GroupPoolStatusRouter).Methods("POST")
	// r.HandleFunc("/cancelGroupPool",social.CancelGroupPoolRouter).Methods("POST")

	r.HandleFunc("/deleteProfile",others.DeleteProfile).Methods("POST")
	
	http.Handle("/", r)
	certificates:=os.Getenv("SSLCertificatePath")
	crt:=certificates+"/all.crt"
	key:=certificates+"/private.key"
	if(runMethod=="Devlopment"){
		port:=os.Getenv("DevlopmentPort")
		url:="0.0.0.0:"+port
		if err := http.ListenAndServeTLS(url,crt,key,nil); err != nil {
			log.Fatal(err)
		}
	}else if(runMethod=="production"){
		port:=os.Getenv("ProductionPort")
		url:="0.0.0.0:"+port
		if err := http.ListenAndServeTLS(url,crt,key,nil); err != nil {
			log.Fatal(err)
		}
	}else{
		log.Println("Run Method not correct")
		return
	}
	
}
