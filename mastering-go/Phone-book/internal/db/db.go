package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	connectionStr := "user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to database!!")
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("cannot connect to database: %v", err)
	}

	return conn, nil
}
