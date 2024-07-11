package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Entry struct {
	Name      string
	Surname   string
	Telephone string
}

const CSVFILE = "../data/data.csv"
const maxStringLength = 7
const phoneLength = 11
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	var data = []Entry{}

	arguments := os.Args
	if err := checkArgumentsLength(arguments); err != nil {
		return
	}

	// Chapter 1 and 2 codes
	data = append(data, Entry{"Mihalis", "Tsoukalos", "2109416471"})
	data = append(data, Entry{"Mary", "Doe", "2109416871"})
	data = append(data, Entry{"John", "Black", "2109416123"})
	manipulatedData := manipulateData(100, data)

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Please provide a search term")
			return
		}

		result, err := search(manipulatedData, arguments[2])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(result)

	case "list":
		usersList, err := readFile(CSVFILE)
		if err != nil {
			fmt.Println(err)
			return
		}

		// For chapter 2
		// usersList, err := list(manipulatedData)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		fmt.Println(usersList)

	default:
		fmt.Println("not valid option")
	}
}

func search(data []Entry, key string) (Entry, error) {
	for _, v := range data {
		if v.Surname == key {
			return v, nil
		}
	}
	return Entry{}, errors.New("not found")
}

// For chapter 2
// func list(data []Entry) ([]Entry, error) {
// 	result := []Entry{}
// 	result = append(result, data...)

// 	return result, nil
// }

func generateRandomString(count int32) string {
	result := ""
	for i := 0; i < int(count); i++ {
		randomInteger := rand.Intn(len(charset))
		result += string(charset[randomInteger])
	}

	return result
}

func generateRandomPhone(phoneLength int32) string {
	result := ""
	for i := 0; i < int(phoneLength); i++ {
		randomInteger := rand.Intn(10)
		result += strconv.Itoa(randomInteger)
	}

	return result
}

func manipulateData(dataCount int32, data []Entry) []Entry {
	for i := 0; i < int(dataCount); i++ {
		data = append(data, Entry{
			Name:      generateRandomString(maxStringLength),
			Surname:   generateRandomString(maxStringLength),
			Telephone: generateRandomPhone(phoneLength),
		})
	}

	return data
}

func checkArgumentsLength(arguments []string) error {
	if len(arguments) == 1 {
		return fmt.Errorf("Please enter required arguments!!")
	}

	return nil
}

func readFile(filePath string) ([]Entry, error) {
	file, err := os.Open(CSVFILE)
	if err != nil {
		return nil, fmt.Errorf("File does not exist")
	}

	defer file.Close()

	reader := csv.NewReader(file)
	fileData, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error in reading the file")
	}

	var data []Entry
	for _, record := range fileData {
		data = append(data, Entry{
			Name:      record[0],
			Surname:   record[1],
			Telephone: record[2],
		})
	}

	return data, nil
}
