package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculator.RegisterSumServiceServer(s, &server{})

	
}