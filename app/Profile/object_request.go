package Profile

type Profile_creation_Header struct{
	Method1 string `header:"method"`
	Agent1 string `header:"agent"`
	Cookie string `header:"Miti-Cookie"`
}