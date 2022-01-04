package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func readingFile(textChannel chan string) {
	inputFile, inputError := os.Open("input.txt")
	//var lock sync.Mutex
	//var arithmetic sync.WaitGroup
	if inputError != nil {
		fmt.Println(inputError)
	} 
	defer inputFile.Close()

    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        textChannel <- scanner.Text()
    }
    close(textChannel)
	// fmt.Println("text is ", readScanner.Text())
	// textChannel <- readScanner.Text()
	// lock.Lock()
	// defer lock.Unlock()
	// time.Sleep(1 * time.Second)
	// textChannel <- readScanner.Text()
	// fmt.Printf("hello")
	// fmt.Println(len(textChannel))
	// fmt.Println(<- textChannel)
	// time.Sleep(10 * time.Second)
	
}

func writeToFile(textChannel chan string) {
	outputFile, outputCreationError := os.Create("output.txt")
     
    if outputCreationError != nil {
        fmt.Println("failed creating file: %s", outputCreationError)
    }
	// lock.Lock()
	// defer lock.Unlock()
	for name := range textChannel {
		outputFile.WriteString(name)
	}
	// outputFile.WriteString(<-textChannel)
	
	// outputFile.WriteString(<-textChannel)
}

func main() {
	//inputFile, inputError := os.Open("input.txt")
	//var lock sync.Mutex
	var arithmetic sync.WaitGroup
	// if inputError != nil {
	// 	fmt.Println(inputError)
	// } 

	outputFile, outputCreationError := os.Create("output.txt")
     
    if outputCreationError != nil {
        fmt.Println("failed creating file: %s", outputCreationError)
    }

	//defer inputFile.Close()
	//defer outputFile.Close()
	textChannel := make(chan string)
	//readScanner := bufio.NewScanner(inputFile)

	// go readingFile(textChannel, readScanner, lock)
	// go writeToFile(textChannel, *outputFile, lock)
	// var stdoutBuff bytes.Buffer
	// defer stdoutBuff.WriteTo(os.Stdout)
	// intStream := make(chan int, 4)
	// go func() {
	// defer close(intStream)
	// defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
	// for i := 0; i < 5; i++ {
	// fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
	// intStream <- i
	// }
	// }()
	// for integer := range intStream {
	// fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	// }


	// for readScanner.Scan() {
		// arithmetic.Add(1)
		// go func() {
		// 	readingFile(textChannel, readScanner, lock)
		// 	arithmetic.Done()
		// }()

		// go func() {
		// 	close(textChannel)
		// }()

		// arithmetic.Add(1)
		// go func() {
		// 	writeToFile(textChannel, *outputFile, lock)
		// 	arithmetic.Done()
		// }()
		// textChannel <- readScanner.Text()
		// outputFile.WriteString(<-textChannel)
		
		//fmt.Println(readScanner.Text())
		go readingFile(textChannel)

		for name := range textChannel {
			outputFile.WriteString(name)
		}

		//go writeToFile(textChannel)

		// go func() {
		// 	readingFile(textChannel, readScanner, lock)
		// 		lock.Lock()
		// 		defer lock.Unlock()
		// 		// time.Sleep(1 * time.Second)
		// 		textChannel <- readScanner.Text()
		// }()
		
	// }

	arithmetic.Wait()
	//close(textChannel)

	// time.Sleep(2 * time.Second)
	
	// fmt.Println(len(textChannel))
	// for i := 0; i < len(textChannel); i++ {
	// 	fmt.Println(<- textChannel)
	// }
}


