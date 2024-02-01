package main

import (
	"net/http"

	"github.com/livefir/fir"
)

func main() {
	storage := NewMemoryStorage()
	controller := fir.NewController("fir_app", fir.DevelopmentMode(true))
	http.Handle("/", controller.RouteFunc(index(storage)))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./templates/static/"))))
	http.ListenAndServe(":9867", nil)
}
