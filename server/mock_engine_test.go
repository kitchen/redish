package main

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

// Incr(key string) (int64, error)
// Decr(key string) (int64, error)
// Incrby(key string, by int64) (int64, error)
// Decrby(key string, by int64) (int64, error)
// Strlen(key string) (int64, error)
// GetSet(key string, value string) (*string, error)
// MGet(keys []string) ([]*string, error)
// MSet(kvs map[string]string) error
// Type(key string) (string, error)
// Expire(key string, seconds int64) (bool, error)
// PExpire(key string, millis int64) (bool, error)
// // seconds since epoch
// ExpireAt(key string, seconds int64) (bool, error)
// // millis since epoch
// PExpireAt(key string, millis int64) (bool, error)
// Persist(key string) (bool, error)
// TTL(key string) (int64, error)
// PTTL(key string) (int64, error)
