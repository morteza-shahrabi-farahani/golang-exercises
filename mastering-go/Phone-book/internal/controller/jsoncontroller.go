package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/pprof"
	"strconv"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/morteza-shahrabi-farahani/golang-exercises/mastering-go/Phone-book/internal/phonebook"
	"github.com/morteza-shahrabi-farahani/golang-exercises/mastering-go/Phone-book/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// @title Phonebook API
// @version 1.0
// @description This is a sample phonebook API server.
// @contact.name API Support
// @contact.email support@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8001
// @BasePath /

// defaultHandler
// @Summary      Handle invalid routes
// @Description  Respond with a 404 error for unknown routes
// @Tags         misc
// @Produce      plain
// @Failure      404  {string}  string  "This page does not exist"
// @Router       / [get]
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from ", r.Host)
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "This page does not exist :(")
}

// deleteHandler
// @Summary      Delete a phonebook entry
// @Description  Delete an entry by its ID
// @Tags         phonebook
// @Param        id   path      int   true  "Entry ID"
// @Produce      plain
// @Success      200  {string}  string  "Deleted successfully"
// @Failure      500  {string}  string  "Internal Server Error"
// @Router       /delete/{id} [delete]
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

// listHandler
// @Summary      List phonebook entries
// @Description  Get all phonebook entries
// @Tags         phonebook
// @Produce      json
// @Success      200  {object}  phonebook.ListResponse
// @Failure      500  {string}  string  "Internal Server Error"
// @Router       /list [get]
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

// insertHandler
// @Summary      Insert a new phonebook entry
// @Description  Add a new entry to the phonebook
// @Tags         phonebook
// @Accept       json
// @Produce      json
// @Param        entry  body      phonebook.Entry  true  "Phonebook Entry"
// @Success      200    {object}  phonebook.InsertResponse
// @Failure      500    {string}  string  "Internal Server Error"
// @Router       /insert [post]
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

// searchHandler
// @Summary      Search phonebook entries
// @Description  Search for an entry by phone number
// @Tags         phonebook
// @Param        phone-number  query     string  true  "Phone number to search"
// @Produce      json
// @Success      200  {object}  phonebook.Entry
// @Failure      500  {string}  string  "Internal Server Error"
// @Router       /search [get]
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

	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("Ready to serve at", "8001")

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		http.ListenAndServe(metrics.METRICS_PORT, nil)
	}()
}
