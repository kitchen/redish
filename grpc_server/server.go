package grpc_server

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/kitchen/redish/engine"
	pb "github.com/kitchen/redish/redish"
)

type redishServer struct {
	engine engine.Engine
	pb.UnimplementedRedishServer
}

func NewServer(engine engine.Engine) *redishServer {
	s := &redishServer{engine: engine}
	return s
}

func (s *redishServer) Get(ctx context.Context, key *pb.Key) (*pb.SingleValue, error) {
	log.Printf("GET %v", key.Key)
	value, err := s.engine.Get(key.Key)
	if err != nil {
		return nil, err
	}

	if value != nil {
		return &pb.SingleValue{Value: &wrappers.StringValue{Value: *value}}, nil
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

func (s *redishServer) Dele(ctx context.Context, keys *pb.KeyList) (*pb.IntValue, error) {
	log.Printf("DELE %v", keys)
	keyStrings := make([]string, len(keys.Keys))
	for i, key := range keys.Keys {
		keyStrings[i] = key.Key
	}
	deleted, err := s.engine.Del(keyStrings)
	if err != nil {
		return nil, err
	}
	return &pb.IntValue{Value: deleted}, nil
}

func (s *redishServer) Exists(ctx context.Context, keys *pb.KeyList) (*pb.IntValue, error) {
	log.Printf("EXISTS %v", keys)
	keyStrings := make([]string, len(keys.Keys))
	for i, key := range keys.Keys {
		keyStrings[i] = key.Key
	}
	exists, err := s.engine.Exists(keyStrings)
	if err != nil {
		return nil, err
	}
	return &pb.IntValue{Value: exists}, nil
}

func (s *redishServer) Incr(ctx context.Context, key *pb.Key) (*pb.IntValue, error) {
	log.Printf("INCR %v", key.Key)

	value, err := s.engine.Incr(key.Key)
	return &pb.IntValue{Value: value}, err
}

func (s *redishServer) Decr(ctx context.Context, key *pb.Key) (*pb.IntValue, error) {
	log.Printf("DECR %v", key.Key)

	value, err := s.engine.Decr(key.Key)
	return &pb.IntValue{Value: value}, err
}

func (s *redishServer) Incrby(ctx context.Context, kv *pb.KeyIntValue) (*pb.IntValue, error) {
	log.Printf("INCRBY %v %v", kv.Key, kv.Value)
	if value, err := s.engine.Incrby(kv.Key, kv.Value); err == nil {
		return &pb.IntValue{Value: value}, err
	} else {
		return nil, err
	}
}

func (s *redishServer) Decrby(ctx context.Context, kv *pb.KeyIntValue) (*pb.IntValue, error) {
	log.Printf("DECRBY %v %v", kv.Key, kv.Value)

	if value, err := s.engine.Decrby(kv.Key, kv.Value); err == nil {
		return &pb.IntValue{Value: value}, err
	} else {
		return nil, err
	}
}

func (s *redishServer) Strlen(ctx context.Context, key *pb.Key) (*pb.IntValue, error) {
	log.Printf("STRLEN %v", key.Key)
	value, err := s.engine.Strlen(key.Key)
	return &pb.IntValue{Value: value}, err
}

func (s *redishServer) Getset(ctx context.Context, keyvalue *pb.KeyValue) (*pb.SingleValue, error) {
	log.Printf("GETSET %v %v", keyvalue.Key, keyvalue.Value)
	if value, err := s.engine.GetSet(keyvalue.Key, keyvalue.Value); value != nil {
		return &pb.SingleValue{Value: &wrappers.StringValue{Value: *value}}, err
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

	if len(values) != len(keys.Keys) {
		return nil, fmt.Errorf("engine returned the wrong number of values (%v) for the number of keys (%v)", len(values), len(keys.Keys))
	}

	singleValues := make([]*pb.SingleValue, len(values))
	for i, value := range values {
		if value != nil {
			singleValues[i] = &pb.SingleValue{Value: &wrappers.StringValue{Value: *value}}
		} else {
			singleValues[i] = &pb.SingleValue{}
		}
	}

	return &pb.ValueList{Values: singleValues}, nil
}

func (s *redishServer) Mset(ctx context.Context, keyvaluelist *pb.KeyValueList) (*pb.OK, error) {
	kvs := make(map[string]string)
	for _, kv := range keyvaluelist.Pairs {
		kvs[kv.Key] = kv.Value
	}
	err := s.engine.MSet(kvs)
	return &pb.OK{}, err
}

func (s *redishServer) Type(ctx context.Context, key *pb.Key) (*pb.SingleValue, error) {
	typeString, err := s.engine.Type(key.Key)
	return &pb.SingleValue{Value: &wrappers.StringValue{Value: typeString}}, err
}

func (s *redishServer) Expire(ctx context.Context, keyintvalue *pb.KeyIntValue) (*pb.IntValue, error) {
	log.Printf("EXPIRE %v %v", keyintvalue.Key, keyintvalue.Value)
	ret, err := s.engine.Expire(keyintvalue.Key, keyintvalue.Value)
	if ret {
		return &pb.IntValue{Value: 1}, err
	} else {
		return &pb.IntValue{Value: 0}, err
	}
}

func (s *redishServer) Pexpire(ctx context.Context, keyintvalue *pb.KeyIntValue) (*pb.IntValue, error) {
	log.Printf("PEXPIRE %v %v", keyintvalue.Key, keyintvalue.Value)
	ret, err := s.engine.PExpire(keyintvalue.Key, keyintvalue.Value)
	if ret {
		return &pb.IntValue{Value: 1}, err
	} else {
		return &pb.IntValue{Value: 0}, err
	}
}

func (s *redishServer) Expireat(ctx context.Context, keyintvalue *pb.KeyIntValue) (*pb.IntValue, error) {
	log.Printf("EXPIREAT %v %v", keyintvalue.Key, keyintvalue.Value)
	ret, err := s.engine.ExpireAt(keyintvalue.Key, keyintvalue.Value)
	if ret {
		return &pb.IntValue{Value: 1}, err
	} else {
		return &pb.IntValue{Value: 0}, err
	}
}

func (s *redishServer) Pexpireat(ctx context.Context, keyintvalue *pb.KeyIntValue) (*pb.IntValue, error) {
	log.Printf("PEXPIREAT %v %v", keyintvalue.Key, keyintvalue.Value)
	ret, err := s.engine.PExpireAt(keyintvalue.Key, keyintvalue.Value)
	if ret {
		return &pb.IntValue{Value: 1}, err
	} else {
		return &pb.IntValue{Value: 0}, err
	}
}

func (s *redishServer) Persist(ctx context.Context, key *pb.Key) (*pb.IntValue, error) {
	log.Printf("PERSIST %v", key.Key)
	ret, err := s.engine.Persist(key.Key)
	if ret {
		return &pb.IntValue{Value: 1}, err
	} else {
		return &pb.IntValue{Value: 0}, err
	}
}

func (s *redishServer) Ttl(ctx context.Context, key *pb.Key) (*pb.IntValue, error) {
	log.Printf("TTL %v", key.Key)
	ret, err := s.engine.TTL(key.Key)
	return &pb.IntValue{Value: ret}, err
}

func (s *redishServer) Pttl(ctx context.Context, key *pb.Key) (*pb.IntValue, error) {
	log.Printf("PTTL %v", key.Key)
	ret, err := s.engine.PTTL(key.Key)
	return &pb.IntValue{Value: ret}, err
}
