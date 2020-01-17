package Social

const(
	MAX_NUMBER_OF_PAIR=5
	MAX_NUMBER_OF_TEMPORARY_MEMBER=3

)

type ChatDetail struct{
	TempUserId string `gorm:"varchar(100)"  json:"TempUserId"`
	ActualUserId string `gorm:"primary_key;varchar(100)"  json:"ActualUserId"`
	ChatId string `gorm:"primary_key;varchar(100)"  json:"ChatId"`
	ChatType string `gorm:"varchar(100)"  json:"ChatType"`
	CreatedAt string `gorm:"varchar(100)" json:"CreatedAt"`
	LastUpdate string `gorm:"varchar(100)" json:"LastUpdate"`
	UserIndex int `gorm:"type:int" json:"Index"`
	Name string `gorm:"type:varchar" json:"Name"`
}