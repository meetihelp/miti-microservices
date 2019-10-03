package Chat

type GetChatDetail_header struct{
	Cookie string `header:"Miti-Cookie"`
}

type ChatDetailDs struct{
	Offset int `json:"offset"`
	Num_of_chat int `json:"num_of_chat"`
}
type GetChat_header struct{
	Cookie string `header:"MitiCookie"`
}
type Chat_header struct{
	Cookie string `header:"Miti-Cookie"`
}

type ChatRequest struct{
	Chat_id string `json:"ChatId"`
	Offset int `json:"Offset"`
	Num_of_chat int `json:"NumOfChat"`
}

type GetChatAfterIndex_header struct{
	Cookie string `header:"Miti-Cookie"`
}
type ChatAfterIndex struct{
	Chat_id string `json:"ChatId"`
	Offset int `json:"Offset"`
	Num_of_chat int `json:"NumOfChat"`
	Index int `json:"Index"`
}