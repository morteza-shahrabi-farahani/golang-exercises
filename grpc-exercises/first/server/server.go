package main

import (
	"context"
	"log"
	"net"
	"time"

	calculator "github.com/morteza-shahrabi-farahani/golang-exercises/grpc-exercises/first/proto/api"
	"google.golang.org/grpc"
)

type server struct{
	calculator.UnsafeCalculatorServiceServer
	// calculator.UnimplementedCalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *calculator.SumRequest) (*calculator.SumResponse, error) {
	firstVar := req.FirstVar
	secondVar := req.SecondVar
	result := &calculator.SumResponse{
		Result: firstVar + secondVar,
	}

	return result, nil
}

func (*server) PrimeNumberDecomposition(req *calculator.PrimeNumberDecompositionRequest, stream calculator.CalculatorService_PrimeNumberDecompositionServer) error {
	input := int(req.GetInput())
	for i := 2; i < input; i++ {
		if input % i == 0 {
			result := &calculator.PrimeNumberDecompositionResponse{
				Result: int32(i),
			}

			stream.Send(result)
			input /= i
			i--
			time.Sleep(1 * time.Second)
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculator.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Println("Failed to serve: %v", err)
	}
}