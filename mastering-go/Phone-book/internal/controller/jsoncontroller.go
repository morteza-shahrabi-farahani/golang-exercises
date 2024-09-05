package controller

import (
	"fmt"
	"log"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from ", r.Host)
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "This page does not exist :(")
}

// func deleteHandler(w http.ResponseWriter, r *http.Request) {
// 	urlVariables := r.URL.Query()
// }
