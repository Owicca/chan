package media

const (
	Vid string = "vid"
	Img string = "img"
)

var (
	ImgMimeTypes map[string]string = map[string]string{
		"image/gif":  "gif",
		"image/jpeg": "jpeg",
		"image/png":  "png",
		"image/webp": "webp",
	}
	VidMimeTypes map[string]string = map[string]string{
		"video/webm": "webm",
		"video/mp4":  "mp4",
	}
	AllowedMedia map[string]string = map[string]string{}
)

func init() {
	for id, ext := range ImgMimeTypes {
		AllowedMedia[id] = ext
	}
	for id, ext := range VidMimeTypes {
		AllowedMedia[id] = ext
	}
}

func IsVid(mime string) bool {
	_, ok := VidMimeTypes[mime]
	return ok
}
