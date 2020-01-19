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

type GetBoardContentResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	BoardContentList []BoardContentList `json:"BoardContentList"`
}

type BoardContentList struct{
	UserId string `gorm:"primary_key;type:varchar" json:"UserId"`
	ContentId string `gorm:"primary_key;type:varchar" json:"ContentId"`
	BoardId string `gorm:"type:varchar" json:"BoardId"`
	ContentText string `gorm:"type:varchar" json:"ContentText"`
	ContentImageId string `gorm:"type:varchar" json:"ContentImageId"`
	AccessType string `gorm:"type:varchar" json:"AccessType"`
	CreatedAt string `gorm:"type:varchar" json:"CreatedAt"`
}