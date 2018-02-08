package misudo

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Ping(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Pong")
}
