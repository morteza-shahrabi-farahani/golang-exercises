package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/morteza-shahrabi-farahani/golang-exercises/mastering-go/Phone-book/internal/phonebook"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from ", r.Host)
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "This page does not exist :(")
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	inputParameter := r.PathValue("id")

	id, err := strconv.Atoi(inputParameter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
	}

	appErr := phonebook.Delete(int64(id))
	if err != nil {
		w.WriteHeader(int(appErr.StatusCode))
		fmt.Fprint(w, appErr.Message)
	}

	w.WriteHeader(http.StatusOK)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	entries, appErr := phonebook.GetList()
	if appErr != nil {
		w.WriteHeader(int(appErr.StatusCode))
		fmt.Fprint(w, appErr.Message)
	}

	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.MarshalIndent(phonebook.ListResponse{Entries: entries}, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResponse))
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
	}

	var entry phonebook.Entry

	err = json.Unmarshal(body, &entry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
	}

	id, appErr := phonebook.Insert(&entry)
	if appErr != nil {
		w.WriteHeader(int(appErr.StatusCode))
		fmt.Fprint(w, appErr.Message)
	}

	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.MarshalIndent(phonebook.InsertResponse{ID: id}, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResponse))
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	data, appErr := phonebook.GetList()
	if appErr != nil {
		w.WriteHeader(int(appErr.StatusCode))
		fmt.Fprint(w, appErr.Message)
	}

	telephone := r.URL.Query().Get("phone-number")
	entry, appErr := phonebook.Serach(data, telephone)
	if appErr != nil {
		w.WriteHeader(int(appErr.StatusCode))
		fmt.Fprint(w, appErr.Message)
	}

	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.MarshalIndent(entry, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResponse))
}

func StartHander() {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:         ":8001",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	mux.Handle("/list", http.HandlerFunc(listHandler))
	mux.Handle("/insert", http.HandlerFunc(insertHandler))
	mux.Handle("/delete/{id}", http.HandlerFunc(deleteHandler))
	mux.Handle("/search/", http.HandlerFunc(searchHandler))
	mux.Handle("/metrics", promhttp.Handler())

	fmt.Println("Ready to serve at", "8001")

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
