package Image
import(
	"os"
	"github.com/jinzhu/gorm"
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

func GetUserImageListDB(db *gorm.DB,userId string) ([]string){
	// db:=database.GetDB()
	userImage:=[]UserImage{}
	db.Where("user_id=?",userId).Find(&userImage)
	var imageList []string
	for _,u:=range userImage{
		imageList=append(imageList,u.ImageId)
	}
	return imageList
}

func IsUserPermittedToSeeImage(db *gorm.DB,userId string,imageId string) (string,string,bool){
	userImage:=UserImage{}
	err:=db.Where("image_id=?",imageId).Find(&userImage).Error
	if(err!=nil){
		return userImage.UserId,"Error",false
	}
	if(userImage.AccessType=="public"){
		return userImage.UserId,"Ok",true
	}else{
		// Write Code to check if user is permitted
		if(userImage.UserId==userId){
			return userImage.UserId,"Ok",true	
		}else{
			return userImage.UserId,"Error",true
		}
		
	}
}

func GetImageURL(db *gorm.DB,userId string,imageId string) (string,bool){
	userImage:=UserImage{}
	db.Where("user_id=? AND image_id=?",userId,imageId).Find(&userImage)
	filename:=userImage.GeneratedName+"."+userImage.Format
	if(userImage.AccessType=="private"){
		signedURL:=GetSignedURL(filename,10)
		return signedURL,false
	}else if(userImage.AccessType=="public"){
		url:=GetPublicImageURL(filename)
		return url,false
	}else{
		return "",false
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

func InsertUserImage(db *gorm.DB,userId string,imageId string){
	// db:=database.GetDB()
	userImage:=UserImage{}
	userImage.UserId=userId
	userImage.ImageId=imageId
	db.Create(&userImage)
}

func EnterUserImage(db *gorm.DB,userImageData UserImage) bool{
	// db:=database.GetDB()
	err:=db.Create(&userImageData).Error
	if(err!=nil){
		return true
	}
	return false
}

func GetUserImageByRequestId(db *gorm.DB,userId string,requestId string) (UserImage,string,bool){
	// db:=database.GetDB()
	userImage:=UserImage{}
	err:=db.Where("user_id=? AND request_id=?",userId,requestId).Find(&userImage).Error
	if(gorm.IsRecordNotFoundError(err)){
		return userImage,"Error",false
	}
	if(err!=nil){
		return userImage,"Error",true
	}
	return userImage,"Ok",false
}