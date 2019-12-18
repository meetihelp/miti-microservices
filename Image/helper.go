package Image

import(
	"os"
	"log"
	"net/http"
	util "app/Util"
	"encoding/json"
	"io"
	// "fmt"
)

const (
	ImageDir ="Image"
)
func GetImageIdFromURL(url string) string{
	// imageId:=""
	// return imageId
	return url
}

func GetImagePath(imageId string) string{
	path:=ImageDir+"/"+imageId
	return path
}

func SendImage(w http.ResponseWriter,path string){
    img, err := os.Open(path)
    if err != nil {
        log.Fatal(err) // perhaps handle this nicer
    }
    defer img.Close()
    w.Header().Set("Content-Type", "image/jpeg") 
    io.Copy(w, img)

}

func SendImageList(w http.ResponseWriter,imageList []string){
	w.Header().Set("Content-Type", "application/json")
	msg:=util.GetMessageDecode(200)
	p:=&SendImageListContent{Code:200,Message:msg,ImageList:imageList}
	enc := json.NewEncoder(w)
	err:= enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateImageFile() (string,*os.File){
	imageId:=util.GenerateToken()
	path:=ImageDir+"/"+imageId+".jpeg"
	dst, err := os.Create(path)
	if err!=nil{
		log.Println(err)
	}
	return imageId,dst
}