package Image
import(
	"os"
	database "miti-microservices/Database"
)
func DoesImageExist(path string) string{
	if _, err := os.Stat(path); err != nil {
        if os.IsNotExist(err) {
            return "Error"
        }
    }
    return "Ok"
}

func GetUserImageListDB(userId string) ([]string){
	db:=database.GetDB()
	userImage:=[]UserImage{}
	db.Where("user_id=?",userId).Find(&userImage)
	var imageList []string
	for _,u:=range userImage{
		imageList=append(imageList,u.ImageId)
	}
	return imageList
}

func GetEventImageListDB(eventId string) ([]string){
	db:=database.GetDB()
	eventImage:=[]EventImage{}
	db.Where("event_id=?",eventId).Find(&eventImage)
	var imageList []string
	for _,e:=range eventImage{
		imageList=append(imageList,e.ImageId)
	}
	return imageList
}

func InsertUserImage(userId string,imageId string){
	db:=database.GetDB()
	userImage:=UserImage{}
	userImage.UserId=userId
	userImage.ImageId=imageId
	db.Create(&userImage)
}

func EnterUserImage(userImageData UserImage){
	db:=database.GetDB()
	db.Create(&userImageData)
}