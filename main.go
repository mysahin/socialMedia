package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"socialMedia/Database"
	"socialMedia/Router"
)

func main() {
	Database.Connect()

	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			accessKey,
			secretKey,
			"",
		),
	})

	if err != nil {
		panic(err)
	}

	uploader := s3manager.NewUploader(awsSession)
	downloader := s3.New(awsSession)

	app := Router.Routes(uploader, downloader, bucketName)

	err = app.Listen(":8000")
	if err != nil {
		panic(err)
	}
}

func init() {
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			accessKey,
			secretKey,
			"",
		),
	})

	if err != nil {
		panic(err)
	}

	uploader = s3manager.NewUploader(awsSession)
	downloader = s3.New(awsSession)
}

var region string = "eu-north-1"
var accessKey string = "AKIAQ3EGTZABT7PRWHUS"
var secretKey string = "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"
var bucketName string = "social-media-mysahin"
var uploader *s3manager.Uploader
var downloader *s3.S3
