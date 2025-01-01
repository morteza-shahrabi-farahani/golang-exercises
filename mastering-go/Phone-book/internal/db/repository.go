package db

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/morteza-shahrabi-farahani/golang-exercises/mastering-go/Phone-book/internal/model"
)

func GetList() ([]model.Entry, *model.PhoeBookError) {
	db, err := ConnectDB()
	if err != nil {
		return nil, &model.PhoeBookError{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM phone_book")
	if err != nil {
		return nil, &model.PhoeBookError{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	var entries []model.Entry
	for rows.Next() {
		var entry model.Entry

		err := rows.Scan(&entry.ID, &entry.Name, &entry.Surname, &entry.PhoneNumber)
		if err != nil {
			return nil, &model.PhoeBookError{Message: err.Error(), StatusCode: http.StatusInternalServerError}
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

func Insert(entry *model.Entry) (int64, *model.PhoeBookError) {
	db, err := ConnectDB()
	if err != nil {
		return 0, &model.PhoeBookError{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	defer db.Close()

	var id int64
	err = db.QueryRow("INSERT INTO phone_book (name, surname, phone_number) VALUES ($1, $2, $3) RETURNING id", entry.Name, entry.Surname, entry.PhoneNumber).Scan(&id)
	if err != nil {
		return 0, &model.PhoeBookError{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	return id, nil
}

func Delete(id int64) *model.PhoeBookError {
	db, err := ConnectDB()
	if err != nil {
		return &model.PhoeBookError{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	defer db.Close()

	result, err := db.Exec("DELETE FROM phone_book WHERE id = ?", id)
	if err != nil {
		return &model.PhoeBookError{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return &model.PhoeBookError{Message: err.Error(), StatusCode: http.StatusInternalServerError}
	}

	if affectedRows == 0 {
		return &model.PhoeBookError{Message: "there is no record with given phone number", StatusCode: http.StatusNotFound}
	}

	return nil
}

func Serach(data []model.Entry, telephone string) (*model.Entry, *model.PhoeBookError) {
	for _, entry := range data {
		if entry.PhoneNumber == telephone {
			return &entry, nil
		}
	}

	return nil, &model.PhoeBookError{Message: "there is no record with given phone number", StatusCode: http.StatusNotFound}
}
