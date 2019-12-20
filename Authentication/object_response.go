package Authentication


type LoadingResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	MoveTo int `json:"MoveTo"`
	Preference int `json:"Preference"`
}

type LoadingToLoginHeader struct{
	ContentType string `json:"Content-Type"`
}
type LoadingToOTPHeader struct{
	ContentType string `json:"Content-Type"`
}
type LoadingToProfileHeader struct{
	ContentType string `json:"Content-Type"`
}
type LoadingToPreferenceHeader struct{
	ContentType string `json:"Content-Type"`
}
type LoadingToFeedHeader struct{
	ContentType string `json:"Content-Type"`
}

type LoginResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	MoveTo int `json:"MoveTo"`
}

type LoginToOTPHeader struct{
	MitiCookie string `json:"Miti-Cookie"`
}

type OTPResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	MoveTo int `json:"MoveTo"`
}

type OTPResponseHeader struct{
	MitiCookie string `json:"Miti-Cookie"`
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
