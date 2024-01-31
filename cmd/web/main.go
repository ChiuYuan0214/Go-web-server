package main

import (
	"hello-world-app/helpers"
	"hello-world-app/pkg/config"
	"hello-world-app/pkg/handlers"
	"hello-world-app/pkg/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNum = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // determine if using SSL/LTS
	app.Session = session

	tc, err := render.CreateTemplateCacheV2()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHanders(repo)
	render.NewTemplates(&app)

	// http.HandleFunc("/", repo.Home)
	// http.HandleFunc("/about", repo.About)
	helpers.LogPort(portNum)

	// _ = http.ListenAndServe(portNum, nil)

	srv := &http.Server{
		Addr:    portNum,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
