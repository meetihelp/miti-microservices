package Chat

type GetChatDetailHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type ChatDetailDs struct{
	CreatedAt string `json:"CreatedAt"`
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

type SendChatImageHeader struct{
	Cookie string `header:"Miti-Cookie"`
	AccessType string `header:"Access-Type"`
	ActualFileName string `header:"Actual-Filename"`
	Format string `header:"Format"`
	Latitude string `header:"Latitutde"`
	Longitude string `header:"Longitude"`
	Dimension string `header:"Dimension"`
	RequestId string `header:"Request-Id"`
	ChatId string `header:"Chat-Id"`
}

type SendMessageRequestHeader struct{
	Cookie string `header:"Miti-Cookie"`
}
type SendMessageRequestDS struct{
	RequestId string `json:"RequestId"`
	MessageContent string `json:"MessageContent"`
	MessageType string `json:"MessageType"`
	Phone string `json:"Phone"`
}

type GetMessageRequestHeader struct{
	Cookie string `header:"Miti-Cookie"`
}