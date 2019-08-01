package main

import (
	// "encoding/json"
	"log"
	// "net/http"
	"github.com/gorilla/mux"
	 "fmt"
 "os"
 "bytes"
 "net/http"
 // "net/url"
  "github.com/aws/aws-sdk-go/aws" 
  "github.com/aws/aws-sdk-go/aws/awsutil" 
  "github.com/aws/aws-sdk-go/aws/credentials" 
  "github.com/aws/aws-sdk-go/service/s3" 
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/google/uuid"
)

func upload(ch chan <- string){
	aws_access_key_id := os.Getenv("aws_access_key_id")
	aws_secret_access_key :=os.Getenv("aws_secret_access_key")
  token := "" 
  creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token) 
    _, err := creds.Get() 
  if err != nil { 
    // handle error
    fmt.Printf(err.Error())
  }  
   cfg := aws.NewConfig().WithRegion("ap-south-1").WithCredentials(creds)
    svc := s3.New(session.New(), cfg) 
  
  file, err := os.Open("upload.png") 
    if err != nil { 
    // handle error
        fmt.Printf("file nhi mil")
  }  
   defer file.Close()  
   fileInfo, _ := file.Stat()  
    size := fileInfo.Size()
       buffer := make([]byte, size) // read file content to buffer 
 
  file.Read(buffer) 
    fileBytes := bytes.NewReader(buffer) 
  fileType := http.DetectContentType(buffer) 
  id, err := uuid.NewUUID()
    if err !=nil {
        // handle error
        fmt.Printf(err.Error())
    }
  path := id.String()+".png" 
    params := &s3.PutObjectInput{ 
    Bucket: aws.String("miti-testbucket"), 
    Key: aws.String(path), 
    Body: fileBytes, 
    ContentLength: aws.Int64(size), 
    ContentType: aws.String(fileType), 
  } 
    resp, err := svc.PutObject(params) 
  if err != nil { 
    // handle error
    fmt.Printf(err.Error())
  }  
   // fmt.Printf("response %s", awsutil.StringValue(resp))
  ch <- awsutil.StringValue(resp)
  
}

func listProducts(w http.ResponseWriter, r *http.Request) {
	// list all products
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/html")
	ch :=make(chan string)
	go upload(ch)
	etag := <-ch
	defer close(ch)
	fmt.Printf(etag)
	w.Write([]byte(etag))

}




func main() {
	r := mux.NewRouter()
	// // match only GET requests on /product/

	r.HandleFunc("/", listProducts).Methods("GET")

	// 	// handle all requests with the Gorilla router.
	http.Handle("/", r)
	if err := http.ListenAndServe("0.0.0.0:8085", nil); err != nil {
		log.Fatal(err)
	}
}
