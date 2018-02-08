package misudo

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

func create(m image.Image, mode string) {
	if mode == "jpeg" {
		createJpegImage(m)
	} else if mode == "png" {
		createPngImage(m)
	} else if mode == "gif" {
		createGifImage(m)
	} else {

	}
}

func createJpegImage(m image.Image) io.Writer {
	out, err := os.Create(saveDirPath + "/test_resized.jpg")
	defer out.Close()
	if err != nil {
		log.Fatal(err)
	}
	jpeg.Encode(out, m, nil)

	return out
}

func createPngImage(m image.Image) io.Writer {
	out, err := os.Create(saveDirPath + "/test_resized.png")
	defer out.Close()
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(out, m)

	return out
}

func createGifImage(m image.Image) io.Writer {
	out, err := os.Create(saveDirPath + "/test_resized.gif")
	defer out.Close()
	if err != nil {
		log.Fatal(err)
	}
	gif.Encode(out, m, nil)

	return out
}
