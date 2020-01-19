package Privacy


import(
	// "log"
	// "fmt"
	// database "miti-microservices/Database"
// 	"github.com/jinzhu/gorm"
//  _ 	"github.com/jinzhu/gorm/dialects/postgres"
)
type Board struct{
	UserId string `gorm:"primary_key;type:varchar" json:"UserId"`
	Date string `gorm:"primary_key;type:varchar" json:"Date"`
	AccessType string `gorm:"type:varchar" json:"AccessType"`
	BoardId string `gorm:"primary_key;type:varchar" json:"BoardId"`
	CreatedAt string `gorm:"type:varchar" json:"CreatedAt"`
	UpdatedAt string `gorm:"type:varchar" json:"UpdatedAt"`//Not for content but accesstype
	RequestId string `gorm:"type:varchar" json:"RequestId"`
}

type BoardContent struct{
	UserId string `gorm:"primary_key;type:varchar" json:"UserId"`
	RequestId string `gorm:"primary_key;type:varchar" json:"RequestId"`
	ContentId string `gorm:"primary_key;type:varchar" json:"ContentId"`
	BoardId string `gorm:"type:varchar" json:"BoardId"`
	ContentText string `gorm:"type:varchar" json:"ContentText"`
	ContentImageId string `gorm:"type:varchar" json:"ContentImageId"`
	AccessType string `gorm:"type:varchar" json:"AccessType"`
	CreatedAt string `gorm:"type:varchar" json:"CreatedAt"`
	AccessRequestId string `gorm:"type:varchar" json:"AccessRequestId"`
	AccessUpdatedAt string `gorm:"type:varchar" json:"AccessUpdatedAt"`
}

func init(){
	// db:=database.GetDB()
	// db.AutoMigrate(&Board{})
	// db.AutoMigrate(&BoardContent{})
}