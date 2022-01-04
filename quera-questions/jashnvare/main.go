package main

import "fmt"

func main() {
	var input int
	var result int
	result = 1
	fmt.Scan(&input)
	for i := 2; i <= input/2; i++ {
		if input%i == 0 {
			result = input / i
			break
		}
	}

	fmt.Println(result)
}
