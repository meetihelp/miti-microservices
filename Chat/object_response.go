package Chat
type SendChatContent struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	Chat []Chat `json:"Chat"`
}

type ChatDetailContent struct{
	ChatDetailResponse []ChatDetailResponse `json:"ChatDetail"`
	Code int `json:"Code"`
	Message string `json:"Message"`
}

type ChatDetailResponse struct{
	UserId string `gorm:"varchar(100)"  json:"UserId"`
	ChatId string `gorm:"primary_key;varchar(100)"  json:"ChatId"`
	ChatType string `gorm:"varchar(100)"  json:"ChatType"`
	CreatedAt string `gorm:"varchar(100)" json:"CreatedAt"`
	LastUpdate string `gorm:"varchar(100)" json:"LastUpdate"`
}

type ChatResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	MessageId string `json:"MessageId"`
	CreatedAt string `json:"CreatedAt"`
}