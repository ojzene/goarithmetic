package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/ojzene/goarithmetic/calculatorpb"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct{}

func (s *server) EvaluateExpression(ctx context.Context, req *calculatorpb.ExpressionRequest) (*calculatorpb.ExpressionResponse, error) {
	expression := req.Expression
	result := evaluateExpression(expression)
	return &calculatorpb.ExpressionResponse{Result: result}, nil
}

func evaluateExpression(expression string) int32 {
	var num1, num2 int32
	operator := ""
	fmt.Sscanf(expression, "%d %s %d", &num1, &operator, &num2)

	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 != 0 {
			return num1 / num2
		} else {
			log.Println("Division by zero")
		}
	}

	return 0
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	reflection.Register(s)

	log.Println("Server started at :50051")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
