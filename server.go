package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ojzene/goarithmetic/calculator/calculator"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) EvaluateExpression(ctx context.Context, req *calculator.ExpressionRequest) (*calculator.ExpressionResponse, error) {
	expression := req.Expression
	result := evaluateExpression(expression)
	return &calculator.ExpressionResponse{Result: result}, nil
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
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	calculator.RegisterCalculatorServiceServer(s, &server{})
	log.Println("Server started at :8888")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
