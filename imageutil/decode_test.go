package imageutil

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode_JPEGQuality(t *testing.T) {
	assert := assert.New(t)

	testFiles := []string{
		"../testdata/garden.jpg",
		"../testdata/garden.png",
	}

	expectedQuality := []int{71, 0}

	for i, fileName := range testFiles {
		filePath, err := filepath.Abs(fileName)
		assert.NoError(err)

		file, err := os.Open(filePath)
		assert.NoError(err)
		defer file.Close()

		img, err := Decode(file)
		assert.NoError(err)

		assert.Equal(expectedQuality[i], img.Quality, fileName)
	}
}
