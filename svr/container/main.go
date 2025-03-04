package main

import (
	"log"
	"net"

	"github.com/mzzz-zzm/blazor-scaffold/svr/gapi"
	"github.com/mzzz-zzm/blazor-scaffold/svr/pb/greet"
	"google.golang.org/grpc"
)

func main() {
	runGrpcServer()
}

func runGrpcServer() {
	server, err := gapi.NewServer()
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	grpcServer := grpc.NewServer()
	greet.RegisterGreeterServer(grpcServer, server)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("cannot listen: ", err)
	}
	log.Printf("server listening at %v", lis.Addr())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("cannot serve: ", err)
	}
}
