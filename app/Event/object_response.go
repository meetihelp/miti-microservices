package Event

import()

type EventResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	EventId string `json:"EventId"`
	EventName string `json:"EventName"`
	EventPicURL string `json:"EventPicURL"`
	EventType string `json:"EventType"`
	Time string `json:"Time"`
	OrganiserId string `json:"OrganiserId"`
	Latitude string `json:"Latitude"`
	Longitude string `json:"Longitude"`
}