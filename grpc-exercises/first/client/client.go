package main

import (
	"context"
	"fmt"
	"io"

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