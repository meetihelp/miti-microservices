package Authentication
type VerificationHeader struct{
	Cookie string `header:"Miti-Cookie"`
}
type LoginHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type VerifyOTPHeader struct{
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