package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ojzene/goarithmetic/calculator/calculator"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := calculator.NewCalculatorServiceClient(conn)
	expression := "10 + 5"
	req := &calculator.ExpressionRequest{Expression: expression}

	resp, err := client.EvaluateExpression(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to evaluate expression: %v", err)
	}

	fmt.Printf("Result of '%s' is %d\n", expression, resp.Result)
}
