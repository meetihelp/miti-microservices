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