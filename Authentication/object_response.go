package Authentication

//Loading Page
type LoadingResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	MoveTo int `json:"MoveTo"`
	Preference int `json:"Preference"`
}

type LoadingPageResponseHeader struct{
	ContentType string `json:"Content-Type"`
}

//Login Page
type LoginResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	MoveTo int `json:"MoveTo"`
}

type LoginResponseHeader struct{
	MitiCookie string `json:"Miti-Cookie"`
	ContentType string `json:"Content-Type"`
}

//OTPStatus
type OTPStatusResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	MoveTo int `json:"MoveTo"`
}

type OTPStatusResponseHeader struct{
	ContentType string `json:"Content-Type"`
}


//Verify User
type VerifyUserResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	MoveTo int `json:"MoveTo"`
}

type VerifyUserResponseHeader struct{
	ContentType string `json:"Content-Type"`
}

//VerifyOTP
type VerifyOTPResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	MoveTo int `json:"MoveTo"`
	Preference int `json:"Preference"`
}

type VerifyOTPResponseHeader struct{
	MitiCookie string `json:"Miti-Cookie"`
	ContentType string `json:"Content-Type"`
}

//ResendOTP
type ResendOTPResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	MoveTo int `json:"MoveTo"`	
}

type ResendOTPResponseHeader struct{
	ContentType string `json:"Content-Type"`
}

//Last











type OTPResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	MoveTo int `json:"MoveTo"`
}

type OTPResponseHeader struct{
	MitiCookie string `json:"Miti-Cookie"`
}



type TempUserResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	UserId string `json:"UserId"`
}
// type TempUserResponse struct{
// 	Code int `json:"Code"`
// 	Message string `json:"Message"`
// 	List TempUserList `json:"List"`
// }

type TempUserList struct{
	UserId string `json:"UserId"`
	ChatList []ChatListElement `json:"ChatList"`
}

type ChatListElement struct{
	ChatId string `json:"ChatId"`
	TempId string `json:"TemporaryUserId"`
}

type GetPhoneStatusResponse struct{
	PhoneStatus []int `json:"PhoneStatus"`
	Code int `json:"Code"`
	Message string `json:"Message"`
}
// type PreferenceContent struct{
// 	Code int `json:"Code"`
// 	Message string `json:"Message"`
// 	Preference int `json:"Preference"`
// }


// type LoadingToLogin struct{
// 	Code int `json:"Code"`
// 	Message string `json:"Message"`
// 	MoveTo int `json:"MoveTo"`
// }

// type LoadingToOTP struct{
// 	Code int `json:"Code"`
// 	Message string `json:"Message"`
// 	MoveTo int `json:"MoveTo"`
// }

// type LoadingToFeed struct{
// 	Code int `json:"Code"`
// 	Message string `json:"Message"`
// 	MoveTo int `json:"MoveTo"`
// }

// type LoadingToProfile struct{
// 	Code int `json:"Code"`
// 	Message string `json:"Message"`
// 	MoveTo int `json:"MoveTo"`
// }

// type LoadingToPreference struct{
// 	Code int `json:"Code"`
// 	Message string `json:"Message"`
// 	MoveTo int `json:"MoveTo"`
// 	Preference int `json:"Preference"`
// }
