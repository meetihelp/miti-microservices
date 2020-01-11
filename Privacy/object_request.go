package Privacy

type UploadBoardContentHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type UploadBoardContentRequest struct{
	RequestId string `json:"RequestId"`
	Date string `json:"RequestId"`
	ContentText string `json:"ContentText"`
	ContentImageId string `json:"ContentImageId"`
}

type ShareBoardHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type ShareBoardRequest struct{
	RequestId string `json:"RequestId"`
	BoardId string `json:"BoardId"`
	AccessType string `json:"AccessType"`
}