package misudo

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
)

type DistSource struct {
	Format string
	Decode image.Image
}

func (d *DistSource) render(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "image/"+d.Format)
	if d.Format == "jpeg" {
		jpeg.Encode(w, d.Decode, nil)
	} else if d.Format == "gif" {
		gif.Encode(w, d.Decode, nil)
	} else if d.Format == "png" {
		png.Encode(w, d.Decode)
	} else {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Unsupported "+d.Format)
	}
}
