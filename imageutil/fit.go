package imageutil

import (
	"image"

	"github.com/disintegration/imaging"
)

// FitCrop fit to narrow side of image and crop square.
func FitCrop(img image.Image) image.Image {
	size := img.Bounds().Size()
	if size.X != size.Y {
		narrowSide := size.Y
		if size.X < size.Y {
			narrowSide = size.X
		}
		img = imaging.CropCenter(img, narrowSide, narrowSide)
	}

	return img
}

// FitCrop for Image struct.
func (img *Image) FitCrop() {
	img.Image = FitCrop(img.Image)
}
