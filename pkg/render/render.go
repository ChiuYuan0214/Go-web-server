package render

import (
	"bytes"
	"fmt"
	"hello-world-app/pkg/config"
	"hello-world-app/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}
var app *config.AppConfig

func NewTemplates(config *config.AppConfig) {
	app = config
}

func renderTemplateV1(w http.ResponseWriter, tmpl string) {
	parsedTmpl, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

var tc = make(map[string]*template.Template)

func renderTemplateV2(w http.ResponseWriter, key string) {
	var tmpl *template.Template
	var err error

	// check if already have the template cache
	tmpl, inMap := tc[key]
	if !inMap {
		// need to create template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(key)
		if err != nil {
			log.Println(err)
		}
		tmpl = tc[key]
	} else {
		log.Println("using cached template")
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(key string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", key),
		"./templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[key] = tmpl
	return nil
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

func RenderTemplateV3(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	tc := app.TemplateCache
	if !app.UseCache {
		tc, _ = CreateTemplateCacheV2()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("cache missed")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCacheV2() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	log.Println("templates created")

	return myCache, nil
}
