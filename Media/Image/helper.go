package Image

import(
	"os"
	"log"
	"net/http"
	util "miti-microservices/Util"
	"encoding/json"
	"io"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "bytes"
    "github.com/aws/aws-sdk-go/service/cloudfront/sign"
    "time"
	"crypto/rsa"

	// "fmt"
)

const (
	ImageDir ="Image"
	S3_REGION = "us-east-2"
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

func UploadToS3(buffer []byte,filename string,bucket string,format string) (int,error){
	 s, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
    if err != nil {
        log.Fatal(err)
    }
    size:=int64(len(buffer))
    imageType:="image/"+format
    // size:=buffer.Len()
    if(bucket==""){
    	bucket="mymititestbucket"
    }
	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
        Bucket:               aws.String(bucket),
        Key:                  aws.String(filename),
        ACL:                  aws.String("private"),
        Body:                 bytes.NewReader(buffer),
        ContentLength:        aws.Int64(size),
        ContentType: aws.String(imageType),
        // ContentType:          aws.String(http.DetectContentType(buffer)),
        // ContentDisposition:   aws.String("image/png"),
        ServerSideEncryption: aws.String("AES256"),
    })
    return int(size),err
}

func getCloudFrontCredentials() (string,*rsa.PrivateKey){
	PrivKeyFileName:=os.Getenv("cloudfront_privkeyfilename")
	keyID:=os.Getenv("cloudfront_keyId")
	// privKey:=os.Getenv("cloudfront_provKey")
	privKey,_:=sign.LoadPEMPrivKeyFile(PrivKeyFileName)
	return keyID,privKey
}
func GetSignedURL(rawURL string,expireDuration time.Duration) string{
	keyID,privKey:=getCloudFrontCredentials()
	signer := sign.NewURLSigner(keyID, privKey)
	signedURL, err := signer.Sign(rawURL, time.Now().Add(expireDuration*time.Hour))
	if err != nil {
	    log.Fatalf("Failed to sign url, err: %s\n", err.Error())
	}
	return signedURL
}

func GetPublicImageBucket() string{
	bucket:=os.Getenv("PublicImageBucket")
	return bucket
}

func GetPrivateImageBucket()string{
	bucket:=os.Getenv("PrivateImageBucket")
	return bucket
}

