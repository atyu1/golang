package main

import (
	"golanglearning/modern-webapps/ch2_basicwebapp/4renderedtemplates/pkg/config"
	"golanglearning/modern-webapps/ch2_basicwebapp/4renderedtemplates/pkg/handlers"
	"golanglearning/modern-webapps/ch2_basicwebapp/4renderedtemplates/pkg/renders"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := renders.CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	renders.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	_ = http.ListenAndServe(portNumber, nil)
}
