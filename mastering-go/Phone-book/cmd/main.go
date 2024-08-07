package main

import (
	"fmt"
	"os"

	"github.com/morteza-shahrabi-farahani/golang-exercises/mastering-go/Phone-book/internal/controller"
	"github.com/morteza-shahrabi-farahani/golang-exercises/mastering-go/Phone-book/internal/db"
	"github.com/morteza-shahrabi-farahani/golang-exercises/mastering-go/Phone-book/internal/phonebook"
)

type PhoneBook []phonebook.Entry

const CSVFILE = "../data/data.csv"

func main() {
	arguments := os.Args
	if err := checkArgumentsLength(arguments); err != nil {
		return
	}

	db, err := db.ConnectDB()
	if err != nil {
		return
	}

	defer db.Close()

	controller.Handler(arguments, db)
}

func checkArgumentsLength(arguments []string) error {
	if len(arguments) == 1 {
		return fmt.Errorf("Please enter required arguments!!")
	}

	return nil
}