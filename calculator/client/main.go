package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cmwylie19/go-grpc-course/calculator/calcpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()
	client := calcpb.NewCalcServiceClient(cc)

	// fmt.Printf("Created client: %f", client)
	doUnary(client)

}

func doUnary(c calcpb.CalcServiceClient) {
	fmt.Println("doUnary RPC...")
	req := &calcpb.CalcRequest{
		Calc: &calcpb.Calc{
			NumOne: 1,
			NumTwo: 5,
		},
	}
	res, err := c.Calc(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Calc: %v\n", res.Result)
}
