package main

import (
	"context"
	"fmt"

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