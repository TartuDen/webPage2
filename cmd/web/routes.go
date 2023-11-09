package main

import (
	"net/http"

	"github.com/TartuDen/webPage2/pkg/config"
	"github.com/TartuDen/webPage2/pkg/handler"
	"github.com/bmizerany/pat"
)

func rounts(a *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handler.Repo.MainHandler))
	mux.Get("/about", http.HandlerFunc(handler.Repo.AboutHandler))
	return mux
}
