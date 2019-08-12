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

func (s *redishServer) Exists(ctx context.Context, keys *pb.KeyList) (*pb.SingleValue, error) {
	log.Printf("EXISTS %v", keys)
	keyStrings := make([]string, len(keys.Keys))
	for i, key := range keys.Keys {
		keyStrings[i] = key.Key
	}
	exists, err := s.engine.Exists(keyStrings)
	if err != nil {
		return nil, err
	}
	return &pb.SingleValue{Value: fmt.Sprintf("%d", exists)}, nil
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

func (s *redishServer) Incrby(ctx context.Context, keyvalue *pb.KeyValue) (*pb.SingleValue, error) {
	log.Printf("INCRBY %v %v", keyvalue.Key, keyvalue.Value)
	if value, err := s.engine.Incrby(keyvalue.Key, keyvalue.Value); err == nil {
		return &pb.SingleValue{Value: value}, err
	} else {
		return nil, err
	}
}

func (s *redishServer) Decrby(ctx context.Context, keyvalue *pb.KeyValue) (*pb.SingleValue, error) {
	log.Printf("DECRBY %v %v", keyvalue.Key, keyvalue.Value)
	if value, err := s.engine.Decrby(keyvalue.Key, keyvalue.Value); err == nil {
		return &pb.SingleValue{Value: value}, err
	} else {
		return nil, err
	}
}

func (s *redishServer) Strlen(ctx context.Context, key *pb.Key) (*pb.SingleValue, error) {
	log.Printf("STRLEN %v", key.Key)
	value, err := s.engine.Strlen(key.Key)
	return &pb.SingleValue{Value: value}, err
}

func (s *redishServer) Getset(ctx context.Context, keyvalue *pb.KeyValue) (*pb.SingleValue, error) {
	log.Printf("GETSET %v %v", keyvalue.Key, keyvalue.Value)
	if value, err := s.engine.GetSet(keyvalue.Key, keyvalue.Value); value != nil {
		return &pb.SingleValue{Value: *value}, err
	} else {
		return &pb.SingleValue{}, err
	}
}

func (s *redishServer) Mget(ctx context.Context, keys *pb.KeyList) (*pb.ValueList, error) {
	log.Printf("MGET %v", keys)
	keyStrings := make([]string, len(keys.Keys))
	for i, key := range keys.Keys {
		keyStrings[i] = key.Key
	}
	values, err := s.engine.MGet(keyStrings)
	if err != nil {
		return nil, err
	}
	singleValues := make([]*pb.SingleValue, len(keys.Keys))
	for i, value := range values {
		if value != nil {
			singleValues[i] = &pb.SingleValue{Value: *value}
		} else {
			singleValues[i] = &pb.SingleValue{}
		}
	}

	return &pb.ValueList{Values: singleValues}, nil
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
