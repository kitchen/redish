package main

import (
	"log"
	"net"

	"github.com/kitchen/redish/engine"
	"github.com/kitchen/redish/grpc_server"
	pb "github.com/kitchen/redish/redish"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	engine := engine.NewEngine()
	srv := grpc_server.NewServer(engine)
	pb.RegisterRedishServer(grpcServer, srv)
	lis, err := net.Listen("tcp", "localhost:4242")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Print("is it working?")
	grpcServer.Serve(lis)
}
