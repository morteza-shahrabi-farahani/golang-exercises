package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/morteza-shahrabi-farahani/golang-exercises/models"
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

	var movie models.Movie
	movie.ID = id
	movie.Title = "Forrest Gump"
	movie.Description = "One of the best movies in the world and in history."
	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()
	movie.MPAARating = "PG-13"
	movie.Rating = 10
	movie.Runtime = 120
	movie.Year = 1993
	movie.ReleaseDate = time.Now()

	app.writeJSON(w, http.StatusOK, movie, "movie")
}

func (app *Application) getAllMovies(w http.ResponseWriter, r *http.Request) {

}
