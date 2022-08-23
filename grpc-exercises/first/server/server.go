package main

import (
	"fmt"
	"log"
	"net"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln("Failed to listen: %v", err)
	}

	fmt.Println(lis)
	// calculator.RegisterCalculatorServiceServer(s, &server{})
}