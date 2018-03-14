// Package imageutil wrapper for jpeg.DecodeWithQuality & remove exif, apply orientation.
package imageutil

import (
	"bytes"
	"errors"
	"image"
	"io"

	"github.com/go-oss/image/jpeg"
	"github.com/rwcarlsen/goexif/exif"

	// use webp decoder
	_ "golang.org/x/image/webp"
)

var (
	// ErrUnsupportedFileType represents unsupported error.
	ErrUnsupportedFileType = errors.New("unsupported file type")
)

// Image represents image.Image wrapper.
type Image struct {
	Image            image.Image
	Quality          int
	OriginalMIMEType MIMEType
}

// Decode image file.
func Decode(file io.Reader) (*Image, error) {
	fileHeader := make([]byte, 512)
	_, err := file.Read(fileHeader)
	if err != nil {
		return nil, err
	}

	multiReader := io.MultiReader(bytes.NewBuffer(fileHeader), file)

	mimeType := DetectMimeType(fileHeader)
	if mimeType == UndefinedType {
		return nil, ErrUnsupportedFileType
	}

	buff := new(bytes.Buffer)
	reader := io.TeeReader(multiReader, buff)

	img := &Image{
		OriginalMIMEType: mimeType,
	}

	if mimeType == JPEG {
		img.Image, img.Quality, err = jpeg.DecodeWithQuality(reader)
		if err != nil {
			return nil, err
		}

		// check EXIF and apply orientation
		meta, decodeErr := exif.Decode(buff)
		// if error occured, exif is invalid and discarded.
		if decodeErr == nil {
			newImg, _ := applyOrientation(img.Image, meta)
			if newImg != nil {
				img.Image = newImg
			}
		}
	} else {
		img.Image, _, err = image.Decode(reader)
		if err != nil {
			return nil, err
		}
	}

	return img, nil
}
