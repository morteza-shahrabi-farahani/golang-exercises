package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

type Config struct {
	port string
	env  string
	dsn  string
}

type AppStatus struct {
	Environment string `json:"environment"`
	Version     string `json:"version"`
	Status      string `json:"status"`
}

type Application struct {
	Cfg    Config      `json:"config"`
	Logger *log.Logger `json:"logger"`
}

func main() {
	var config Config
	config.port = ":8080"
	config.env = "backend development with golang"
	config.dsn = "postgres://postgres:123456@localhost/go_movies?sslmode=disable"

	var app Application
	app.Cfg = config
	app.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	db, err := dbOpen(config)
	if err != nil {
		fmt.Println("error is ", err)
	}
	defer db.Close()

	server := &http.Server{
		Addr:         app.Cfg.port,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Println("server is trying to Run")

	err = server.ListenAndServe()
	if err != nil {
		log.Println("error listening to server")
	}

	fmt.Println("server is runnig now")
}

func dbOpen(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.dsn)
	if err != nil {
		fmt.Println("error is ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		fmt.Println("error is ", err)
	}

	return db, nil
}
