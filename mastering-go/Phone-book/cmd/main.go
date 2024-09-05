package main

import (
	"os"

	"github.com/morteza-shahrabi-farahani/golang-exercises/mastering-go/Phone-book/internal/controller"
	"github.com/morteza-shahrabi-farahani/golang-exercises/mastering-go/Phone-book/internal/db"
	"github.com/morteza-shahrabi-farahani/golang-exercises/mastering-go/Phone-book/internal/phonebook"
)

type PhoneBook []phonebook.Entry

const CSVFILE = "../data/data.csv"

func main() {
	arguments := os.Args

	db, err := db.ConnectDB()
	if err != nil {
		return
	}

	defer db.Close()

	controller.CommandLineHandler(arguments, db)
}
