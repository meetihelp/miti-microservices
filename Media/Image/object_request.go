package Image

type GetImageByIdHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type GetImageByIdDS struct{
	ImageId string `json:"ImageId"`
}

type GetUserImageListHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type GetUserImageListDS struct{
	UserId string `json:"UserId"`
}

type GetEventImageListHeader struct{
	Cookie string `header:"Miti-Cookie"`
}
type GetEventImageListDS struct{
	EventId string `json:"EventId"`
}

type UploadProfilePicHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type UploadImageHeader struct{
	Cookie string `header:"Miti-Cookie"`
	AccessType string `header:"AccessType"`
	ActualFileName string `header:"ActualFileName"`
	Format string `header:"Format"`
	Latitude string `header:"Latitutde"`
	Longitude string `header:"Longitude"`
	Dimension string `header:"Dimension"`
	RequestId string `header:"RequestId"`
}

type UploadImageResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	ImageId string `json:"ImageId"`
	URL string `json:"URL"`
}