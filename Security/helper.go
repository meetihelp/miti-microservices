package Security

type Profile struct{
	UserId string `json:"UserId"`
	Name string `json:"Name"`
	DateOfBirth string `json:"DateOfBirth"`
	Job string `"json:"Job"`
	ProfilePicID string `json:"ProfilePicId"`
	// ProfilePicURL string `gorm:"type:varchar(100)"  json:"ProfilePicURl"`
	Gender string `json:"Gender"`
	Sex string `json:"Sex"`
	RelationshipStatus string `json:"RelationshipStatus"`
	ParentsAddress string `json:"ParentsAddress"`
	Language string `json:"Language"`
	Country string `json:"Country"`
	Extraversion int `json:"Extraversion"`
	Agreeableness int `json:"Agreeableness"`
	Conscientiousness int `json:"Conscientiousness"`
	EmotionalStability int `json:"EmotionalStability"`
	Intellect int `json:"Intellect"`
	MakingChoice string `json:"MakingChoice"`
	Pincode string `json:"Pincode"`
	// Latitude string `gorm:"type:varchar(20)" json:"Latitude"`
	// Longitude string `gorm:"type:varchar(20)" json:"Longitude"`
}