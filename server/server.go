package main

import (
	"context"
	"log"
	"net"

	calcpb "github.com/matheustp/calculator-grpc/pb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calc(ctx context.Context, req *calcpb.CalculatorRequest) (*calcpb.CalculatorResponse, error) {
	return &calcpb.CalculatorResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}, nil
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error on listen: %v", err)
	}
	defer l.Close()
	log.Println("Start listening")
	s := grpc.NewServer()
	calcpb.RegisterCalculatorServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("Error when starting to serve: %v", err)
	}

}
