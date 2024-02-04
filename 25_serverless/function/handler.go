package function

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"strings"
)

type imageType string

const (
	imageTypeBMP imageType = "bmp"
	imageTypeGIF imageType = "gif"
	imageTypePNG imageType = "png"
)

type imageConverter interface {
	Convert(i image.Image) ([]byte, error)
	ImageType() imageType
}

type UploadHandlerConfig struct {
	Bucket string
}

type uploadHandler struct {
	imageConverters []imageConverter
	s3              *S3
	config          UploadHandlerConfig
}

func NewUploadHandler(s3 *S3, config UploadHandlerConfig) *uploadHandler {
	return &uploadHandler{
		s3:     s3,
		config: config,
	}
}

func (h *uploadHandler) WithImageConverter(ic imageConverter) *uploadHandler {
	h.imageConverters = append(h.imageConverters, ic)

	return h
}

// Handle reads an object from s3 bucket, check if it's jpeg image, converts it to target formats, and saves converted images to `results` s3 bucket
func (h uploadHandler) Handle(ctx context.Context, event interface{}) error {
	ev, err := MapToStruct[Event](event.(map[string]interface{}))
	if err != nil {
		return fmt.Errorf("unable to convert event to Event type: %w", err)
	}
	// ev := event.(Event)

	objectKey := ev.ObjectKey()
	bucketName := ev.BucketName()

	if objectKey == "" || bucketName == "" {
		return fmt.Errorf("invalid event: %v", ev)
	}

	jpegBuf, err := h.s3.DownloadFile(ctx, objectKey, bucketName)
	if err != nil {
		return fmt.Errorf("unable to download file %s from %s bucket: %w", objectKey, bucketName, err)
	}

	img, err := jpeg.Decode(bytes.NewReader(jpegBuf))
	if err != nil {
		return fmt.Errorf("unable to decode jpeg: %w", err)
	}

	fileName := ev.FileNameWithoutExtension()
	// convert to target format
	for _, converter := range h.imageConverters {
		convertedBuf, err := converter.Convert(img)
		if err != nil {
			return fmt.Errorf("unable to convert image to %s format: %w", converter.ImageType(), err)
		}

		// save to s3
		objectKey := fmt.Sprintf("%s.%s", fileName, converter.ImageType())
		bucketName := h.bucketName(converter.ImageType())

		err = h.s3.UploadFile(ctx, objectKey, bucketName, convertedBuf)
		if err != nil {
			return fmt.Errorf("unable to upload image of type %s to s3 %s: %w", converter.ImageType(), bucketName, err)
		}

		fmt.Printf("Image of type %s has been uploaded to s3 %s\n", converter.ImageType(), bucketName)
	}

	return nil
}

func (h *uploadHandler) bucketName(t imageType) string {
	imageType := strings.ToLower(strings.ReplaceAll(string(t), " ", "_"))

	return fmt.Sprintf("%s/%s_images/", h.config.Bucket, imageType)
}
