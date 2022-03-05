package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		fmt.Println("error is ", err)
		app.errorJSON(w, err)
		return
	}

	fmt.Println("id is: ", id)

	movie, err := app.Models.DB.Get(id)

	// var movie models.Movie
	// movie.ID = id
	// movie.Title = "Forrest Gump"
	// movie.Description = "One of the best movies in the world and in history."
	// movie.CreatedAt = time.Now()
	// movie.UpdatedAt = time.Now()
	// movie.MPAARating = "PG-13"
	// movie.Rating = 10
	// movie.Runtime = 120
	// movie.Year = 1993
	// movie.ReleaseDate = time.Now()

	err = app.writeJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *Application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.Models.DB.All()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}

func (app *Application) getAllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := app.Models.DB.AllGenres()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, genres, "genres")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *Application) getAllMoviesByGenre(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	var genre_name string
	genre_name = params.ByName("genre_name")

	fmt.Println(genre_name)

	movies, err := app.Models.DB.GetAllMoviesByGenreName(genre_name)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}
