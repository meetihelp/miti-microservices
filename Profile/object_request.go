package Profile

type ProfileCreationHeader struct{
	Method1 string `header:"method"`
	Agent1 string `header:"agent"`
	Cookie string `header:"Miti-Cookie"`
}

type GetProfileHeader struct{
	Cookie string `header:"Miti-Cookie"`	
}

type InsertQuestionHeader struct{
	Method1 string `header:"method"`
	Agent1 string `header:"agent"`
	Cookie string `header:"Miti-Cookie"`
}

type UpdateQuestionResponseHeader struct{
	Method1 string `header:"method"`
	Agent1 string `header:"agent"`
	Cookie string `header:"Miti-Cookie"`	
}
type UpdatePreferenceResponseHeader struct{
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

type ProfileRequest struct{
	UserId string `json:"UserId"`
}

type PreferenceRequest struct{
	Page int `json:"Page"`
	I1 string `json:"I1"`
	I2 string `json:"I2"`
}

type ProfileReactionHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type ProfileReactionRequest struct{
	UserId string `json:"UserId"`
	Reaction string `json:"Reaction"`
}

type CreateStatusHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type GetStatusHeader struct{
	Cookie string `header:"Miti-Cookie"`	
}

type GetStatusRequest struct{
	ChatId string `json:"ChatId"`
}