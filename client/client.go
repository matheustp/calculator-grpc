package main

import (
	"context"
	"log"

	calcpb "github.com/matheustp/calculator-grpc/pb"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error when dialing: %v", err)
	}
	defer cc.Close()
	log.Println("Success dialing server")
	c := calcpb.NewCalculatorClient(cc)
	numToSum := &calcpb.CalculatorRequest{
		Num1: 40,
		Num2: 2,
	}
	r, err := c.Calc(context.Background(), numToSum)
	if err != nil {
		log.Panicln("Error when calling function: %v", err)
	}
	log.Printf("Result: %v", r.GetResult())
}
