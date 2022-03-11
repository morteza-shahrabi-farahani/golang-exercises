package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *Application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (app *Application) routes() http.Handler {
	var router = httprouter.New()
	secure := alice.New(app.checkToken)

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	router.HandlerFunc(http.MethodGet, "/movie/:id", app.getOneMovie)
	router.HandlerFunc(http.MethodGet, "/movies", app.getAllMovies)
	router.HandlerFunc(http.MethodGet, "/genres", app.getAllGenres)
	router.HandlerFunc(http.MethodGet, "/genres/:genre_name", app.getAllMoviesByGenre)
	// router.HandlerFunc(http.MethodPost, "/admin/add-movie", app.addMovie)
	// router.HandlerFunc(http.MethodPut, "/admin/movie/:id/edit", app.editMovie)
	// router.HandlerFunc(http.MethodDelete, "/admin/movie/:id/delete", app.deleteMovie)
	router.POST("/admin/add-movie", app.wrap(secure.ThenFunc(app.addMovie)))
	router.PUT("/admin/movie/:id/edit", app.wrap(secure.ThenFunc(app.editMovie)))
	router.DELETE("/admin/movie/:id/delete", app.wrap(secure.ThenFunc(app.deleteMovie)))
	router.HandlerFunc(http.MethodPost, "/login", app.login)

	return app.enableCORS(router)
}
