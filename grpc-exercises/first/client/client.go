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

	// doUnary(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	doBiDirectionalStreaming(c)
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

func doBiDirectionalStreaming(c calculator.CalculatorServiceClient) {
	requests := []*calculator.MaximumRequest {
		&calculator.MaximumRequest{
			Input: 1,
		},
		&calculator.MaximumRequest{
			Input: 2,
		},
		&calculator.MaximumRequest{
			Input: 3,
		},
		&calculator.MaximumRequest{
			Input: 1,
		},
		&calculator.MaximumRequest{
			Input: 10,
		},
		&calculator.MaximumRequest{
			Input: 5,
		},
		&calculator.MaximumRequest{
			Input: 7,
		},
	}

	stream, err := c.Maximum(context.Background())
	if err != nil {
		log.Println("error while calling Maximum rpc: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		// function to send a bunch of messages to the server (go routine)
		for _, req := range requests {
			fmt.Println("sending messages: ")
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		// function to receive a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while receiving: %v", err)
				break
			}

			fmt.Println("received: ", res.GetResult())
		}
		
		close(waitc)
	}()

	// block until everything is done
	<-waitc

}