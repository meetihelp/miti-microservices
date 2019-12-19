package Image

import(
	"net/http"
	"fmt"
	// CD "miti-microservices/Model/CreateDatabase"
	util "miti-microservices/Util"
	// "io/ioutil"
	"io"
	// "encoding/json"
	profile "miti-microservices/Profile"
)

func UploadProfilePic(w http.ResponseWriter,r *http.Request){
	uploadProfilePicHeader:=UploadProfilePicHeader{}
	util.GetHeader(r,&uploadProfilePicHeader)
	sessionId:=uploadProfilePicHeader.Cookie
	userId,getChatStatus:=util.GetUserIdFromSession(sessionId)
	// fmt.Println(userId)
	if getChatStatus=="Error"{
		util.Message(w,1003)
		return
	}
	err := r.ParseMultipartForm(100000)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    m := r.MultipartForm
    files := m.File["myfiles"]
    for i, _ := range files {
        //for each fileheader, get a handle to the actual file
        file, err := files[i].Open()
        defer file.Close()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        //create destination file making sure the path is writeable.
        // dst, err := os.Create("/home/sanat/" + files[i].Filename)
        // imageId,dst:=CreateImageFile(files[i].Filename)
        fmt.Println(files[i].Filename)
        imageId,dst:=CreateImageFile()
        profile.UpdateProfilePicDB(userId,imageId)
        InsertUserImage(userId,imageId)
        if _, err := io.Copy(dst, file); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

    }

}