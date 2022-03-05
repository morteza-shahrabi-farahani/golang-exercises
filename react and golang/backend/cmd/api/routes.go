package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() http.Handler {
	var router = httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	router.HandlerFunc(http.MethodGet, "/movie/:id", app.getOneMovie)
	router.HandlerFunc(http.MethodGet, "/movies", app.getAllMovies)

	return app.enableCORS(router)
}