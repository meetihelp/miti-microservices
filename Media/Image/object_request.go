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
	AccessType string `header:"Access-Type"`
	ActualFileName string `header:"Actual-Filename"`
	Format string `header:"Format"`
	Latitude string `header:"Latitutde"`
	Longitude string `header:"Longitude"`
	Dimension string `header:"Dimension"`
	RequestId string `header:"RequestId"`
}

