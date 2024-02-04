package function

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/png"

	"golang.org/x/image/bmp"
)

type ImageToBMPConverter struct{}

func NewImageToBMPConverter() *ImageToBMPConverter {
	return &ImageToBMPConverter{}
}

func (c *ImageToBMPConverter) Convert(img image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bmp.Encode(buf, img); err != nil {
		return nil, fmt.Errorf("unable to encode png: %w", err)
	}

	return buf.Bytes(), nil
}

func (c ImageToBMPConverter) ImageType() imageType {
	return imageTypeBMP
}

type ImageToGIFConverter struct{}

func NewImageToGIFConverter() *ImageToGIFConverter {
	return &ImageToGIFConverter{}
}

func (c *ImageToGIFConverter) Convert(img image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := gif.Encode(buf, img, nil); err != nil {
		return nil, fmt.Errorf("unable to encode png: %w", err)
	}

	return buf.Bytes(), nil
}

func (c ImageToGIFConverter) ImageType() imageType {
	return imageTypeGIF
}

type ImageToPNGConverter struct{}

func NewImageToPNGConverter() *ImageToPNGConverter {
	return &ImageToPNGConverter{}
}

func (c *ImageToPNGConverter) Convert(img image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return nil, fmt.Errorf("unable to encode png: %w", err)
	}

	return buf.Bytes(), nil
}

func (c ImageToPNGConverter) ImageType() imageType {
	return imageTypePNG
}
