package imageutil

import (
	"net/http"
)

// MIMEType represents MIME type for image file.
type MIMEType uint8

const (
	// UndefinedType represents unknown image type.
	UndefinedType MIMEType = iota
	// JPEG represents jpeg image type.
	JPEG
	// PNG represents png image type.
	PNG
	// GIF represents gif image type.
	GIF
	// WEBP represents webp image type.
	WEBP
)

func (m MIMEType) String() string {
	switch m {
	case JPEG:
		return "image/jpeg"
	case PNG:
		return "image/png"
	case GIF:
		return "image/gif"
	case WEBP:
		return "image/webp"
	default:
		return ""
	}
}

// Ext returns file extension.
func (m MIMEType) Ext() string {
	switch m {
	case JPEG:
		return ".jpg"
	case PNG:
		return ".png"
	case GIF:
		return ".gif"
	case WEBP:
		return ".webp"
	default:
		return ""
	}
}

// MIMETypeByType returns MIMEType by type string.
func MIMETypeByType(typ string) MIMEType {
	switch typ {
	case "image/jpeg":
		return JPEG
	case "image/png":
		return PNG
	case "image/gif":
		return GIF
	case "image/webp":
		return WEBP
	default:
		return MIMEType(0)
	}
}

// DetectMimeType detects MIMEType for image file.
func DetectMimeType(fileHeader []byte) MIMEType {
	contentType := http.DetectContentType(fileHeader)
	return MIMETypeByType(contentType)
}
