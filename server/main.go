package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/kitchen/redish/engine"
	pb "github.com/kitchen/redish/redish"
	"google.golang.org/grpc"
)

type redishServer struct {
	engine engine.Engine
	pb.UnimplementedRedishServer
}

func newServer() *redishServer {
	s := &redishServer{engine: engine.NewEngine()}
	return s
}

func (s *redishServer) Get(ctx context.Context, key *pb.Key) (*pb.SingleValue, error) {
	log.Printf("GET %v", key.Key)
	value, err := s.engine.Get(key.Key)
	if err != nil {
		return nil, err
	}

	if value != nil {
		return &pb.SingleValue{Value: *value}, nil
	}
	return &pb.SingleValue{}, nil
}

func (s *redishServer) Set(ctx context.Context, sr *pb.SetRequest) (*pb.OK, error) {
	log.Printf("SET %v \"%v\"", sr.Key, sr.Value)
	if err := s.engine.Set(sr.Key, sr.Value); err != nil {
		return nil, err
	}
	return &pb.OK{}, nil
}

func (s *redishServer) Dele(ctx context.Context, keys *pb.KeyList) (*pb.SingleValue, error) {
	log.Printf("DELE %v", keys)
	keyStrings := make([]string, len(keys.Keys))
	for i, key := range keys.Keys {
		keyStrings[i] = key.Key
	}
	deleted, err := s.engine.Del(keyStrings)
	if err != nil {
		return nil, err
	}
	return &pb.SingleValue{Value: fmt.Sprintf("%d", deleted)}, nil
}

func (s *redishServer) Incr(ctx context.Context, key *pb.Key) (*pb.SingleValue, error) {
	log.Printf("INCR %v", key.Key)

	value, err := s.engine.Incr(key.Key)
	return &pb.SingleValue{Value: value}, err
}

func (s *redishServer) Decr(ctx context.Context, key *pb.Key) (*pb.SingleValue, error) {
	log.Printf("DECR %v", key.Key)

	value, err := s.engine.Decr(key.Key)
	return &pb.SingleValue{Value: value}, err
}

func main() {
	grpcServer := grpc.NewServer()
	srv := newServer()
	pb.RegisterRedishServer(grpcServer, srv)
	lis, err := net.Listen("tcp", "localhost:4242")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Print("is it working?")
	grpcServer.Serve(lis)
}
