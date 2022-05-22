package media

import (
	"io"

	"github.com/Owicca/chan/models/logs"
	"github.com/abema/go-mp4"
	"github.com/ebml-go/webm"
	"upspin.io/errors"
)

func GetMediaXYMp4(f io.ReadSeeker, mime string) (int, int) {
	const op errors.Op = "models.media.GetMediaXYMp4"
	x, y := -1, -1

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

	return x, y
}

func GetMediaXYWebm(f io.ReadSeeker, mime string) (int, int) {
	const op errors.Op = "models.media.GetMediaXYWebm"
	x, y := -1, -1

	m := webm.WebM{}
	r, err := webm.Parse(f, &m)
	if err != nil {
		logs.LogErr(op, errors.Errorf("err while parsing %s (%s)", mime, err))
		return x, y
	}
	defer r.Shutdown()
	vidTrack := m.FindFirstVideoTrack()
	if vidTrack == nil {
		logs.LogErr(op, errors.Str("no video track found"))
		return x, y
	}

	return int(vidTrack.Video.DisplayHeight), int(vidTrack.Video.DisplayWidth)
}
