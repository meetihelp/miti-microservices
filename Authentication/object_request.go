package Authentication

//Loading Page
type LoadingPageHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

//Login
type LoginHeader struct{
	Cookie string `header:"Miti-Cookie"`
}
type LoginRequest struct{
	Phone string `validate:"required" json:"Phone"`
}

//OTPStatus
type OTPStatusHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

//VerifyUser
type VerifyUserHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

//VerifyOTP
type VerifyOTPHeader struct{
    Cookie string `header:"Miti-Cookie"`
}

type VerifyOTPRequest struct{
	OTP string `validate:"required" json:"OTP"`
}

//ResendOTP
type ResendOTPHeader struct{
	Cookie string `header:"Miti-Cookie"`	
}

//Get TempUser Id
type GetTempUserIdHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

//Last












type VerificationHeader struct{
	Cookie string `header:"Miti-Cookie"`
}



type UpdatePasswordHeader struct{
	Cookie string `header:"Miti-Cookie"`
}
type LogoutHeader struct{
	Cookie string `header:"Miti-Cookie"`
}
type RegisterHeader struct{
	Method string `header:"Method"`
	Agent string `header:"UserAgent"`
}

type SMSHeader struct{
	Cookie string `header:"Miti-Cookie"`
}



type GetPhoneStatusHeader struct{
	Cookie string `header:"Miti-Cookie"`	
}

type GetPhoneStatusRequest struct{
	PhoneList []PhoneDS `json:"PhoneList"`
}

type PhoneDS struct{
	Phone string `json:"Phone"`
}