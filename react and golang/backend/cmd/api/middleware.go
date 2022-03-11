package main

import "net/http"

func (app *Application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")

		next.ServeHTTP(w, r)
	})
}

func (app *Application) checkToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)) {
		w.Header().Set("Vary", "Authorization")
		authHeader := w.Header().Get("Authorization")

		if authHeader == "" {
			
		}
	}
}
