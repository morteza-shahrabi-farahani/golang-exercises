package phonebook

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/morteza-shahrabi-farahani/golang-exercises/mastering-go/Phone-book/internal/db"
)

func GetList() ([]Entry, *PhoeBookError) {
	db, err := db.ConnectDB()
	if err != nil {
		return nil, &PhoeBookError{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM phone_book")
	if err != nil {
		return nil, &PhoeBookError{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	var entries []Entry
	for rows.Next() {
		var entry Entry

		err := rows.Scan(&entry.ID, &entry.Name, &entry.Surname, &entry.PhoneNumber)
		if err != nil {
			return nil, &PhoeBookError{Message: err.Error(), StatusCode: http.StatusInternalServerError}
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

func Insert(entry *Entry) (int64, *PhoeBookError) {
	db, err := db.ConnectDB()
	if err != nil {
		return 0, &PhoeBookError{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	defer db.Close()

	result, err := db.Exec("INSERT INTO phone_book (name, surname, phone_number) VALUES ($1, $2, $3)", entry.Name, entry.Surname, entry.PhoneNumber)
	if err != nil {
		return 0, &PhoeBookError{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, &PhoeBookError{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	return id, nil
}

func Delete(phoneNumber string) *PhoeBookError {
	db, err := db.ConnectDB()
	if err != nil {
		return &PhoeBookError{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	defer db.Close()

	result, err := db.Exec("DELETE FROM phone_book WHERE phone_number = ?", phoneNumber)
	if err != nil {
		return &PhoeBookError{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return &PhoeBookError{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	if affectedRows == 0 {
		return &PhoeBookError{Message: "there is no record with given phone number", StatusCode: http.StatusNotFound}
	}

	return nil
}

func Serach(data []Entry, telephone string) (*Entry, *PhoeBookError) {
	for _, entry := range data {
		if entry.PhoneNumber == telephone {
			return &entry, nil
		}
	}

	return nil, &PhoeBookError{Message: "there is no record with given phone number", StatusCode: http.StatusNotFound}
}
