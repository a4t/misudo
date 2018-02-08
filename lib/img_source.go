package misudo

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"net/http"
)

type ImgSource struct {
	Tmp    io.Reader
	Raw    []byte
	Format string
	Decode string
}

func (i *ImgSource) getExtension() {
	_, format, err := image.DecodeConfig(bytes.NewReader(i.Raw))
	if err != nil {
		fmt.Println(err)
		return
	}
	i.Format = format
}

func (i *ImgSource) setData() {
	data, _ := ioutil.ReadAll(i.Tmp)
	i.Raw = data
}

func (i *ImgSource) getSource(r *http.Request) {
	file, _, err := r.FormFile("file")
	i.Tmp = file
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	i.setData()
	i.getExtension()
}
