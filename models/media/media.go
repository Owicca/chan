package media

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/Owicca/chan/models/logs"
	"golang.org/x/image/draw"
	"gorm.io/gorm"
	"upspin.io/errors"
)

const (
	MaxFileSize int64 = 4194304
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

	mimeType := DetectMimeType(f)
	ext, allowed := AllowedMedia[mimeType]
	if !allowed {
		return m, errors.Errorf("Mime-type %s is not allowed (post_id: %d)", mimeType, m.Object_id)
	}
	if ext == "" {
		return m, errors.Errorf("Invalid extension (post_id: %d)", m.Object_id)
	}

	name := time.Now().UnixMilli()
	staticDir := GetStaticDir()

	m.Path = fmt.Sprintf("%d.%s", name, ext)
	m.Thumb = fmt.Sprintf("%d_thumb.%s", name, "png")
	m.Type = Img
	if IsVid(mimeType) {
		m.Type = Vid
	}

	osPath := fmt.Sprintf("%s/%s", staticDir, m.Path)
	osF, err := os.Create(osPath)
	if err != nil {
		return m, err
	}
	if mimeType != "video/webm" {
		// TODO: this is such a hack!
		// Webm lib does a seek after the file is closed, causing a panic.
		// So I let the runtime Close when it closes all the fd's on ps close.
		defer osF.Close()
	}

	osThumb := fmt.Sprintf("%s/%s", staticDir, m.Thumb)

	io.Copy(osF, f)
	f.Seek(0, io.SeekStart)

	osF.Seek(0, io.SeekStart)
	m.X, m.Y = GetMediaXY(osF, mimeType)

	width, height := 200, 200

	defer ResizeMedia(osF, osPath, osThumb, width, height, mimeType)

	return m, nil
}

// get mime type of io.ReadSeeker
func DetectMimeType(f io.ReadSeeker) string {
	buf := make([]byte, 512)
	f.Read(buf)
	f.Seek(0, io.SeekStart)

	return http.DetectContentType(buf)
}

// get resolution of media
func GetMediaXY(f io.ReadSeeker, mime string) (int, int) {
	const op errors.Op = "models.media.GetMediaXY"
	x, y := -1, -1

	if IsVid(mime) {
		//log.Println("getmediaxy is vid", mime)
		if mime == "video/webm" {
			x, y = GetMediaXYWebm(f, mime)
		} else if mime == "video/mp4" {
			x, y = GetMediaXYMp4(f, mime)
		}
	} else {
		//log.Println("getmediaxy is img", mime)
		res, err := Decode(f, mime)
		if err != nil {
			logs.LogErr(op, errors.Errorf("err on image decoding (%s)", mime, err))
			return x, y
		}
		bounds := res.Bounds()
		x, y = bounds.Max.X, bounds.Max.Y
	}

	return x, y
}

func GetStaticDir() string {
	wd, _ := os.Getwd()
	return fmt.Sprintf("%s/static/media", wd)
}

func ResizeMedia(f io.ReadSeeker, from string, to string, width int, height int, mime string) {
	const op errors.Op = "models.media.ResizeMedia"

	if IsVid(mime) {
		//log.Println("is vid")

		// -y: force overwrite
		// -i: input file
		// -an: only video
		// -q 0: get 0th track
		// scale: scale video to specified w and h
		// -frames:v : get N frames
		filter := `-y -i %s -an -q 0 -vf scale=w=%[2]d:h=%[3]d -frames:v 1 %s`
		params := fmt.Sprintf(filter, from, width, height, to)

		cmd := exec.Command("ffmpeg", strings.Fields(params)...)

		var errOut bytes.Buffer
		cmd.Stderr = &errOut

		_, err := cmd.Output()
		if err != nil {
			logs.LogErr(op, errors.Errorf("could not read output for vid thumbnail cmd (%s)", err))
			log.Println(errOut.String())
			return
		}
	} else {
		//log.Println("is image")
		f.Seek(0, io.SeekStart)
		dec, err := Decode(f, mime)
		if err != nil {
			logs.LogErr(op, errors.Errorf("err on image decoding (%s)", err))
			return
		}

		dr := image.Rect(0, 0, dec.Bounds().Max.X/2, dec.Bounds().Max.Y/2)
		dst, _ := os.Create(to)
		defer dst.Close()

		var res image.Image = ScaleTo(dec, dr, draw.NearestNeighbor)

		err = Encode(dst, res, mime)
		if err != nil {
			logs.LogErr(op, errors.Errorf("err on image encoding (%s)", err))
			return
		}
	}
}
