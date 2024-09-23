package main

import (
	"github.com/morteza-shahrabi-farahani/golang-exercises/mastering-go/Phone-book/internal/controller"
	"github.com/morteza-shahrabi-farahani/golang-exercises/mastering-go/Phone-book/internal/phonebook"
)

type PhoneBook []phonebook.Entry

const CSVFILE = "../data/data.csv"

func main() {
	// arguments := os.Args

	// controller.CommandLineHandler(arguments)
	controller.StartHander()
}
