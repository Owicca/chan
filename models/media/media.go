package media

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Owicca/chan/models/logs"
	"github.com/abema/go-mp4"
	"github.com/ebml-go/webm"
	"golang.org/x/image/draw"
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
	m.Thumb = fmt.Sprintf("%d_thumb.%s", name, ext)
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
	f.Seek(0, io.SeekStart)
	io.Copy(osT, f)

	osF.Seek(0, io.SeekStart)
	m.X, m.Y = GetMediaXY(osF, mimeType)

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
		log.Println("getmediaxy is vid", mime)
		if mime == "video/webm" {
			m := webm.WebM{}
			_, err := webm.Parse(f, &m)
			if err != nil {
				logs.LogErr(op, errors.Errorf("err while parsing %s (%s)", mime, err))
				return x, y
			}
			vidTrack := m.FindFirstVideoTrack()
			if vidTrack == nil {
				logs.LogErr(op, errors.Str("no video track found"))
				return x, y
			}
			x, y = int(vidTrack.Video.DisplayHeight), int(vidTrack.Video.DisplayWidth)
		} else if mime == "video/mp4" {
			// TODO: run ffmpeg from cli until I research https://pkg.go.dev/github.com/abema/go-mp4
			probe, err := mp4.Probe(f)
			if err != nil {
				logs.LogErr(op, errors.Errorf("error on probing (%s)", err))
				return x, y
			}
			for _, track := range probe.Tracks {
				if track.AVC != nil {
					x, y = int(track.AVC.Width), int(track.AVC.Height)
					break
				}
			}
		}
	} else {
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

func IsVid(mime string) bool {
	_, ok := VidMimeTypes[mime]
	return ok
}

func ResizeMedia(ff *io.Reader, mime string) {
	name := "1652969265523"
	ext := "png"
	pth := fmt.Sprintf("%s/%s.%s", GetStaticDir(), name, ext)
	th := fmt.Sprintf("%s/%s_thumb.%s", GetStaticDir(), name, ext)
	f, _ := os.Open(pth)
	defer f.Close()
	//mime := "image/png"
	src, _, err := image.Decode(f)
	if err != nil {
		log.Println("err on image decoding", err)
		return
	}

	dr := image.Rect(0, 0, src.Bounds().Max.X/2, src.Bounds().Max.Y/2)
	dst, _ := os.Create(th)
	defer dst.Close()

	var res image.Image = scaleTo(src, dr, draw.NearestNeighbor)

	err = png.Encode(dst, res)
	if err != nil {
		log.Println("err on png encode", err)
		return
	}
	cmdStr := `ffmpeg -i "%s" -an -q 0 -vf scale="'if(gt(iw,ih),-1,200):if(gt(iw,ih),200,-1)', crop=200:200:exact=1" -vframes 1 "%s"`
	_ = fmt.Sprintf(cmdStr, pth, th)
}

// Rescale `src` to `rect` using `scale`,
// Rescaling is done inplace
func scaleTo(src image.Image, rect image.Rectangle, scale draw.Scaler) image.Image {
	dst := image.NewRGBA(rect)

	scale.Scale(dst, rect, src, src.Bounds(), draw.Over, nil)

	return dst
}
