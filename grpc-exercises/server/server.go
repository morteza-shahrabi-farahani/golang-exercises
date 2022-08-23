package main

import (
	"log"
	"net"

	"github.com/morteza-shahrabi-farahani/golang-exercises/grpc-excercises/proto"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	
	calculator.RegisterCalculatorServiceServer(s, &server{})
	
}