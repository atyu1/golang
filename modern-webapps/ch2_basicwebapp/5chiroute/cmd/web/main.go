package main

import (
	"golanglearning/modern-webapps/ch2_basicwebapp/5chiroute/pkg/config"
	"golanglearning/modern-webapps/ch2_basicwebapp/5chiroute/pkg/handlers"
	"golanglearning/modern-webapps/ch2_basicwebapp/5chiroute/pkg/renders"
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

	r := routes(&app)
	_ = http.ListenAndServe(portNumber, r)
}
