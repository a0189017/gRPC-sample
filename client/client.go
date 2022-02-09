package main

import (
	"context"
	"fmt"
	"log"

	calculatorPB "grpc/proto/calculator"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()

	client := calculatorPB.NewCalculatorServiceClient(conn)

	sendRequest(client)
}
func sendRequest(client calculatorPB.CalculatorServiceClient) {
	fmt.Println("Staring to do a Unary RPC")
	req := &calculatorPB.CalculatorRequest{
		A: 3,
		B: 10,
	}

	res, err := client.Total(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling CalculatorService: %v \n", err)
	}

	log.Printf("Response from CalculatorService: %v", res.Result)
}
