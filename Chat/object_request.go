package Chat

type GetChatDetailHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type ChatDetailDs struct{
	Offset int `json:"Offset"`
	NumOfChat int `json:"NumOfChat"`
}
type GetChatHeader struct{
	Cookie string `header:"Miti-Cookie"`
}
type ChatHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type ChatRequest struct{
	ChatId string `json:"ChatId"`
	Offset int `json:"Offset"`
	NumOfChat int `json:"NumOfChat"`
}

type GetChatAfterIndexHeader struct{
	Cookie string `header:"Miti-Cookie"`
}
type ChatAfterTime struct{
	ChatId string `json:"ChatId"`
	// Offset int `json:"Offset"`
	NumOfChat int `json:"NumOfChat"`
	// Index int `json:"Index"`
	CreatedAt string `json:"CreatedAt"`
}