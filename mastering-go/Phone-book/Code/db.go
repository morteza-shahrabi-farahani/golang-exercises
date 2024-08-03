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

		err := rows.Scan(&entry.ID, &entry.Name, &entry.Surname, &entry.PhoneNumber)
		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

func insert(db *sql.DB, entry *Entry) (int64, error) {
	result, err := db.Exec("INSERT INTO phone_book (name, surname, phone_number) VALUES ($1, $2, $3)", entry.Name, entry.Surname, entry.PhoneNumber)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func delete(db *sql.DB, phoneNumber string) error {
	result, err := db.Exec("DELETE FROM phone_book WHERE phone_number = ?", phoneNumber)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return fmt.Errorf("there is no record with given phone number")
	}

	return nil
}
