package imageutil

import (
	"image/jpeg"
	"io"
)

const defaultJPEGQuality = 95

// EncodeJPEG encodes image struct to jpeg image.
func (img *Image) EncodeJPEG(writer io.Writer) error {
	quality := defaultJPEGQuality
	if img.Quality > 0 && img.Quality < quality {
		quality = img.Quality
	}
	return jpeg.Encode(writer, img.Image, &jpeg.Options{Quality: quality})
}
