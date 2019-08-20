package grpc_server

import (
	"github.com/kitchen/redish/engine"
	"github.com/stretchr/testify/mock"
)

type mockEngine struct {
	engine.UnimplementedEngine
	mock.Mock
}

func (engine *mockEngine) Get(key string) (*string, error) {
	args := engine.Called(key)
	err := args.Error(1)
	if ret := args.Get(0); ret != nil {
		return ret.(*string), err
	}

	return nil, err
}

func (engine *mockEngine) Set(key string, value string) error {
	args := engine.Called(key, value)
	return args.Error(0)
}

func (engine *mockEngine) Del(keys []string) (int64, error) {
	args := engine.Called(keys)
	return int64(args.Int(0)), args.Error(1)
}

func (engine *mockEngine) Exists(keys []string) (int64, error) {
	args := engine.Called(keys)
	return int64(args.Int(0)), args.Error(1)
}

func (engine *mockEngine) Incr(key string) (int64, error) {
	args := engine.Called(key)
	return int64(args.Int(0)), args.Error(1)
}

func (engine *mockEngine) Decr(key string) (int64, error) {
	args := engine.Called(key)
	return int64(args.Int(0)), args.Error(1)
}

func (engine *mockEngine) Incrby(key string, by int64) (int64, error) {
	args := engine.Called(key, by)
	return int64(args.Int(0)), args.Error(1)
}

func (engine *mockEngine) Decrby(key string, by int64) (int64, error) {
	args := engine.Called(key, by)
	return int64(args.Int(0)), args.Error(1)
}

func (engine *mockEngine) Strlen(key string) (int64, error) {
	args := engine.Called(key)
	return int64(args.Int(0)), args.Error(1)
}

func (engine *mockEngine) GetSet(key string, value string) (*string, error) {
	args := engine.Called(key, value)
	err := args.Error(1)
	if ret := args.Get(0); ret != nil {
		return ret.(*string), err
	}

	return nil, err
}

func (engine *mockEngine) MGet(keys []string) ([]*string, error) {
	args := engine.Called(keys)
	err := args.Error(1)
	if ret := args.Get(0); ret != nil {
		return ret.([]*string), err
	}

	return nil, err
}

func (engine *mockEngine) MSet(kvs map[string]string) error {
	args := engine.Called(kvs)
	return args.Error(0)
}

func (engine *mockEngine) Type(key string) (string, error) {
	args := engine.Called(key)
	return args.String(0), args.Error(1)
}

func (engine *mockEngine) Expire(key string, seconds int64) (bool, error) {
	args := engine.Called(key, seconds)
	return args.Bool(0), args.Error(1)
}

func (engine *mockEngine) PExpire(key string, millis int64) (bool, error) {
	args := engine.Called(key, millis)
	return args.Bool(0), args.Error(1)
}

func (engine *mockEngine) ExpireAt(key string, seconds int64) (bool, error) {
	args := engine.Called(key, seconds)
	return args.Bool(0), args.Error(1)
}

func (engine *mockEngine) PExpireAt(key string, millis int64) (bool, error) {
	args := engine.Called(key, millis)
	return args.Bool(0), args.Error(1)
}

func (engine *mockEngine) Persist(key string) (bool, error) {
	args := engine.Called(key)
	return args.Bool(0), args.Error(1)
}

func (engine *mockEngine) TTL(key string) (int64, error) {
	args := engine.Called(key)
	return int64(args.Int(0)), args.Error(1)
}

func (engine *mockEngine) PTTL(key string) (int64, error) {
	args := engine.Called(key)
	return int64(args.Int(0)), args.Error(1)
}
