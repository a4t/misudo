package misudo

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

var saveDirPath = "data/"
var fileLength = 32

type FileOutput struct {
	Name   string
	Format string
	Source image.Image
}

func (f *FileOutput) setName() {
	f.Name = saveDirPath + randString(fileLength) + "." + f.Format
}

func (f *FileOutput) create() {
	f.setName()
	if f.Format == "jpeg" {
		f.createJpegImage()
	} else if f.Format == "png" {
		f.createPngImage()
	} else if f.Format == "gif" {
		f.createGifImage()
	}
}

func (f *FileOutput) createJpegImage() {
	out, err := os.Create(f.Name)
	defer out.Close()
	if err != nil {
		log.Fatal(err)
	}
	jpeg.Encode(out, f.Source, nil)
}

func (f *FileOutput) createPngImage() {
	out, err := os.Create(f.Name)
	defer out.Close()
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(out, f.Source)
}

func (f *FileOutput) createGifImage() {
	out, err := os.Create(f.Name)
	defer out.Close()
	if err != nil {
		log.Fatal(err)
	}
	gif.Encode(out, f.Source, nil)
}
