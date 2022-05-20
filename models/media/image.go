package media

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"golang.org/x/image/webp"
)

type Imager interface {
	Decode(r io.Reader, mime string) (image.Image, error)
	Encode(w io.Writer, m image.Image) error
}

func Decode(r io.Reader, mime string) (image.Image, error) {
	switch mime {
	case "image/png":
		return png.Decode(r)
	case "image/jpeg":
		return jpeg.Decode(r)
	case "image/gif":
		return gif.Decode(r)
	case "image/webp":
		return webp.Decode(r)
	}

	res, _, err := image.Decode(r)
	return res, err
}

func Encode(w io.Writer, m image.Image, mime string) error {
	switch mime {
	case "image/png":
		return png.Encode(w, m)
	case "image/jpeg":
		opt := &jpeg.Options{
			Quality: 100,
		}
		return jpeg.Encode(w, m, opt)
	case "image/gif":
		opt := &gif.Options{
			NumColors: 256,
		}
		return gif.Encode(w, m, opt)
	}

	return nil
}
