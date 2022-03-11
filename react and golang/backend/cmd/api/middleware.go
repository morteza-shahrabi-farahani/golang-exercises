package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pascaldekloe/jwt"
)

func (app *Application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		next.ServeHTTP(w, r)
	})
}

func (app *Application) checkToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Vary", "Authorization")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			app.errorJSON(w, errors.New("Invalid auth header1"))
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			app.errorJSON(w, errors.New("Invalid auth header2"))
			return
		}

		if headerParts[0] != "Bearer" {
			app.errorJSON(w, errors.New("Invalid auth header3"))
			return
		}

		token := headerParts[1]
		fmt.Println("get token is here: ", token)
		claims, err := jwt.HMACCheck([]byte(token), []byte(app.Cfg.jwtKey))
		if err != nil {
			app.errorJSON(w, errors.New("unauthorized - failed algorithm"))
			return
		}

		if !claims.Valid(time.Now()) {
			app.errorJSON(w, errors.New("unauthorized - token expired"))
			return
		}

		if !claims.AcceptAudience("morteza-shahrabi-farahani.com") {
			app.errorJSON(w, errors.New("unauthorized - failed audience"))
			return
		}

		if claims.Issuer != "morteza-shahrabi-farahani.com" {
			app.errorJSON(w, errors.New("unauthorized - failed issuer"))
			return
		}

		userId, err := strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil {
			app.errorJSON(w, errors.New("unauthorized"))
			return
		}

		log.Printf("user is, user Id %d", userId)
		next.ServeHTTP(w, r)

	})
}
