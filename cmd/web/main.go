package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TartuDen/webPage2/pkg/config"
	"github.com/TartuDen/webPage2/pkg/handler"
	"github.com/TartuDen/webPage2/pkg/renderer"
)

const Port = ":8080"

func main() {
	var app config.AppConfig

	tc, err := renderer.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}
	app.TemplateCache = tc
	//this var to set use cache true or false, when in Dev mode
	app.UseCache = false

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	//passing app variable to the render package
	renderer.NewTemplate(&app)

	http.HandleFunc("/", handler.Repo.MainHandler)
	http.HandleFunc("/about", handler.Repo.AboutHandler)
	fmt.Println("Server started on Port:", Port)
	err = http.ListenAndServe(Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
