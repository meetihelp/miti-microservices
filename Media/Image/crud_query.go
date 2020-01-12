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

func IsUserPermittedToSeeImage(userId string,imageId string) (string,string){
	db:=database.GetDB()
	userImage:=UserImage{}
	db.Where("image_id=?",imageId).Find(&userImage)
	if(userImage.AccessType=="Public"){
		return userImage.UserId,"Ok"
	}else{
		// Write Code to check if user is permitted
		return userImage.UserId,"Ok"
	}
}

func GetImageURL(userId string,imageId string) string{
	db:=database.GetDB()
	userImage:=UserImage{}
	db.Where("user_id=? AND image_id=?",userId,imageId).Find(&userImage)
	filename:=userImage.GeneratedName+"."+userImage.Format
	if(userImage.AccessType=="Private"){
		signedURL:=GetSignedURL(filename,10)
		return signedURL
	}else if(userImage.AccessType=="Public"){
		url:=GetPublicImageURL(filename)
		return url
	}else{
		return ""
	}
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

func GetUserImageByRequestId(userId string,requestId string) (UserImage,string){
	db:=database.GetDB()
	userImage:=UserImage{}
	db.Where("user_id=? AND request_id=?",userId,requestId).Find(&userImage)
	status:=""
	if(userImage.ImageId==""){
		status="Error"
	}else{
		status="Ok"
	}
	return userImage,status
}