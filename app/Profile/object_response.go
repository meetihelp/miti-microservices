package Profile

type SendQuestion_Content struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	Question []Question `json:"Question"`
}

