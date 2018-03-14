package imageutil

import (
	"image"
	"os"
	"path/filepath"
	"testing"

	"github.com/corona10/goimagehash"
	"github.com/stretchr/testify/assert"
)

func TestFitCrop_HeightGreaterThanWidth(t *testing.T) {
	assert := assert.New(t)

	var file *os.File
	{
		filePath, err := filepath.Abs("../testdata/FitCrop_HeightGreaterThanWidth.jpg")
		assert.NoError(err)
		file, err = os.Open(filePath)
		defer file.Close()
		assert.NoError(err)
	}

	img, _, err := image.Decode(file)
	assert.NoError(err)

	img = FitCrop(img)

	bounds := img.Bounds()
	assert.Equal(int(690), bounds.Dx())
	assert.Equal(int(690), bounds.Dy())

	hash, err := goimagehash.AverageHash(img)
	assert.NoError(err)
	expected := goimagehash.NewImageHash(4111518241746910974, goimagehash.AHash)
	distance, err := hash.Distance(expected)
	assert.NoError(err)
	assert.True(distance < 2, "filename: distance=%d", distance)
}

func TestFitCrop_WidthGreaterThanHeight(t *testing.T) {
	assert := assert.New(t)

	var file *os.File
	{
		filePath, err := filepath.Abs("../testdata/FitCrop_WidthGreaterThanHeight.jpg")
		assert.NoError(err)
		file, err = os.Open(filePath)
		defer file.Close()
		assert.NoError(err)
	}

	img, _, err := image.Decode(file)
	assert.NoError(err)

	img = FitCrop(img)

	bounds := img.Bounds()
	assert.Equal(int(1151), bounds.Dx())
	assert.Equal(int(1151), bounds.Dy())

	hash, err := goimagehash.AverageHash(img)
	assert.NoError(err)
	expected := goimagehash.NewImageHash(33776998211911551, goimagehash.AHash)
	distance, err := hash.Distance(expected)
	assert.NoError(err)
	assert.True(distance < 2, "filename: distance=%d", distance)
}

func TestFitCrop_SameSizeWidthHeight(t *testing.T) {
	assert := assert.New(t)

	var file *os.File
	{
		filePath, err := filepath.Abs("../testdata/FitCrop_SameSizeWidthHeight.jpg")
		assert.NoError(err)
		file, err = os.Open(filePath)
		defer file.Close()
		assert.NoError(err)
	}

	img, _, err := image.Decode(file)
	assert.NoError(err)

	img = FitCrop(img)

	bounds := img.Bounds()
	assert.Equal(int(1150), bounds.Dx())
	assert.Equal(int(1150), bounds.Dy())

	hash, err := goimagehash.AverageHash(img)
	assert.NoError(err)
	expected := goimagehash.NewImageHash(1099511627775, goimagehash.AHash)
	distance, err := hash.Distance(expected)
	assert.NoError(err)
	assert.True(distance < 2, "filename: distance=%d", distance)
}
