package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// TODO fill these in!
const (
	S3_REGION = "us-west-1"
	S3_BUCKET = "pfplugins"
)

func main() {
	key := os.Getenv("SPACES_KEY")
	secret := os.Getenv("SPACES_SECRET")
	fmt.Println(key)

	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:    aws.String("https://sfo2.digitaloceanspaces.com"),
		Region:      aws.String(S3_REGION),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Upload
	err = AddFileToS3(s, os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
}

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func AddFileToS3(s *session.Session, fileDir string, key string) error {

	// Open the file for use
	file, err := os.Open(fileDir)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:  aws.String(S3_BUCKET),
		Key:     aws.String(key),
		Body:    bytes.NewReader(buffer),
		ACL:     aws.String("public-read"),
		Tagging: aws.String("version=0.0.1"),
	})
	return err
}
