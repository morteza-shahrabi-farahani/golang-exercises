package main

import (
	"example.com/m/v2/pkg/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/uploadFile", handlers.UploadFile).Methods("POST")
	router.HandleFunc("/downloadFile", handlers.DownloadFile).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}

}

