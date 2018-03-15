#  image

`image/jpeg` base pure go image utility.

[![GoDoc](https://godoc.org/github.com/go-oss/image?status.svg)](https://godoc.org/github.com/go-oss/image)
[![Build Status](https://travis-ci.org/go-oss/image.svg?branch=master)](https://travis-ci.org/go-oss/image)
[![codecov](https://codecov.io/gh/go-oss/image/branch/master/graph/badge.svg)](https://codecov.io/gh/go-oss/image)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-oss/image)](https://goreportcard.com/report/github.com/go-oss/image)

## features
- Decode jpeg with quality value.
- Apply EXIF orientation to image pixels.

## usage
To install `imageutil` package:

```
go get github.com/go-oss/image/imageutil
```

Or, just use `jpeg` package:

```
go get github.com/go-oss/image/jpeg
```

### example
```
reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
img, err := imageutil.Decode(reader)
if err != nil {
    log.Fatal(err)
}

// print metadata
fmt.Printf("Original MIME Type: %s", img.OriginalMIMEType)
fmt.Printf("Original Image Quality: %d", img.Quality)

// crop square
img.FitCrop()

// encode to jpeg image same original image quality
err = img.EncodeJPEG(ioutil.Discard)
if err != nil {
    log.Fatal(err)
}
```
