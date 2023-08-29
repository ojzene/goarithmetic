package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/ojzene/goarithmetic/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultCalculation = "10 + 5"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultCalculation, "Arithmetic to calculate")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := calculatorpb.NewCalculatorServiceClient(conn)

	expression := defaultCalculation

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &calculatorpb.ExpressionRequest{Expression: expression}
	resp, err := client.EvaluateExpression(ctx, req)
	if err != nil {
		log.Fatalf("Failed to evaluate expression: %v", err)
	}

	fmt.Printf("Result of '%s' is %d\n", expression, resp.Result)
}
