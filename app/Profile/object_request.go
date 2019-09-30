package Profile

type Profile_creation_Header struct{
	Method1 string `header:"method"`
	Agent1 string `header:"agent"`
	Cookie string `header:"Miti-Cookie"`
}

type InsertQuestion_Header struct{
	Method1 string `header:"method"`
	Agent1 string `header:"agent"`
	Cookie string `header:"Miti-Cookie"`
}

type UpdateQuestionResponse_Header struct{
	Method1 string `header:"method"`
	Agent1 string `header:"agent"`
	Cookie string `header:"Miti-Cookie"`	
}
type QuestionRequest struct{
	Offset int `json:"Offset"`
	NumOfQuestion int `json:"NumOfQuestion"`
}

type Response struct{
	QuestionId string `json:"QuestionID"`
	Response int `json:"Response"`
}

type ResponseWrapper struct{
	IPIP []Response `json:"IPIP"`
}