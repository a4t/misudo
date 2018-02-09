package misudo

import (
	"image"
	"net/http"
	"strconv"

	"github.com/nfnt/resize"
)

type Resizer struct {
	Source image.Image
	Width  uint
	Height uint
	Dist   image.Image
}

func (d *Resizer) setResizeImageSize(r *http.Request, i ImgSource) {
	d.Width = d.getResizeParam(r.Form["w"])
	d.Height = d.getResizeParam(r.Form["h"])

	if uint(i.Width) < d.Width {
		d.Width = uint(i.Width)
	}

	if uint(i.Height) < d.Height {
		d.Height = uint(i.Height)
	}
}

func (d Resizer) getResizeParam(p []string) uint {
	if len(p) != 0 {
		size, _ := strconv.ParseUint(p[0], 10, 64)
		return uint(size)
	}
	return 0
}

func (d *Resizer) make() {
	d.Dist = resize.Resize(d.Width, d.Height, d.Source, resize.Lanczos3)
}
