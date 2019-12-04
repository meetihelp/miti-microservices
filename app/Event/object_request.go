package Event

type CreateEventHeader struct{
	Cookie string `header:"Miti-Cookie"`
}
type CreateEventDS struct{
	EventName string `json:"EventName"`
	EventPic string `json:"EventPic"`
	EventType string `json:"EventType"`
	Time string `json:"Time"`
	Latitude string `json:"Latitude"`
	Longitude string `json:"Longitude"`
}