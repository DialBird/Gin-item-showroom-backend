package main

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"os"
)

func loadDotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func sdkHandler(c *gin.Context) {
	svc := s3.New(
		session.Must(session.NewSession()),
		aws.NewConfig().WithRegion("ap-northeast-1"),
	)
	file, err := os.Open("test.jpg")
	if err != nil {
		fmt.Printf("err opening file: %s", err)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	var size = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	path := "/media/" + file.Name()

	input := &s3.PutObjectInput{
		Body:                 fileBytes,
		Bucket:               aws.String(os.Getenv("BUCKET_NAME")),
		Key:                  aws.String(path),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(fileType),
		ServerSideEncryption: aws.String("AES256"),
		Tagging:              aws.String("key1=value1&key2=value2"),
	}
	result, err := svc.PutObject(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)

	c.String(200, "Image Sended")
}

func main() {
	loadDotenv()

	se, err := mgo.Dial(os.Getenv("MONGO_HOST") + ":27017")
	if err != nil {
		panic(err)
	}
	defer se.Close()
	se.SetMode(mgo.Monotonic, true)

	r := gin.Default()

	r.GET("sdk", sdkHandler)

	SetItemRouter(r, se)

	r.Run()
}
