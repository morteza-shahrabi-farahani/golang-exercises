package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/morteza-shahrabi-farahani/golang-exercises/models"
	"github.com/pascaldekloe/jwt"
	"golang.org/x/crypto/bcrypt"
)

var validUser = models.User{
	ID:       1,
	Email:    "morteza.shahrabii@gmail.com",
	Username: "morteza",
	Password: "$2a$14$CrsAqF9j/c518Nx.h4HzSepsKaMjg4RJV.A8KhoHMD43nrXAsQoay",
}

type Credentials struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password string  `json:"password"`
}

func (app *Application) login(w http.ResponseWriter, r *http.Request) {
	var cred Credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		fmt.Println(err)
		app.errorJSON(w, err)
		return
	}

	hashPass := validUser.Password
	fmt.Println(cred.Password)
	err = bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(cred.Password))
	if err != nil {
		fmt.Println(err)
		app.errorJSON(w, errors.New("unauthorized"))
		return
	}

	var claims jwt.Claims
	claims.Subject = fmt.Sprint(validUser.ID)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = "morteza-shahrabi-farahani.com"
	claims.Audiences = []string{"morteza-shahrabi-farahani.com"}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(app.Cfg.jwtKey))
	if err != nil {
		fmt.Println(err)
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusOK, jwtBytes, "response")

}

