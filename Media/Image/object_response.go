package Image

type SendImageListContent struct{
	ImageList []string `json:"ImageList"`
	Code int `json:"Code"`
	Message string `json:"Message"`
}

type UploadImageResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	ImageId string `json:"ImageId"`
	URL string `json:"URL"`
	RequestId string `json:"RequstId"`
}