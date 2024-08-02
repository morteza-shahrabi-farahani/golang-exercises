package main

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

func getList(db *sql.DB) ([]Entry, error) {
	rows, err := db.Query("SELECT * FROM phone_book")
	if err != nil {
		return nil, err
	}

	var entries []Entry
	for rows.Next() {
		var entry Entry

		err := rows.Scan(&entry.ID, &entry.Name, &entry.Surname, &entry.Telephone)
		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	return entries, nil
}
