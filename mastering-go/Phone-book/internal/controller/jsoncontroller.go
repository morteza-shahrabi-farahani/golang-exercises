package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/morteza-shahrabi-farahani/golang-exercises/mastering-go/Phone-book/internal/phonebook"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from ", r.Host)
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "This page does not exist :(")
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	phoneNumber := r.PathValue("phone-number")

	err := phonebook.Delete(phoneNumber)
	if err != nil {
		w.WriteHeader(int(err.StatusCode))
		fmt.Fprintf(w, err.Message)
	}

	w.WriteHeader(http.StatusOK)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	entries, appErr := phonebook.GetList()
	if appErr != nil {
		w.WriteHeader(int(appErr.StatusCode))
		fmt.Fprintf(w, appErr.Message)
	}

	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.MarshalIndent(entries, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(jsonResponse))
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}

	var entry phonebook.Entry

	err = json.Unmarshal(body, &entry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}

	id, appErr := phonebook.Insert(&entry)
	if appErr != nil {
		w.WriteHeader(int(appErr.StatusCode))
		fmt.Fprintf(w, appErr.Message)
	}

	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.MarshalIndent(id, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(jsonResponse))
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	data, appErr := phonebook.GetList()
	if appErr != nil {
		w.WriteHeader(int(appErr.StatusCode))
		fmt.Fprintf(w, appErr.Message)
	}

	telephone := r.URL.Query().Get("phone-number")
	entry, appErr := phonebook.Serach(data, telephone)
	if appErr != nil {
		w.WriteHeader(int(appErr.StatusCode))
		fmt.Fprintf(w, appErr.Message)
	}

	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.MarshalIndent(entry, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(jsonResponse))
}

func StartHander() {
	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         "8000",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	mux.Handle("/list", http.HandlerFunc(listHandler))
	mux.Handle("/insert", http.HandlerFunc(insertHandler))
	mux.Handle("/delete/{id}", http.HandlerFunc(deleteHandler))
	mux.Handle("/search/", http.HandlerFunc(searchHandler))

	err := s.ListenAndServe()
	if err != nil {
		return
	}
}
