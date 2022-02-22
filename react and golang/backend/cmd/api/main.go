package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const version = "1.0.0"

type Config struct {
	port string
	env string
}

type AppStatus struct {
	Environment string `json:"environment"`
	Version string `json:"version"`
	Status string `json:"status"`
}

func main() {
	var config Config

	config.port = ":8080"
	config.env = "backend development with golang"

	http.HandleFunc("/status", statusHandler)
	http.ListenAndServe(config.port, nil)
	
	fmt.Println("Running")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {

	var status AppStatus
	status.Environment = "backend development with golang"
	status.Version = version
	status.Status = "OK"

	statusJson, err := json.Marshal(status)
	if err != nil {
		fmt.Fprintf(w, "error is: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(statusJson)
}