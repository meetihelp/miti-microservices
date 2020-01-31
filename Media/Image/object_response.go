package Image

//Upload Image
type UploadImageResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	ImageId string `json:"ImageId"`
	URL string `json:"URL"`
	RequestId string `json:"RequstId"`
	CreatedAt string `json:"CreatedAt"`
}

type UploadImageResponseHeader struct{
	ContentType string `json:"Content-Type"`
}

//Get image by id
type GetImageByIdResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	ImageURL string `json:"ImageURL"`
}

type GetImageByIdResponseHeader struct{
	ContentType string `json:"Content-Type"`	
}

//Last

type SendImageListContent struct{
	ImageList []string `json:"ImageList"`
	Code int `json:"Code"`
	Message string `json:"Message"`
}



