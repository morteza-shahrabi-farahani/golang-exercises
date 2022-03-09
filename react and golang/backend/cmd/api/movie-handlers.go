package main

import (
	"encoding/json"
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

func (app *Application) addMovie(w http.ResponseWriter, r *http.Request) {
	var movieReceiver models.MovieReceiver
	fmt.Println(r.Method)
	fmt.Println(r.Body)
	// log.Println("Hello")
	err := json.NewDecoder(r.Body).Decode(&movieReceiver)
	if err != nil {
		fmt.Println(err)
		app.errorJSON(w, err)
		return
	}
	fmt.Println(movieReceiver.Title)

	var newMovie models.Movie
	newMovie.Title = movieReceiver.Title
	newMovie.Description = movieReceiver.Description
	newMovie.ReleaseDate, _ = time.Parse("2006-01-02", movieReceiver.ReleaseDate)
	newMovie.MPAARating = movieReceiver.MPAARating
	newMovie.Rating, _ = strconv.Atoi(movieReceiver.Rating)
	newMovie.Runtime, _ = strconv.Atoi(movieReceiver.Runtime)
	newMovie.Year = newMovie.ReleaseDate.Year()
	newMovie.CreatedAt = time.Now()
	newMovie.UpdatedAt = time.Now()

	err = app.Models.DB.AddMovieToDB(newMovie)
	if err != nil {
		fmt.Println(err)
		app.errorJSON(w, err)
		return
	}
	// fmt.Println(movieReceiver.Title)

	type jsonResp struct {
		OK bool `json:"ok"`
	}

	ok := jsonResp{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}


func (app *Application) editMovie(w http.ResponseWriter, r *http.Request) {
	var movieReceiver models.MovieReceiver
	fmt.Println(r.Method)
	fmt.Println(r.Body)
	// log.Println("Hello")
	err := json.NewDecoder(r.Body).Decode(&movieReceiver)
	if err != nil {
		fmt.Println(err)
		app.errorJSON(w, err)
		return
	}
	fmt.Println(movieReceiver.Title)

	var newMovie models.Movie
	newMovie.ID, _ = strconv.Atoi(movieReceiver.ID)
	newMovie.Title = movieReceiver.Title
	newMovie.Description = movieReceiver.Description
	newMovie.ReleaseDate, _ = time.Parse("2006-01-02", movieReceiver.ReleaseDate)
	newMovie.MPAARating = movieReceiver.MPAARating
	newMovie.Rating, _ = strconv.Atoi(movieReceiver.Rating)
	newMovie.Runtime, _ = strconv.Atoi(movieReceiver.Runtime)
	newMovie.Year = newMovie.ReleaseDate.Year()
	newMovie.CreatedAt = newMovie.CreatedAt
	newMovie.UpdatedAt = time.Now()

	err = app.Models.DB.AddMovieToDB(newMovie)
	if err != nil {
		fmt.Println(err)
		app.errorJSON(w, err)
		return
	}
	// fmt.Println(movieReceiver.Title)

	type jsonResp struct {
		OK bool `json:"ok"`
	}

	ok := jsonResp{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

