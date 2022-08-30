package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	calculator "github.com/morteza-shahrabi-farahani/golang-exercises/grpc-exercises/first/proto/api"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("error in connecting to server!")
	}
	defer cc.Close()

	c := calculator.NewCalculatorServiceClient(cc)

	doUnary(c)
	doServerStreaming(c)
}

func doUnary(c calculator.CalculatorServiceClient) {
	request := &calculator.SumRequest{
		FirstVar: 10,
		SecondVar: 20,
	}

	response, err := c.Sum(context.Background(), request)
	if err != nil {
		fmt.Println("error while calling Sum API")
	}

	fmt.Println("response: ", response)
}

func doServerStreaming(c calculator.CalculatorServiceClient) {
	request := &calculator.PrimeNumberDecompositionRequest{
		Input: 120,
	}

	response, err := c.PrimeNumberDecomposition(context.Background(), request)
	if err != nil {
		fmt.Println("error while calling Prime Number Decomposition API")
	}

	for {
		msg, err := response.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error while receiving data from stream server")
		}

		fmt.Println("response: ", msg.GetResult())
	}
}

func doClientStreaming(c calculator.CalculatorServiceClient) {
	// stream, err := c.Average(context.Background())
	// if err != nil {
	// 	fmt.Println("error while calling Average API")
	// }
	// for i := 0; i < 10; i++ {
	// 	stream.Send(&calculator.AverageRequest{
	// 		Input: int32(i),
	// 	})
	// }
	// stream.CloseSend()
	// msg, err := stream.Recv()
	// if err != nil {
	// 	fmt.Println("error while receiving data from stream server")
	// }
	// fmt.Println("response: ", msg.GetResult())

	requests := []*calculator.AverageRequest{
		&calculator.AverageRequest{
			Input: 1,
		}, 
		&calculator.AverageRequest{
			Input: 2,
		}, 
		&calculator.AverageRequest{
			Input: 3,
		}, 
		&calculator.AverageRequest{
			Input: 4,
		}, 
	}

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Println("error while calling Average API: %v", err)
	}

	for _, req := range requests {
		stream.Send(req)
		time.Sleep(time.Second)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("error while receiving data from server: %v", err)
	}

	fmt.Println("response: ", resp.GetResult())
}