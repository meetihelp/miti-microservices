package Chat
type SendChat_Content struct{
	Code int `json:"code"`
	Message string `json:"message"`
	Chat []Chat `json:"chat"`
}

type ChatDetail_Content struct{
	ChatDetail []ChatDetail `json:"chatdetail"`
	Code int `json:"status"`
	Message string `json:"message"`
}