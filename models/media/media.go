package media

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"gorm.io/gorm"
	"upspin.io/errors"
)

type Media struct {
	Object_id   int
	Object_type string
	Deleted_at  int
	Name        string
	Type        string
	Path        string
	Thumb       string
	X           int
	Y           int
	Size        int64
}

func TotalMediaSize(db *gorm.DB) int64 {
	var count int64

	db.Raw(`
	SELECT SUM(m.size) FROM media m
	WHERE
	m.deleted_at = 0
	`).Scan(&count)

	return count
}

func CreateMedia(m *Media, f io.ReadSeeker) (*Media, error) {
	ext := "png"

	buf := make([]byte, 512)
	f.Read(buf)
	f.Seek(0, io.SeekStart)
	mimeType := http.DetectContentType(buf)
	ext, allowed := AllowedMedia()[mimeType]
	if !allowed {
		return m, errors.Errorf("Mime-type %s is not allowed (post_id: %d)", mimeType, m.Object_id)
	}
	if ext == "" {
		return m, errors.Errorf("Invalid extension (post_id: %d)", m.Object_id)
	}

	name := time.Now().UnixMilli()
	staticDir := GetStaticDir()

	m.Path = fmt.Sprintf("%d.%s", name, ext)
	m.Thumb = fmt.Sprintf("%ds.%s", name, ext)
	m.Type = Img
	if IsVid(mimeType) {
		m.Type = Vid
	}

	osPath := fmt.Sprintf("%s/%s", staticDir, m.Path)
	osF, err := os.Create(osPath)
	if err != nil {
		return m, err
	}
	defer osF.Close()

	osThumb := fmt.Sprintf("%s/%s", staticDir, m.Thumb)
	osT, err := os.Create(osThumb)
	if err != nil {
		return m, err
	}
	defer osT.Close()

	io.Copy(osF, f)
	io.Copy(osT, f)

	m.X, m.Y = GetMediaXY(f)

	return m, nil
}

func GetMediaXY(f io.Reader) (int, int) {
	return 100, 100
}

func GetStaticDir() string {
	wd, _ := os.Getwd()
	return fmt.Sprintf("%s/static/media", wd)
}

func AllowedMedia() map[string]string {
	return map[string]string{
		"image/gif":  "gif",
		"image/jpeg": "jpeg",
		"image/png":  "png",
		"image/webp": "webp",
		"video/webm": "webm",
	}
}

func IsVid(mime string) bool {
	_, ok := VidMimeTypes[mime]
	return ok
}
