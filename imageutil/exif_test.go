package imageutil

import (
	"image"
	"os"
	"path/filepath"
	"testing"

	"github.com/corona10/goimagehash"
	"github.com/stretchr/testify/assert"
)

var testfiles = []string{
	"../testdata/Landscape_1.jpg",
	"../testdata/Landscape_2.jpg",
	"../testdata/Landscape_3.jpg",
	"../testdata/Landscape_4.jpg",
	"../testdata/Landscape_5.jpg",
	"../testdata/Landscape_6.jpg",
	"../testdata/Landscape_7.jpg",
	"../testdata/Landscape_8.jpg",
	"../testdata/Portrait_1.jpg",
	"../testdata/Portrait_2.jpg",
	"../testdata/Portrait_3.jpg",
	"../testdata/Portrait_4.jpg",
	"../testdata/Portrait_5.jpg",
	"../testdata/Portrait_6.jpg",
	"../testdata/Portrait_7.jpg",
	"../testdata/Portrait_8.jpg",
	"../testdata/Large_Portrait_8.jpg",
	"../testdata/Without_exif.jpg",
}

var imageHashes = []uint64{
	1731642887319092348,
	1731642887587527804,
	1731642887587527804,
	1731642887319092348,
	1731642887319092348,
	1731642887587527804,
	1731642887587527804,
	1731642887587527804,
	18442266965844736254,
	18442266965844728062,
	18442266965844736254,
	18442266965844736254,
	18442266965844736254,
	18442266965844728062,
	18442266965844736254,
	18442266965844736254,
	308498827398019326,
	308498827398019326,
}

func TestRemoveExif(t *testing.T) {
	assert := assert.New(t)

	for i, filename := range testfiles {
		p, err := filepath.Abs(filename)
		assert.NoError(err, filename)
		file, err := os.Open(p)
		assert.NoError(err, filename)

		func(file *os.File) {
			defer file.Close()
			r, err := RemoveExif(file)
			assert.NotNil(r, filename)
			assert.NoError(err, filename)

			img, _, err := image.Decode(r)
			assert.NoError(err, filename)

			hash, err := goimagehash.AverageHash(img)
			assert.NoError(err)
			expected := goimagehash.NewImageHash(imageHashes[i], goimagehash.AHash)
			distance, err := hash.Distance(expected)
			assert.NoError(err)
			assert.True(distance < 8, "filename: distance=%d", distance)
		}(file)
	}
}
