package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	connectionStr := "user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}

	// connectionStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	// conn, err := sql.Open("postgres", connectionStr)
	// if err != nil {
	// 	panic(err)
	// }

	rows, err := conn.Query("SELECT version();")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var version string
		rows.Scan(&version)
		fmt.Println(version)
	}

	rows.Close()

	conn.Close()
}
