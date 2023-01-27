package s3

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Upload(b64 string) (string, error) {
	sess := session.Must(session.NewSession())

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	dec, err := base64.StdEncoding.DecodeString(b64)
	fmt.Println(err)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("go-server-test-files"),
		Key:    aws.String("filename"),
		Body:   bytes.NewReader(dec),
	})
	if err != nil {
		fmt.Println(err)
		return "404", fmt.Errorf("failed to upload file, %v", err)
	}
	fmt.Printf("file uploaded to, %s\n", result.Location)
	return result.Location, nil
}
