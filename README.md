# image

`image/jpeg` base pure go image utility.

[![Go Reference][go-dev-img]][go-dev-url]
![Test][github-actions-img]
[![codecov][codecov-img]][codecov-url]
[![Maintainability][codeclimate-img]][codeclimate-url]
[![Go Report Card][goreportcard-img]][goreportcard-url]

## features
- Decode jpeg with quality value.
- Apply EXIF orientation to image pixels.

## compatibility
- Go 1.13+

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

[go-dev-img]: https://pkg.go.dev/badge/github.com/go-oss/image.svg
[go-dev-url]: https://pkg.go.dev/github.com/go-oss/image
[github-actions-img]: https://github.com/go-oss/image/workflows/Test/badge.svg?branch=master
[codecov-img]: https://codecov.io/gh/go-oss/image/branch/master/graph/badge.svg
[codecov-url]: https://codecov.io/gh/go-oss/image
[codeclimate-img]: https://api.codeclimate.com/v1/badges/1d34d57d8c26f502ce7d/maintainability
[codeclimate-url]: https://codeclimate.com/github/go-oss/image/maintainability
[goreportcard-img]: https://goreportcard.com/badge/github.com/go-oss/image
[goreportcard-url]: https://goreportcard.com/report/github.com/go-oss/image
