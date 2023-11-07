package handler

import (
	"net/http"

	"github.com/TartuDen/webPage2/pkg/renderer"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	renderer.RendererTemplate(w, "home.page.html")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	renderer.RendererTemplate(w, "about.page.html")
}
