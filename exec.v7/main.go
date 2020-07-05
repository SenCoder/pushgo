package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
)

func main() {
	accessKey := "xxxxxxxxxxxxx"
	secretKey := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	endpoint := "http://xx.xx.xx.xx:7480" // endpoint

	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:         aws.String(endpoint),
		Region:           aws.String("US"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(false), //virtual-host style方式，不要修改
	})

	if err != nil {
		log.Fatal(err)
	}

	//service := s3.New(sess)
	uploader := s3manager.NewUploader(sess)

	fp, err := os.Open("s3_test.go")
	defer fp.Close()

	//ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	//defer cancel()

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("maa-dev-bucket"),
		Key:    aws.String("main.go"),
		Body:   fp,
	})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("upload success")
	}

}
