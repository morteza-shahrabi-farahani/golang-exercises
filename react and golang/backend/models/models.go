package models

import (
	"database/sql"
	"time"
)

type Models struct {
	DB DBModel
}

func NewModel(db *sql.DB) Models {
	return Models{DB: DBModel{DB: db}}
}

type Movie struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Year        int            `json:"year"`
	ReleaseDate time.Time      `json:"release_date"`
	Runtime     int            `json:"runtime"`
	Rating      int            `json:"rating"`
	MPAARating  string         `json:"mpaa_rating"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	MovieGenres map[int]string `json:"movie_genres"`
}

type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type MovieGenre struct {
	ID        int       `json:"id"`
	MovieID   int       `json:"-"`
	GenreID   int       `json:"-"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type MovieReceiver struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Year        string         `json:"year"`
	ReleaseDate string         `json:"release_date"`
	Runtime     string         `json:"runtime"`
	Rating      string         `json:"rating"`
	MPAARating  string         `json:"mpaa_rating"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	MovieGenres map[int]string `json:"movie_genres"`
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
