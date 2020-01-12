package Privacy

type UploadBoardContentResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	CreatedAt string `json:"CreatedAt"`
	RequestId string `json:"RequestId"`
	BoardId string `json:"BoardId"`
	ContentId string `json:"ContentId"`
}

type ShareBoardResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	RequestId string `json:"RequestId"`
	UpdatedAt string `json:"UpdatedAt"`
}

type ShareBoardContentResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	AccessRequestId string `json:"AccessRequestId"`
	AccessUpdatedAt string `json:"AccessUpdatedAt"`
}