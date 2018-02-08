package misudo

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"log"
)

type Decoder struct {
	Raw    []byte
	Format string
	Decode image.Image
}

func (d *Decoder) decode() image.Image {
	if d.Format == "jpeg" {
		d.jpegDecoder()
	} else if d.Format == "png" {
		d.pngDecoder()
	} else if d.Format == "gif" {
		d.gifDecoder()
	} else {

	}
	return d.Decode
}

func (d *Decoder) jpegDecoder() {
	img, err := jpeg.Decode(bytes.NewReader(d.Raw))
	if err != nil {
		log.Fatal(err)
	}
	d.Decode = img
}

func (d *Decoder) gifDecoder() {
	img, err := jpeg.Decode(bytes.NewReader(d.Raw))
	if err != nil {
		log.Fatal(err)
	}
	d.Decode = img
}

func (d *Decoder) pngDecoder() {
	img, err := png.Decode(bytes.NewReader(d.Raw))
	if err != nil {
		log.Fatal(err)
	}
	d.Decode = img
}
