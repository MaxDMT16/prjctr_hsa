package function

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3 struct {
	session *session.Session
}

func NewS3(config *aws.Config) *S3 {
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}

	return &S3{
		session: sess,
	}
}

func (s *S3) DownloadFile(ctx context.Context, fileName, bucketName string) ([]byte, error) {
	downloader := s3manager.NewDownloader(s.session)

	file := aws.WriteAtBuffer{}

	_, err := downloader.Download(&file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(fileName),
		})
	if err != nil {
		return nil, fmt.Errorf("unable to download file %s from bucket %s: %w", fileName, bucketName, err)
	}

	return file.Bytes(), nil
}

func (s *S3) UploadFile(ctx context.Context, fileName, bucketName string, file []byte) error {
	uploader := s3manager.NewUploader(s.session)

	fmt.Printf("uploading file - bucket: %s\tfile: %s\n", bucketName, fileName)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(file),
	})

	if err != nil {
		return fmt.Errorf("unable to upload %s to bucket %s: %w", fileName, bucketName, err)
	}

	return nil
}
