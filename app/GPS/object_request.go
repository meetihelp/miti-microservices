package GPS

type UpdateUserLocationHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type GetUserListByLocationHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type GetUserListByLocationDS struct{
	Latitude string `json:"Latitude"`
	Longitude string `json:"Longitude"`
	Distance string `json:"Distance"`
}