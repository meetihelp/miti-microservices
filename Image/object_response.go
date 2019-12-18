package Image

type SendImageListContent struct{
	ImageList []string `json:"ImageList"`
	Code int `json:"Code"`
	Message string `json:"Message"`
}