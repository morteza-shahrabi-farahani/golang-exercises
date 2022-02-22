package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *Application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
	if err != nil {
		fmt.Println("error is ", err)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *Application) errorJSON(w http.ResponseWriter, err error) {
	type jsonError struct {
		Message string `json:"message"`
	}

	var error = jsonError{Message: err.Error()}

	app.writeJSON(w, http.StatusBadRequest, error, "Error")
}
