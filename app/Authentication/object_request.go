package Authentication
type Login_header struct{
	Cookie string `header:"Miti-Cookie"`
}

type Verify_OTP_Header struct{
    Cookie string `header:"Miti-Cookie"`
}

type Update_password_header struct{
	Cookie string `header:"Miti-Cookie"`
}
type Logout_header struct{
	Cookie string `header:"Miti-Cookie"`
}
type Register_Header struct{
	Method string `header:"Method"`
	Agent string `header:"User-Agent"`
}