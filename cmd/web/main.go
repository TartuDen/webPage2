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
		log.Fatal("cannot create template cache",err)
	}
	app.TemplateCache=tc

	http.HandleFunc("/", handler.MainHandler)
	http.HandleFunc("/about", handler.AboutHandler)
	fmt.Println("Server started on Port:", Port)
	err = http.ListenAndServe(Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
