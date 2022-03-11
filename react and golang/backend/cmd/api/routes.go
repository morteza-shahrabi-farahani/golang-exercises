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
	router.HandlerFunc(http.MethodGet, "/genres", app.getAllGenres)
	router.HandlerFunc(http.MethodGet, "/genres/:genre_name", app.getAllMoviesByGenre)
	router.HandlerFunc(http.MethodPost, "/admin/add-movie", app.addMovie)
	router.HandlerFunc(http.MethodPut, "/admin/movie/:id/edit", app.editMovie)
	router.HandlerFunc(http.MethodDelete, "/admin/movie/:id/delete", app.deleteMovie)
	router.HandlerFunc(http.MethodPost, "/login", app.login)

	return app.enableCORS(router)
}
