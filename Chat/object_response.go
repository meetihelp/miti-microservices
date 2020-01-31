package Chat

type GetChatDetailResponseHeader struct{
	ContentType string `json:"Content-Type"`
}

type GetChatResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	Chat []Chat `json:"Chat"`
}

type GetChatResponseHeader struct{
	ContentType string `json:"Content-Type"`
}

type ChatResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	MessageId string `json:"MessageId"`
	RequestId string `json:"RequestId"`
	CreatedAt string `json:"CreatedAt"`
	MessageType string `json:""MessageType`
	Chat []Chat `json:"Chat"`
}

type ChatResponseHeader struct{
	ContentType string `json:"Content-Type"`
}






type SendChatResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	Chat []Chat `json:"Chat"`
}

type ChatDetailResponse struct{
	ChatDetailContent []ChatDetailContent `json:"ChatDetail"`
	Code int `json:"Code"`
	Message string `json:"Message"`
}


type ChatDetailContent struct{
	UserId string `gorm:"varchar(100)"  json:"UserId"`
	UserId2 string `json:"UserId2"`
	ChatId string `gorm:"primary_key;varchar(100)"  json:"ChatId"`
	ChatType string `gorm:"varchar(100)"  json:"ChatType"`
	CreatedAt string `gorm:"varchar(100)" json:"CreatedAt"`
	LastUpdate string `gorm:"varchar(100)" json:"LastUpdate"`
	Name string `json:"Name"`
}



type SendChatImageResponse struct{
	Code int `json:"Code`
	Message string `json:"Message"`
	ImageId string `json:"ImageId"`
	MessageId string `json:"MessageId"`
	RequestId string `json:"RequestId"`
	CreatedAt string `json:"CreatedAt"`
	MessageType string `json:""MessageType`
	URL string `json:"URL"`
	Chat []Chat `json:"Chat"`
}

type SendChatImageResponseHeader struct{
	ContentType string `json:"Content-Type"`
}

type SendMessageRequestResponse struct{
	Code int `json:"Code`
	Message string `json:"Message"`
	RequestId string `json:"RequestId"`
	CreatedAt string `json:"CreatedAt"`
}

type SendMessageRequestResponseHeader struct{
	ContentType string `json:"Content-Type"`
}
type GetMessageRequestResponseHeader struct{
	ContentType string `json:"Content-Type"`
}

type GetMessageRequestResponse struct{
	Code int `json:"Code`
	Message string `json:"Message"`
	RequestList []MessageRequestDS `json:"RequestList"`
}

type ActionMessageRequestResponseHeader struct{
	ContentType string `json:"Content-Type"`
}
type MessageRequestDS struct{
	// UserId string `json:"UserId"`
	Name string `json:"Name"`
	Phone string `json:"Phone"`
	CreatedAt string `json:"CreatedAt"`
}

type ActionMessageRequestResponse struct{
	Code int `json:"Code`
	Message string `json:"Message"`
	RequestId string `json:"RequestId"`
	CreatedAt string `json:"CreatedAt"`
}