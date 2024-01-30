package main

import (
	"hello-world-app/helpers"
	"hello-world-app/pkg/config"
	"hello-world-app/pkg/handlers"
	"hello-world-app/pkg/render"
	"log"
	"net/http"
)

const portNum = ":8080"

func main() {
	var app config.AppConfig
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
