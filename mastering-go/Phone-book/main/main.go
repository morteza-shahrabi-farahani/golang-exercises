package main

import (
	"fmt"
	"os"

	phonebook "github.com/morteza-shahrabi-farahani/golang-exercises/mastering-go/Phone-book/db"
)

type PhoneBook []phonebook.Entry

const CSVFILE = "../data/data.csv"

func main() {
	arguments := os.Args
	if err := checkArgumentsLength(arguments); err != nil {
		return
	}

	db, err := phonebook.ConnectDB()
	if err != nil {
		return
	}

	defer db.Close()

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Please provide a search term")
			return
		}

		usersList, err := phonebook.GetList(db)
		if err != nil {
			fmt.Println(err)
			return
		}

		result, err := serach(usersList, arguments[2])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(result)

	case "list":
		usersList, err := phonebook.GetList(db)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(usersList)

	case "insert":
		if err := validateInsert(arguments); err != nil {
			fmt.Println(err)
			return
		}

		id, err := phonebook.Insert(db, &phonebook.Entry{Name: arguments[2], Surname: arguments[3], PhoneNumber: arguments[4]})
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("successfully inserted with id = %d \n", id)

	case "delete":
		if err := validateDelete(arguments); err != nil {
			fmt.Println(err)
			return
		}

		if err := phonebook.Delete(db, arguments[2]); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("successfully deleted")

	default:
		fmt.Println("not valid option")
	}
}

func checkArgumentsLength(arguments []string) error {
	if len(arguments) == 1 {
		return fmt.Errorf("Please enter required arguments!!")
	}

	return nil
}

// func readFile(filePath string) ([]Entry, error) {
// 	file, err := os.Open(CSVFILE)
// 	if err != nil {
// 		return nil, fmt.Errorf("File does not exist")
// 	}

// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	fileData, err := reader.ReadAll()
// 	if err != nil {
// 		return nil, fmt.Errorf("error in reading the file")
// 	}

// 	var data []Entry
// 	for _, record := range fileData {
// 		data = append(data, Entry{
// 			Name:      record[0],
// 			Surname:   record[1],
// 			Telephone: record[2],
// 		})
// 	}

// 	sort.Sort(PhoneBook(data))

// 	return data, nil
// }

// func writeToFile(filePath string, data *Entry) error {
// 	file, err := os.OpenFile(CSVFILE, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
// 	if err != nil {
// 		return fmt.Errorf("File does not exist")
// 	}

// 	defer file.Close()

// 	csvWriter := csv.NewWriter(file)
// 	temp := []string{data.Name, data.Surname, data.PhoneNumber}

// 	fmt.Println(temp)
// 	err = csvWriter.Write(temp)
// 	if err != nil {
// 		return err
// 	}
// 	csvWriter.Flush()

// 	return nil
// }

// func deleteAndWrite(filePath string, data []Entry, userPhone string) error {
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return fmt.Errorf("File does not exist")
// 	}
// 	defer file.Close()

// 	for index, entry := range data {
// 		if entry.PhoneNumber == userPhone {
// 			data = append(data[:index], data[index+1:]...)
// 			break
// 		}
// 	}

// 	csvWriter := csv.NewWriter(file)
// 	for _, entry := range data {
// 		temp := []string{entry.Name, entry.Surname, entry.PhoneNumber}
// 		if err := csvWriter.Write(temp); err != nil {
// 			return err
// 		}
// 	}
// 	csvWriter.Flush()

// 	return nil
// }

func serach(data []phonebook.Entry, telephone string) (*phonebook.Entry, error) {
	for _, entry := range data {
		if entry.PhoneNumber == telephone {
			return &entry, nil
		}
	}

	return nil, fmt.Errorf("Not found!!")
}

func validateDelete(arguments []string) error {
	if len(arguments) != 3 {
		return fmt.Errorf("not enought arguments for delete")
	}

	return nil
}

func validateInsert(arguments []string) error {
	if len(arguments) != 5 {
		return fmt.Errorf("not enought arguments for insert")
	}

	return nil
}

func (phoneBook PhoneBook) Len() int {
	return len(phoneBook)
}

func (phoneBook PhoneBook) Less(i, j int) bool {
	if phoneBook[i].Surname == phoneBook[j].Surname {
		return phoneBook[i].Name < phoneBook[j].Name
	}

	return phoneBook[i].Surname < phoneBook[j].Surname
}

func (phoneBook PhoneBook) Swap(i, j int) {
	phoneBook[i], phoneBook[j] = phoneBook[j], phoneBook[i]
}
