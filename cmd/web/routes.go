package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mojafa/bookie/pkg/config"
	"github.com/mojafa/bookie/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	// return mux

	r := chi.NewRouter()
	// r.Use(WriteToConsole)
	r.Use(NoSurf)
	r.Use(SessionLoad)
	r.Use(middleware.Recoverer)
	r.Get("/", handlers.Repo.Home)
	r.Get("/about", handlers.Repo.About)

	// enabling static files
	fileServer := http.FileServer(http.Dir("./static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return r
}
