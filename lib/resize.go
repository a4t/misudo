package misudo

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Response struct {
	Status   int
	Filepath string
}

func Resize(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()
	imgSource := ImgSource{}
	imgSource.getSource(r)

	decoder := Decoder{
		Raw:    imgSource.Raw,
		Format: imgSource.Format,
	}
	img := decoder.decode()

	resizer := Resizer{Source: img}
	resizer.getResizeImageSize(r)
	resizer.make()

	if r.Form["output"][0] == "file" {
		fileOutput(w, imgSource, resizer)
	} else {
		responseOutput(w, imgSource, resizer)
	}
}

func fileOutput(w http.ResponseWriter, imgSource ImgSource, resizer Resizer) {
	fileOutput := FileOutput{
		Format: imgSource.Format,
		Source: resizer.Dist,
	}
	fileOutput.create()

	response := Response{Status: 200, Filepath: fileOutput.Name}
	res, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func responseOutput(w http.ResponseWriter, imgSource ImgSource, resizer Resizer) {
	distSource := DistSource{
		Decode: resizer.Dist,
		Format: imgSource.Format,
	}
	distSource.render(w)
}
