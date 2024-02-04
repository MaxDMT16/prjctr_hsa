package main

import (
	"os"
	function "prjctr/md/25_serverless"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	awsRegion := os.Getenv("REGION") //"eu-west-1"
	bucket := os.Getenv("BUCKET")        //"prjctr-image-converter"

	s3 := function.NewS3(&aws.Config{
		Region: aws.String(awsRegion),
	})

	config := function.UploadHandlerConfig{
		Bucket: bucket,
	}

	p := function.NewUploadHandler(s3, config).
		WithImageConverter(function.NewImageToBMPConverter()).
		WithImageConverter(function.NewImageToGIFConverter()).
		WithImageConverter(function.NewImageToPNGConverter())

	lambda.Start(p.Handle)

	// tmp := `{"Records":[{"awsRegion":"eu-west-1","eventName":"ObjectCreated:Put","eventSource":"aws:s3","eventTime":"2024-02-03T17:02:31.930Z","eventVersion":"2.1","requestParameters":{"sourceIPAddress":"95.67.100.225"},"responseElements":{"x-amz-id-2":"9UUpkXeG0Q/dUHnYTusr2T0dq4PsfkrKq7LqAz0RXbyfGM31Ea6aCP0W0EFCt75FBiYMN7WCjUwI3n8TAiZe+kBqADx6tvfZ","x-amz-request-id":"RJYF2PN7XDFN9Z1B"},"s3":{"bucket":{"arn":"arn:aws:s3:::prjctr-image-converter","name":"prjctr-image-converter","ownerIdentity":{"principalId":"A1956E5E64ZSP4"}},"configurationId":"tf-s3-lambda-20240203142946382200000001","object":{"eTag":"ace5a28bee85fb9a792ce55857033e48","key":"jpeg_images/tiger-jpg.jpeg","sequencer":"0065BE71A7DD69C9F1","size":241514},"s3SchemaVersion":"1.0"},"userIdentity":{"principalId":"A1956E5E64ZSP4"}}]}`

	// var event function.Event
	// err := json.Unmarshal([]byte(tmp), &event)
	// if err != nil {
	// 	panic(err)
	// }

	// err = p.Handle(context.Background(), event)
	// if err != nil {
	// 	panic(err)
	// }
}
