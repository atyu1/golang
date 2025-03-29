package main

import (
	"golanglearning/1cookies/pkg/handlers"
	"golanglearning/1cookies/pkg/renders"
	"golanglearning/modern-webapps/ch5_cookies/1cookies/pkg/config"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction //Set in prod true

	app.Session = session

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
