package s3

// import (
// 	constant "API/constant"
// 	"encoding/base64"
// 	"fmt"
// 	"os"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/s3/s3manager"
// )

// func Upload(b64 string) error {
// 	sess := session.Must(session.NewSession())

// 	// Create an uploader with the session and default options
// 	uploader := s3manager.NewUploader(sess)

// 	dec, err := base64.StdEncoding.DecodeString(b64)
// 	if err != nil {
// 		return fmt.Errorf("failed to open file %q, %v", filename, err)
// 	}

//		// Upload the file to S3.
//		result, err := uploader.Upload(&s3manager.UploadInput{
//			Bucket: aws.String(constant.AWS_BUCKET),
//			Key:    aws.String(constant.AWS_KEY),
//			Body:   f,
//		})
//		if err != nil {
//			return fmt.Errorf("failed to upload file, %v", err)
//		}
//		fmt.Printf("file uploaded to, %s\n", aws.StringValue(result.Location))
//	}
const a = "string"
