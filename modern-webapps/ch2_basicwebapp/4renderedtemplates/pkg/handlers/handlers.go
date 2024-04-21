package handlers

import (
	"golanglearning/modern-webapps/ch2_basicwebapp/3templates/pkg/renders"
	"net/http"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "home.page.tmpl.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "about.page.tmpl.html")
}
