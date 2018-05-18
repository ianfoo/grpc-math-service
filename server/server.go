package server

import (
	"context"
	"log"

	mather "github.com/ianfoo/grpc-math-service"
	"github.com/ianfoo/grpc-math-service/pb"
)

// Server performs arithmetic operations for clients.
type Server struct{}

// Add handles add requests from clients.
func (Server) Add(ctx context.Context, r *pb.AddRequest) (*pb.AddResponse, error) {
	log.Println("servicing add request")
	return &pb.AddResponse{
		Sum: mather.Add(r.A, r.B),
	}, nil
}

// Subtract handles subtract requests from clients.
func (Server) Subtract(ctx context.Context, r *pb.SubtractRequest) (*pb.SubtractResponse, error) {
	log.Println("servicing subtract request")
	return &pb.SubtractResponse{
		Difference: mather.Subtract(r.A, r.B),
	}, nil
}

// Multiply handles multiply requests from clients.
func (Server) Multiply(ctx context.Context, r *pb.MultiplyRequest) (*pb.MultiplyResponse, error) {
	log.Println("servicing multiply request")
	return &pb.MultiplyResponse{
		Product: mather.Multiply(r.A, r.B),
	}, nil
}

// Divide handles divide requests from clients.
func (Server) Divide(ctx context.Context, r *pb.DivideRequest) (*pb.DivideResponse, error) {
	log.Println("servicing divide request")
	q, rem, err := mather.Divide(r.Dividend, r.Divisor)
	if err != nil {
		return nil, err
	}
	return &pb.DivideResponse{
		Quotient:  q,
		Remainder: rem,
	}, nil
}
