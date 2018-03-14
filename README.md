# image

`image/jpeg` base pure go image utility.

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
