package main

import (
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"myaddsvc/pb"
	"myaddsvc/pkg/addendpoint"
	"myaddsvc/pkg/addservice"
	"myaddsvc/pkg/addtransport"
	"net"
	"os"
)

func main()  {
	var (
		service  = addservice.New()
		endpoints = addendpoint.New(service)
		grpcServer = addtransport.NewGRPCServer(endpoints)
	)

	grpcListener, err := net.Listen("tcp", ":8000")
	if err != nil{
		log.Fatalf("gRPC listen error: %v", err)
		os.Exit(1)
	}
	log.Fatalf("gRPC listening on port :8000")
	baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	pb.RegisterAddServer(baseServer, grpcServer)
	baseServer.Serve(grpcListener)

}
