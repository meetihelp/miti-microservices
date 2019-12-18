package GPS
import()

type SendUserListContent struct{
	UserList []string `json:"UserList"`
	Code int `json:"Code"`
	Message string `json:"Message"`
}

type SendEventListContent struct{
	EventList []string `json:"EventList"`
	Code int `json:"Code"`
	Message string `json:"Message"`
}