package Profile

type SendQuestionContent struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	Question []Question `json:"Question"`
}

type ProfileResponse struct{
	UserId string `gorm:"primary_key;type:varchar(100)"  json:"UserId"`
	Name string `gorm:"type:varchar(40)" validate:"required" json:"Name"`
	DateOfBirth string `gorm:"type:varchar(100)" validate:"required" json:"DateOfBirth"`
	Job string `gorm:"type:varchar(30)" validate:"required" "json:"Job"`
	ProfilePicURL string `gorm:"type:varchar(100)"  json:"ProfilePicURL"`
	Gender string `gorm:"type:varchar(10)" validate:"required" json:"Gender"`
	Language string `gorm:"type:varchar(20)" validate:"required" json:"Language"`
	Country string `gorm:"type:varchar(30)" json:"Country"`
	Sex string `gorm:"type:varchar(10)"  json:"Sex"`
	RelationshipStatus string `gorm:"type:varchar(10)"  json:"RelationshipStatus"`
	ParentsAddress string `gorm:"type:varchar(10)"  json:"ParentsAddress"`
}

type ProfileResponseContent struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	ProfileResponse ProfileResponse `json:"ProfileResponse"` 
}