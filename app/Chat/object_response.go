package Chat
type SendChatContent struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	Chat []Chat `json:"Chat"`
}

type ChatDetailContent struct{
	ChatDetail []ChatDetail `json:"ChatDetail"`
	Code int `json:"Code"`
	Message string `json:"Message"`
}