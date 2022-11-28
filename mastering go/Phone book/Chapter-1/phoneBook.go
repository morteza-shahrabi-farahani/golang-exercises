package main

import (
	"errors"
	"fmt"
	"os"
)

type Entry struct {
	Name 		string
	Surname 	string
	Telephone 	string
}

var data = []Entry{}

func search(key string) (Entry, error) {
	for _, v := range data {
		if v.Surname == key {
			return v, nil
		}
	}
	return Entry{}, errors.New("Not found")
}

func list() ([]Entry, error) {
	result := []Entry{}
	for _, v := range data {
		result = append(result, v)
	}

	return result, nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please enter required arguments!!")
		return
	}

	data = append(data, Entry{"Mihalis", "Tsoukalos", "2109416471"})
	data = append(data, Entry{"Mary", "Doe", "2109416871"})
	data = append(data, Entry{"John", "Black", "2109416123"})

	switch arguments[1] {
		case "search":
			if len(arguments) != 3 {
				fmt.Println("Please provide a search term")
				return
			}

			result, err := search(arguments[2])
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(result)

		case "list":
			userList, err := list()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(userList)

		default:
			fmt.Println("not valid option")
	}
}