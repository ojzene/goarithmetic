package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/ojzene/goarithmetic/calculatorpb"

	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := calculatorpb.NewCalculatorServiceClient(conn)
	expression := "10 + 5"
	req := &calculatorpb.ExpressionRequest{Expression: expression}

	resp, err := client.EvaluateExpression(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to evaluate expression: %v", err)
	}

	fmt.Printf("Result of '%s' is %d\n", expression, resp.Result)
}
