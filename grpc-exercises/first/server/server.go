package main

import (
	"context"
	"fmt"
	"io"
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
	for i := 2; i <= input; i++ {
		if input % i == 0 {
			result := &calculator.PrimeNumberDecompositionResponse{
				Result: int32(i),
			}

			stream.Send(result)
			fmt.Println(result)
			input /= i
			i--
			time.Sleep(1 * time.Second)
		}
	}

	return nil
}

func (*server) Average(stream calculator.CalculatorService_AverageServer) error {
	var result float32
	var sum int32
	var counter = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			result = float32(sum) / float32(counter)
			return stream.SendAndClose(&calculator.AverageResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Println("Failed to receive: %v", err)
		}

		counter++
		input := req.GetInput()
		sum += input
	}
}

func (*server) Maximum(stream calculator.CalculatorService_MaximumServer) error {
	var max int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		input := req.GetInput()
		if input > max {
			max = input
		}

		sendErr := stream.Send(&calculator.MaximumResponse{
			Result: max,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", sendErr)
			return sendErr
		}
	}
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