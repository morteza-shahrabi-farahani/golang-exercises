package main

import (
	"encoding/json"
	"net/http"
)

func (app *Application) statusHandler(w http.ResponseWriter, r *http.Request) {
	var status AppStatus
	status.Environment = app.Cfg.env
	status.Version = version
	status.Status = "OK"

	statusJson, err := json.Marshal(status)
	if err != nil {
		app.Logger.Printf("error is: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(statusJson)
}