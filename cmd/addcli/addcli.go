package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"myaddsvc/pb"
)

func main()  {
	fmt.Println("Client is running...")

	cc, err := grpc.Dial("localhost:8001", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("could not connect to server: %v", err)
	}
	defer cc.Close()
	c := pb.NewAddClient(cc)

	test1(c)
}

func test1(c pb.AddClient) {
	fmt.Println("Starting to do a test")

	req := &pb.SumRequest{
		A: 25,
		B: 12,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil{
		log.Fatalf("Error while calling sum RPC: %v", err)
	}
	log.Printf("Response from server: %v", res.V)
}