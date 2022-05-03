package main

import (
	"fmt"
	"log"
	"net/http"

	"gihub.com/olamilekan000/gowebtmpl/pkg/config"
	"gihub.com/olamilekan000/gowebtmpl/pkg/handlers"
	"gihub.com/olamilekan000/gowebtmpl/pkg/render"
)

const portNumber = ":9995"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc

	// app := config.AppConfig{
	// 	TemplateCache: tc,
	// }

	fmt.Println(app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
