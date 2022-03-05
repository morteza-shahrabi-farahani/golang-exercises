package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Get(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, description, year, release_date, rating, runtime, mpaa_rating, created_at,
		 updated_at from movies where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)
	// fmt.Println(row)

	var movie Movie
	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Rating,
		&movie.Runtime,
		&movie.MPAARating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	query = `select movies_genres.id, movies_genres.movie_id, movies_genres.genre_id, genres.genre_name
	from movies_genres Left join genres on (movies_genres.genre_id = genres.id) 
	where movies_genres.movie_id = $1`

	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	var genres = make(map[int]string)
	for rows.Next() {
		var movieGenre MovieGenre
		err := rows.Scan(
			&movieGenre.ID,
			&movieGenre.MovieID,
			&movieGenre.GenreID,
			&movieGenre.Genre.GenreName,
		)
		if err != nil {
			return nil, err
		}
		genres[movieGenre.ID] = movieGenre.Genre.GenreName
	}
	movie.MovieGenres = genres

	return &movie, nil

}

func (db *DBModel) All() ([]Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select * from movies`
	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie

	for rows.Next() {
		var movie Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.Rating,
			&movie.MPAARating,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		genreQuery := `select movies_genres.id, movies_genres.movie_id, movies_genres.genre_id, genres.genre_name
		from movies_genres Left join genres on (movies_genres.genre_id = genres.id) 
		where movies_genres.movie_id = $1`

		genreRows, _ := db.DB.QueryContext(ctx, genreQuery, movie.ID)
		defer genreRows.Close()

		var genres = make(map[int]string)
		for genreRows.Next() {
			var movieGenre MovieGenre
			err := genreRows.Scan(
				&movieGenre.ID,
				&movieGenre.MovieID,
				&movieGenre.GenreID,
				&movieGenre.Genre.GenreName,
			)
			if err != nil {
				return nil, err
			}
			genres[movieGenre.ID] = movieGenre.Genre.GenreName
		}
		movie.MovieGenres = genres
		movies = append(movies, movie)
	}

	return movies, nil
}

func (db *DBModel) AllGenres() ([]Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select * from genres`
	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	// defer rows.Close()
	var genres []Genre
	for rows.Next() {
		var genre Genre
		err := rows.Scan(
			&genre.ID,
			&genre.GenreName,
			&genre.CreatedAt,
			&genre.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		genres = append(genres, genre)
	}

	return genres, nil
}

func (db *DBModel) GetAllMoviesByGenreName(genre_name string) ([]Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	fmt.Println("name is:", genre_name)
	genre_name = strings.ReplaceAll(genre_name, " ", "")
	fmt.Println("name is:", genre_name)
	getGenreIdquery := `select * from genres where genre_name = $1`
	row := db.DB.QueryRowContext(ctx, getGenreIdquery, genre_name)
	fmt.Println(row)

	var genre Genre
	err := row.Scan(
		&genre.ID,
		&genre.GenreName,
		&genre.CreatedAt,
		&genre.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	fmt.Println(genre.ID)

	query := fmt.Sprintf("select * from movies where id in (select movie_id from movies_genres where genre_id = %d)", genre.ID)
	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie

	for rows.Next() {
		var movie Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.Rating,
			&movie.MPAARating,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		// fmt.Println(movie)

		genreQuery := `select movies_genres.id, movies_genres.movie_id, movies_genres.genre_id, genres.genre_name
		from movies_genres Left join genres on (movies_genres.genre_id = genres.id) 
		where movies_genres.movie_id = $1`

		genreRows, _ := db.DB.QueryContext(ctx, genreQuery, movie.ID)
		defer genreRows.Close()

		var genres = make(map[int]string)
		for genreRows.Next() {
			var movieGenre MovieGenre
			err := genreRows.Scan(
				&movieGenre.ID,
				&movieGenre.MovieID,
				&movieGenre.GenreID,
				&movieGenre.Genre.GenreName,
			)
			if err != nil {
				return nil, err
			}
			genres[movieGenre.ID] = movieGenre.Genre.GenreName
		}
		movie.MovieGenres = genres
		movies = append(movies, movie)
	}

	return movies, nil
}
