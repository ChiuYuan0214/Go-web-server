package handlers

import (
	"hello-world-app/pkg/config"
	"hello-world-app/pkg/models"
	"hello-world-app/pkg/render"
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

// set the repository for handlers
func NewHanders(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// n, err := fmt.Fprintf(w, "This is the home page")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("bytes written: %s", strconv.Itoa(n))
	render.RenderTemplateV3(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := map[string]string{}
	stringMap["test"] = "Hello, again."

	// send the data to template
	render.RenderTemplateV3(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
