package imageutil

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"io"

	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
)

// RemoveExif remove JPEG EXIF metadata and rotate image for orientation tag.
func RemoveExif(r io.Reader) (io.Reader, error) {
	buff := new(bytes.Buffer)
	defer buff.Reset()
	r = io.TeeReader(r, buff)

	img, err := jpeg.Decode(r)
	if err != nil {
		return nil, err
	}

	// check EXIF and apply orientation
	meta, err := exif.Decode(buff)
	if err == nil {
		// エラーが発生した場合は入力画像の EXIF が不正なだけなので無視する
		newImg, _ := applyOrientation(img, meta)
		if newImg != nil {
			img = newImg
		}
	}

	// Remove EXIF metadata
	reader, writer := io.Pipe()
	go func() {
		var err error
		defer func() {
			if cause := recover(); cause != nil && err == nil {
				err = fmt.Errorf("panic: %+v", cause)
			}
			writer.CloseWithError(err)
		}()
		err = jpeg.Encode(writer, img, &jpeg.Options{Quality: 100})
	}()
	return reader, nil
}

func applyOrientation(img image.Image, meta *exif.Exif) (image.Image, error) {
	tag, err := meta.Get(exif.Orientation)
	if err != nil {
		return nil, err
	}
	orientation, err := tag.Int(0)
	if err != nil {
		return nil, err
	}
	if orientation < 1 || orientation > 8 {
		return nil, errors.New("invalid orientation")
	}

	// ref: http://www.exif.org/Exif2-2.PDF
	// ref: http://www.cipa.jp/std/documents/j/DC-008-2012_J.pdf
	switch orientation {
	case 1: // The 0th row is at the visual top of the image, and the 0th column is the visual left-hand side.
		return img, nil
	case 2: // The 0th row is at the visual top of the image, and the 0th column is the visual right-hand side.
		img = imaging.FlipH(img)
	case 3: // The 0th row is at the visual bottom of the image, and the 0th column is the visual right-hand side.
		img = imaging.FlipH(img)
		img = imaging.FlipV(img)
	case 4: // The 0th row is at the visual bottom of the image, and the 0th column is the visual left-hand side.
		img = imaging.FlipV(img)
	case 5: // The 0th row is the visual left-hand side of the image, and the 0th column is the visual top.
		img = imaging.Rotate270(img)
		img = imaging.FlipH(img)
	case 6: // The 0th row is the visual right-hand side of the image, and the 0th column is the visual top.
		img = imaging.Rotate270(img)
	case 7: // The 0th row is the visual right-hand side of the image, and the 0th column is the visual bottom.
		img = imaging.Rotate90(img)
		img = imaging.FlipH(img)
	case 8: // The 0th row is the visual left-hand side of the image, and the 0th column is the visual bottom.
		img = imaging.Rotate90(img)
	}

	return img, nil
}
