package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"myaddsvc/pb"
	"myaddsvc/pkg/addendpoint"
	"myaddsvc/pkg/addservice"
	"myaddsvc/pkg/addtransport"
)

func main()  {
	var (
		service  = addservice.New()
		endpoints = addendpoint.New(service)
		grpcServer = addtransport.NewGRPCServer(endpoints)
	)
	_ = grpcServer
	fmt.Println("Server is running...")

	// Make a listener
	lis, err := net.Listen("tcp", "0.0.0.0:8001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Make a gRPC server
	grpcS := grpc.NewServer()
	pb.RegisterAddServer(grpcS, grpcServer)

	// Run the gRPC server
	if err := grpcS.Serve(lis); err != nil{
		log.Fatalf("Failed to serve: %v", err)
	}

	//************ Old method *******************
	//grpcListener, err := net.Listen("tcp", ":8000")
	//if err != nil{
	//	log.Fatalf("gRPC listen error: %v", err)
	//	os.Exit(1)
	//}
	//log.Fatalf("gRPC listening on port :8000")
	//baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	//pb.RegisterAddServer(baseServer, grpcServer)
	//baseServer.Serve(grpcListener)

}
