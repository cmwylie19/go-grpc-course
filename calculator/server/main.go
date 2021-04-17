package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/cmwylie19/go-grpc-course/calculator/calcpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calc(ctx context.Context, req *calcpb.CalcRequest) (*calcpb.CalcResponse, error) {
	fmt.Printf("Calc functiom invoked: %v", req)
	num_one := req.GetCalc().GetNumOne()
	num_two := req.GetCalc().GetNumTwo()

	result := num_one + num_two
	res := &calcpb.CalcResponse{
		Result: result,
	}
	return res, nil
}

func (*server) Primes(req *calcpb.PrimesRequest, stream calcpb.CalcService_PrimesServer) error {
	fmt.Printf("Primes function was invoked: %v", req)
	Num := req.GetPrimes().GetNum()
	k := 2
	for Num > 1 {
		if Num%k == 0 {
			fmt.Println(k)
			Num = Num / k
		} else {
			k = k + 1
		}
	}
}
func main() {
	fmt.Println("Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalcServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
