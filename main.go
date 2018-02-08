package main

import (
	"flag"
	"log"
	"net/http"

	misudo "github.com/a4t/misudo/lib"
	"github.com/julienschmidt/httprouter"
)

var (
	port = flag.String("p", "8080", "Web server port")
)

func main() {
	flag.Parse()
	router := routing()
	log.Fatal(http.ListenAndServe(":"+*port, router))
}

func routing() http.Handler {
	router := httprouter.New()
	router.GET("/ping", misudo.Ping)
	router.POST("/resize", misudo.Resize)

	return router
}
