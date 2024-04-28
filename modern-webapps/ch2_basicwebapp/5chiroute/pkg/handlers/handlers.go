package handlers

import (
	"golanglearning/modern-webapps/ch2_basicwebapp/5chiroute/pkg/config"
	"golanglearning/modern-webapps/ch2_basicwebapp/5chiroute/pkg/models"
	"golanglearning/modern-webapps/ch2_basicwebapp/5chiroute/pkg/renders"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "home.page.tmpl.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringmap := map[string]string{}
	stringmap["test"] = "Test Data from Template Variable"

	renders.RenderTemplate(w, "about.page.tmpl.html", &models.TemplateData{
		StringMap: stringmap,
	})
}
