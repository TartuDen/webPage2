package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TartuDen/webPage2/pkg/handler"
)

const Port = ":8080"

func main() {
	http.HandleFunc("/", handler.MainHandler)
	http.HandleFunc("/about", handler.AboutHandler)
	fmt.Println("Server started on Port:", Port)
	err := http.ListenAndServe(Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
